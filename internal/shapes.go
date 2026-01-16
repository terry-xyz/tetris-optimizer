// Package internal defines the 19 canonical tetromino shapes and matching utilities.
package internal

// Point represents a coordinate offset from the top-left origin.
type Point struct {
	Row, Col int // Row and column offset from origin (0,0)
}

// Tetromino represents a tetromino piece with its shape and label.
type Tetromino struct {
	Label  byte    // 'A'-'Z' identifier for this piece
	Coords []Point // 4 coordinate offsets defining the shape
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
	if len(coords) == 0 { // Handle empty input
		return coords
	}

	minRow, minCol := coords[0].Row, coords[0].Col // Initialize with first point
	for _, p := range coords[1:] {                 // Find minimum row and column
		if p.Row < minRow { // Found smaller row
			minRow = p.Row
		}
		if p.Col < minCol { // Found smaller column
			minCol = p.Col
		}
	}

	normalized := make([]Point, len(coords)) // Allocate result slice
	for i, p := range coords {               // Shift all points by min values
		normalized[i] = Point{Row: p.Row - minRow, Col: p.Col - minCol}
	}

	sortPoints(normalized) // Sort for consistent comparison
	return normalized
}

// sortPoints sorts points in row-major order (by row, then by column).
// Uses bubble sort since n is always 4 (tetromino cells).
func sortPoints(points []Point) {
	n := len(points)
	for i := 0; i < n-1; i++ { // Outer loop: n-1 passes needed for n elements
		for j := 0; j < n-i-1; j++ { // Inner loop: -i because last i elements already sorted
			if points[j].Row > points[j+1].Row || // Sort by row first
				(points[j].Row == points[j+1].Row && points[j].Col > points[j+1].Col) { // Then by column
				points[j], points[j+1] = points[j+1], points[j] // Swap elements
			}
		}
	}
}

// MatchShape checks if the given coordinates match any of the 19 canonical shapes.
func MatchShape(coords []Point) bool {
	if len(coords) != 4 { // Early exit: tetrominoes always have exactly 4 cells
		return false
	}

	normalized := Normalize(coords) // Normalize for comparison

	for _, shape := range CanonicalShapes { // Compare against each canonical shape
		if pointsEqual(normalized, shape) { // Found a match
			return true
		}
	}
	return false
}

// pointsEqual checks if two sorted point slices are equal.
func pointsEqual(a, b []Point) bool {
	if len(a) != len(b) { // Different lengths can't be equal
		return false
	}
	for i := range a { // Compare each point
		if a[i].Row != b[i].Row || a[i].Col != b[i].Col { // Mismatch found
			return false
		}
	}
	return true
}

// ParseGrid extracts coordinates from a 4x4 grid representation.
// Returns the normalized coordinates of all '#' cells.
func ParseGrid(lines []string) []Point {
	var coords []Point             // Collect coordinates
	for row, line := range lines { // Iterate through each row
		for col, ch := range line { // Iterate through each character
			if ch == '#' { // Found a filled cell
				coords = append(coords, Point{Row: row, Col: col})
			}
		}
	}
	return Normalize(coords) // Normalize before returning
}
