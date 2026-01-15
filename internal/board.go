// Package internal provides board representation and operations for tetromino placement.
package internal

// Board represents the game board as a 2D grid.
type Board struct {
	Grid [][]byte
	Size int
}

// NewBoard creates a new empty board of the given size.
func NewBoard(size int) *Board {
	grid := make([][]byte, size)
	for i := range grid {
		grid[i] = make([]byte, size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return &Board{Grid: grid, Size: size}
}

// Copy creates a deep copy of the board.
func (b *Board) Copy() *Board {
	newBoard := &Board{
		Grid: make([][]byte, b.Size),
		Size: b.Size,
	}
	for i := range b.Grid {
		newBoard.Grid[i] = make([]byte, b.Size)
		copy(newBoard.Grid[i], b.Grid[i])
	}
	return newBoard
}

// Clear resets the board to empty state.
func (b *Board) Clear() {
	for i := range b.Grid {
		for j := range b.Grid[i] {
			b.Grid[i][j] = '.'
		}
	}
}

// CanPlace checks if a tetromino can be placed at the given position.
// It checks bounds first (early exit optimization), then collision.
func (b *Board) CanPlace(t *Tetromino, row, col int) bool {
	// Check bounds first for all coordinates
	for _, p := range t.Coords {
		r, c := row+p.Row, col+p.Col
		if r < 0 || r >= b.Size || c < 0 || c >= b.Size {
			return false
		}
	}

	// Check collision
	for _, p := range t.Coords {
		r, c := row+p.Row, col+p.Col
		if b.Grid[r][c] != '.' {
			return false
		}
	}

	return true
}

// Place puts a tetromino on the board at the given position.
// Assumes CanPlace has already been called.
func (b *Board) Place(t *Tetromino, row, col int) {
	for _, p := range t.Coords {
		r, c := row+p.Row, col+p.Col
		b.Grid[r][c] = t.Label
	}
}

// String returns the board as a string for output.
func (b *Board) String() string {
	result := make([]byte, 0, b.Size*(b.Size+1))
	for _, row := range b.Grid {
		result = append(result, row...)
		result = append(result, '\n')
	}
	return string(result)
}

// CountEmpty returns the number of empty cells on the board.
func (b *Board) CountEmpty() int {
	count := 0
	for _, row := range b.Grid {
		for _, cell := range row {
			if cell == '.' {
				count++
			}
		}
	}
	return count
}
