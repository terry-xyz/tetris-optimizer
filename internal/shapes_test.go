package internal

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name   string
		coords []Point
		want   []Point
	}{
		{
			name:   "empty",
			coords: []Point{},
			want:   []Point{},
		},
		{
			name:   "already normalized",
			coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			want:   []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name:   "offset from origin",
			coords: []Point{{2, 3}, {2, 4}, {3, 3}, {3, 4}},
			want:   []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name:   "negative offset",
			coords: []Point{{-1, -1}, {-1, 0}, {0, -1}, {0, 0}},
			want:   []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name:   "unsorted input",
			coords: []Point{{1, 1}, {0, 0}, {1, 0}, {0, 1}},
			want:   []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Normalize(tt.coords)
			if !pointsEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatchShape_AllCanonicalShapes ensures all 19 rotational variants are recognized.
func TestMatchShape_AllCanonicalShapes(t *testing.T) {
	for i, shape := range CanonicalShapes {
		t.Run(shapeName(i), func(t *testing.T) {
			if !MatchShape(shape) {
				t.Errorf("MatchShape(%v) = false, want true", shape)
			}
		})
	}
}

// TestMatchShape_ShiftedShapes verifies normalization handles arbitrary offsets.
func TestMatchShape_ShiftedShapes(t *testing.T) {
	for i, shape := range CanonicalShapes {
		t.Run(shapeName(i)+"_shifted", func(t *testing.T) {
			// Shift by (5, 5)
			shifted := make([]Point, len(shape))
			for j, p := range shape {
				shifted[j] = Point{Row: p.Row + 5, Col: p.Col + 5}
			}
			if !MatchShape(shifted) {
				t.Errorf("MatchShape(shifted %v) = false, want true", shape)
			}
		})
	}
}

func TestMatchShape_InvalidShapes(t *testing.T) {
	tests := []struct {
		name   string
		coords []Point
	}{
		{
			name:   "diagonal",
			coords: []Point{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
		},
		{
			name:   "disconnected",
			coords: []Point{{0, 0}, {0, 1}, {2, 0}, {2, 1}},
		},
		{
			name:   "three cells",
			coords: []Point{{0, 0}, {0, 1}, {0, 2}},
		},
		{
			name:   "five cells",
			coords: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
		},
		{
			name:   "plus shape",
			coords: []Point{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if MatchShape(tt.coords) {
				t.Errorf("MatchShape(%v) = true, want false", tt.coords)
			}
		})
	}
}

// TestMatchShape_CountIs19 validates spec requirement: I×2, O×1, T×4, S×2, Z×2, L×4, J×4 = 19.
func TestMatchShape_CountIs19(t *testing.T) {
	if len(CanonicalShapes) != 19 {
		t.Errorf("CanonicalShapes has %d shapes, want 19", len(CanonicalShapes))
	}
}

func TestParseGrid(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []Point
	}{
		{
			name:  "O-piece centered",
			lines: []string{"....", ".##.", ".##.", "...."},
			want:  []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name:  "I-piece horizontal",
			lines: []string{"....", "####", "....", "...."},
			want:  []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			name:  "I-piece vertical",
			lines: []string{"...#", "...#", "...#", "...#"},
			want:  []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		},
		{
			name:  "L-piece",
			lines: []string{"#...", "#...", "##..", "...."},
			want:  []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}},
		},
		{
			name:  "T-piece",
			lines: []string{"###.", ".#..", "....", "...."},
			want:  []Point{{0, 0}, {0, 1}, {0, 2}, {1, 1}},
		},
		{
			name:  "S-piece",
			lines: []string{".##.", "##..", "....", "...."},
			want:  []Point{{0, 1}, {0, 2}, {1, 0}, {1, 1}},
		},
		{
			name:  "Z-piece",
			lines: []string{"##..", ".##.", "....", "...."},
			want:  []Point{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseGrid(tt.lines)
			if !pointsEqual(got, tt.want) {
				t.Errorf("ParseGrid() = %v, want %v", got, tt.want)
			}
			// Also verify it matches a canonical shape
			if !MatchShape(got) {
				t.Errorf("ParseGrid() result doesn't match any canonical shape")
			}
		})
	}
}

func TestPointsEqual(t *testing.T) {
	tests := []struct {
		name string
		a, b []Point
		want bool
	}{
		{
			name: "equal",
			a:    []Point{{0, 0}, {0, 1}},
			b:    []Point{{0, 0}, {0, 1}},
			want: true,
		},
		{
			name: "different length",
			a:    []Point{{0, 0}},
			b:    []Point{{0, 0}, {0, 1}},
			want: false,
		},
		{
			name: "different values",
			a:    []Point{{0, 0}, {0, 1}},
			b:    []Point{{0, 0}, {0, 2}},
			want: false,
		},
		{
			name: "empty",
			a:    []Point{},
			b:    []Point{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointsEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("pointsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper to get readable shape names
func shapeName(idx int) string {
	names := []string{
		"I_horiz", "I_vert",
		"O",
		"T_down", "T_right", "T_up", "T_left",
		"S_horiz", "S_vert",
		"Z_horiz", "Z_vert",
		"L_up", "L_right", "L_down", "L_left",
		"J_up", "J_right", "J_down", "J_left",
	}
	if idx < len(names) {
		return names[idx]
	}
	return "unknown"
}
