package internal

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"size 0", 0},
		{"size 1", 1},
		{"size 2", 2},
		{"size 4", 4},
		{"size 10", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBoard(tt.size)
			if b.Size != tt.size {
				t.Errorf("NewBoard(%d).Size = %d, want %d", tt.size, b.Size, tt.size)
			}
			if len(b.Grid) != tt.size {
				t.Errorf("NewBoard(%d) grid rows = %d, want %d", tt.size, len(b.Grid), tt.size)
			}
			for i, row := range b.Grid {
				if len(row) != tt.size {
					t.Errorf("NewBoard(%d) grid row %d cols = %d, want %d", tt.size, i, len(row), tt.size)
				}
				for j, cell := range row {
					if cell != '.' {
						t.Errorf("NewBoard(%d) cell [%d][%d] = %c, want '.'", tt.size, i, j, cell)
					}
				}
			}
		})
	}
}

// TestBoardCopy verifies deep copy behavior, critical for immutable backtracking.
func TestBoardCopy(t *testing.T) {
	b := NewBoard(4)
	b.Grid[0][0] = 'A'
	b.Grid[1][1] = 'B'

	copy := b.Copy()

	if copy.Size != b.Size {
		t.Errorf("Copy().Size = %d, want %d", copy.Size, b.Size)
	}
	if copy.Grid[0][0] != 'A' || copy.Grid[1][1] != 'B' {
		t.Error("Copy() did not preserve values")
	}

	// Mutation must not propagate back to original (required for backtracking)
	copy.Grid[0][0] = 'X'
	if b.Grid[0][0] != 'A' {
		t.Error("Modifying copy affected original board")
	}
}

func TestBoardClear(t *testing.T) {
	b := NewBoard(3)
	b.Grid[0][0] = 'A'
	b.Grid[1][1] = 'B'
	b.Grid[2][2] = 'C'

	b.Clear()

	for i, row := range b.Grid {
		for j, cell := range row {
			if cell != '.' {
				t.Errorf("Clear() left cell [%d][%d] = %c, want '.'", i, j, cell)
			}
		}
	}
}

// TestBoardCanPlace tests bounds checking before collision (early exit optimization).
func TestBoardCanPlace(t *testing.T) {
	piece := &Tetromino{
		Label:  'A',
		Coords: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, // I-piece horizontal
	}

	tests := []struct {
		name     string
		size     int
		row, col int
		occupied []Point
		want     bool
	}{
		{"fits at origin", 4, 0, 0, nil, true},
		{"out of bounds right", 4, 0, 1, nil, false},
		{"out of bounds bottom", 3, 0, 0, nil, false},
		{"collision", 4, 0, 0, []Point{{0, 2}}, false},
		{"negative row", 4, -1, 0, nil, false},
		{"negative col", 4, 0, -1, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBoard(tt.size)
			for _, p := range tt.occupied {
				b.Grid[p.Row][p.Col] = 'X'
			}
			got := b.CanPlace(piece, tt.row, tt.col)
			if got != tt.want {
				t.Errorf("CanPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardPlace(t *testing.T) {
	b := NewBoard(4)
	piece := &Tetromino{
		Label:  'A',
		Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, // O-piece
	}

	b.Place(piece, 1, 1)

	expected := []Point{{1, 1}, {1, 2}, {2, 1}, {2, 2}}
	for _, p := range expected {
		if b.Grid[p.Row][p.Col] != 'A' {
			t.Errorf("Place() cell [%d][%d] = %c, want 'A'", p.Row, p.Col, b.Grid[p.Row][p.Col])
		}
	}
}

func TestBoardString(t *testing.T) {
	b := NewBoard(2)
	b.Grid[0][0] = 'A'
	b.Grid[0][1] = 'A'
	b.Grid[1][0] = '.'
	b.Grid[1][1] = 'B'

	got := b.String()
	want := "AA\n.B\n"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestBoardCountEmpty(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		occupied int
		want     int
	}{
		{"empty 2x2", 2, 0, 4},
		{"full 2x2", 2, 4, 0},
		{"half 4x4", 4, 8, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBoard(tt.size)
			count := 0
			for i := 0; i < tt.size && count < tt.occupied; i++ {
				for j := 0; j < tt.size && count < tt.occupied; j++ {
					b.Grid[i][j] = 'X'
					count++
				}
			}
			got := b.CountEmpty()
			if got != tt.want {
				t.Errorf("CountEmpty() = %d, want %d", got, tt.want)
			}
		})
	}
}
