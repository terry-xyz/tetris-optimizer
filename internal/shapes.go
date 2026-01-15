// Package internal defines the 19 canonical tetromino shapes and matching utilities.
package internal

// Point represents a coordinate offset from the top-left origin.
type Point struct {
	Row, Col int
}

// Tetromino represents a tetromino piece with its shape and label.
type Tetromino struct {
	Label  byte
	Coords []Point
}

// CanonicalShapes contains all 19 rotational variants of the 7 standard tetrominoes.
// Each shape is represented as coordinate offsets normalized to origin (0,0).
var CanonicalShapes = [][]Point{
	// I-piece (2 variants)
	// ####
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, // I horizontal
	// #
	// #
	// #
	// #
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, // I vertical

	// O-piece (1 variant)
	// ##
	// ##
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, // O

	// T-piece (4 variants)
	// ###
	//  #
	{{0, 0}, {0, 1}, {0, 2}, {1, 1}}, // T down
	// #
	// ##
	// #
	{{0, 0}, {1, 0}, {1, 1}, {2, 0}}, // T right
	//  #
	// ###
	{{0, 1}, {1, 0}, {1, 1}, {1, 2}}, // T up
	//  #
	// ##
	//  #
	{{0, 1}, {1, 0}, {1, 1}, {2, 1}}, // T left

	// S-piece (2 variants)
	//  ##
	// ##
	{{0, 1}, {0, 2}, {1, 0}, {1, 1}}, // S horizontal
	// #
	// ##
	//  #
	{{0, 0}, {1, 0}, {1, 1}, {2, 1}}, // S vertical

	// Z-piece (2 variants)
	// ##
	//  ##
	{{0, 0}, {0, 1}, {1, 1}, {1, 2}}, // Z horizontal
	//  #
	// ##
	// #
	{{0, 1}, {1, 0}, {1, 1}, {2, 0}}, // Z vertical

	// L-piece (4 variants)
	// #
	// #
	// ##
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}}, // L up
	// ###
	// #
	{{0, 0}, {0, 1}, {0, 2}, {1, 0}}, // L right
	// ##
	//  #
	//  #
	{{0, 0}, {0, 1}, {1, 1}, {2, 1}}, // L down
	//   #
	// ###
	{{0, 2}, {1, 0}, {1, 1}, {1, 2}}, // L left

	// J-piece (4 variants)
	//  #
	//  #
	// ##
	{{0, 1}, {1, 1}, {2, 0}, {2, 1}}, // J up
	// #
	// ###
	{{0, 0}, {1, 0}, {1, 1}, {1, 2}}, // J right
	// ##
	// #
	// #
	{{0, 0}, {0, 1}, {1, 0}, {2, 0}}, // J down
	// ###
	//   #
	{{0, 0}, {0, 1}, {0, 2}, {1, 2}}, // J left
}

// Normalize converts a set of coordinates to be relative to origin (0,0).
// Returns sorted coordinates for consistent comparison.
func Normalize(coords []Point) []Point {
	if len(coords) == 0 {
		return coords
	}

	// Find minimum row and column
	minRow, minCol := coords[0].Row, coords[0].Col
	for _, p := range coords[1:] {
		if p.Row < minRow {
			minRow = p.Row
		}
		if p.Col < minCol {
			minCol = p.Col
		}
	}

	// Normalize coordinates
	normalized := make([]Point, len(coords))
	for i, p := range coords {
		normalized[i] = Point{Row: p.Row - minRow, Col: p.Col - minCol}
	}

	// Sort for consistent comparison (row-major order)
	sortPoints(normalized)
	return normalized
}

// sortPoints sorts points in row-major order (by row, then by column).
func sortPoints(points []Point) {
	n := len(points)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if points[j].Row > points[j+1].Row ||
				(points[j].Row == points[j+1].Row && points[j].Col > points[j+1].Col) {
				points[j], points[j+1] = points[j+1], points[j]
			}
		}
	}
}

// MatchShape checks if the given coordinates match any of the 19 canonical shapes.
// Coordinates should be pre-normalized.
func MatchShape(coords []Point) bool {
	if len(coords) != 4 {
		return false
	}

	normalized := Normalize(coords)

	for _, shape := range CanonicalShapes {
		if pointsEqual(normalized, shape) {
			return true
		}
	}
	return false
}

// pointsEqual checks if two sorted point slices are equal.
func pointsEqual(a, b []Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Row != b[i].Row || a[i].Col != b[i].Col {
			return false
		}
	}
	return true
}

// ParseGrid extracts coordinates from a 4x4 grid representation.
// Returns the normalized coordinates of all '#' cells.
func ParseGrid(lines []string) []Point {
	var coords []Point
	for row, line := range lines {
		for col, ch := range line {
			if ch == '#' {
				coords = append(coords, Point{Row: row, Col: col})
			}
		}
	}
	return Normalize(coords)
}
