// Package internal provides board representation and operations for tetromino placement.
package internal

// Board represents the game board as a 2D grid.
type Board struct {
	Grid [][]byte // 2D slice storing cell values ('.' for empty, 'A'-'Z' for pieces)
	Size int      // Width and height of the square board
}

// NewBoard creates a new empty board of the given size.
func NewBoard(size int) *Board {
	grid := make([][]byte, size) // Allocate rows
	for i := range grid {
		grid[i] = make([]byte, size) // Allocate columns for each row
		for j := range grid[i] {
			grid[i][j] = '.' // '.' represents empty cell in output format
		}
	}
	return &Board{Grid: grid, Size: size}
}

// Copy creates a deep copy of the board.
func (b *Board) Copy() *Board {
	newBoard := &Board{
		Grid: make([][]byte, b.Size), // Allocate new grid
		Size: b.Size,
	}
	for i := range b.Grid {
		newBoard.Grid[i] = make([]byte, b.Size) // Allocate new row
		copy(newBoard.Grid[i], b.Grid[i])       // Copy row contents
	}
	return newBoard
}

// Clear resets the board to empty state.
func (b *Board) Clear() {
	for i := range b.Grid {
		for j := range b.Grid[i] {
			b.Grid[i][j] = '.' // Reset each cell to empty
		}
	}
}

// CanPlace checks if a tetromino can be placed at the given position.
// It checks bounds first (early exit optimization), then collision.
func (b *Board) CanPlace(t *Tetromino, row, col int) bool {
	for _, p := range t.Coords { // Check each cell of the tetromino
		r, c := row+p.Row, col+p.Col                      // Calculate absolute position
		if r < 0 || r >= b.Size || c < 0 || c >= b.Size { // Out of bounds check
			return false
		}
	}

	for _, p := range t.Coords { // Check for collision with existing pieces
		r, c := row+p.Row, col+p.Col // Calculate absolute position
		if b.Grid[r][c] != '.' {     // Cell already occupied
			return false
		}
	}

	return true
}

// Place puts a tetromino on the board at the given position.
// Assumes CanPlace has already been called.
func (b *Board) Place(t *Tetromino, row, col int) {
	for _, p := range t.Coords { // Place each cell of the tetromino
		r, c := row+p.Row, col+p.Col // Calculate absolute position
		b.Grid[r][c] = t.Label       // Mark cell with piece label
	}
}

// String returns the board as a string for output.
func (b *Board) String() string {
	result := make([]byte, 0, b.Size*(b.Size+1)) // Pre-allocate: Size cells + 1 newline per row
	for _, row := range b.Grid {
		result = append(result, row...) // Append entire row
		result = append(result, '\n')   // Add newline after each row
	}
	return string(result)
}

// CountEmpty returns the number of empty cells on the board.
func (b *Board) CountEmpty() int {
	count := 0
	for _, row := range b.Grid {
		for _, cell := range row {
			if cell == '.' { // Cell is empty
				count++
			}
		}
	}
	return count
}
