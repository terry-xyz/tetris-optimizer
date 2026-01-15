// Package internal implements the backtracking algorithm to find the smallest square grid.
package internal

import (
	"context"
	"math"
)

// Result represents the outcome of solving.
type Result struct {
	Board   *Board
	Timeout bool
}

// Solve finds the smallest square grid that fits all tetrominoes.
// Returns the solution board or nil if timeout/cancelled.
func Solve(ctx context.Context, pieces []*Tetromino) *Result {
	if len(pieces) == 0 {
		return &Result{Board: NewBoard(0)}
	}

	// Calculate minimum possible size: ceil(sqrt(4 * num_pieces))
	minSize := int(math.Ceil(math.Sqrt(float64(4 * len(pieces)))))

	// Try increasing sizes until solution found
	for size := minSize; ; size++ {
		// Check for cancellation
		select {
		case <-ctx.Done():
			return &Result{Timeout: true}
		default:
		}

		b := NewBoard(size)
		if solve(ctx, b, pieces, 0) {
			return &Result{Board: b}
		}

		// Check for cancellation after attempt
		select {
		case <-ctx.Done():
			return &Result{Timeout: true}
		default:
		}
	}
}

// solve recursively places tetrominoes using backtracking.
// Returns true if all pieces are placed successfully.
func solve(ctx context.Context, b *Board, pieces []*Tetromino, idx int) bool {
	// Check for cancellation periodically
	select {
	case <-ctx.Done():
		return false
	default:
	}

	// All pieces placed successfully
	if idx >= len(pieces) {
		return true
	}

	piece := pieces[idx]

	// Try each position on the board (top-left scan)
	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {
			if b.CanPlace(piece, row, col) {
				// Create a copy for immutable backtracking
				newBoard := b.Copy()
				newBoard.Place(piece, row, col)

				if solve(ctx, newBoard, pieces, idx+1) {
					// Copy solution back to original board
					*b = *newBoard
					return true
				}
			}
		}
	}

	return false
}
