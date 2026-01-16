// Package internal implements the backtracking algorithm to find the smallest square grid.
package internal

import (
	"context"
	"math"
)

// Result represents the outcome of solving.
type Result struct {
	Board   *Board // Solution board (nil if timeout)
	Timeout bool   // True if solve was cancelled or timed out
}

// Solve finds the smallest square grid that fits all tetrominoes.
// Returns the solution board or nil if timeout/cancelled.
func Solve(ctx context.Context, pieces []*Tetromino) *Result {
	if len(pieces) == 0 { // No pieces to place
		return &Result{Board: NewBoard(0)}
	}

	minSize := int(math.Ceil(math.Sqrt(float64(4 * len(pieces))))) // Minimum size: ceil(sqrt(total_cells))

	for size := minSize; ; size++ { // Try increasing board sizes until solution found
		select {
		case <-ctx.Done(): // Check for cancellation before attempting
			return &Result{Timeout: true}
		default: // Continue if not cancelled
		}

		b := NewBoard(size)           // Create fresh board for this size
		if solve(ctx, b, pieces, 0) { // Attempt to place all pieces
			return &Result{Board: b} // Solution found
		}

		select {
		case <-ctx.Done(): // Check for cancellation after attempt
			return &Result{Timeout: true}
		default: // Continue to next size
		}
	}
}

// solve recursively places tetrominoes using backtracking.
// Returns true if all pieces are placed successfully.
func solve(ctx context.Context, b *Board, pieces []*Tetromino, idx int) bool {
	select {
	case <-ctx.Done(): // Check for cancellation periodically
		return false
	default: // Continue if not cancelled
	}

	if idx >= len(pieces) { // All pieces placed successfully
		return true
	}

	piece := pieces[idx] // Get current piece to place

	for row := 0; row < b.Size; row++ { // Try each row position
		for col := 0; col < b.Size; col++ { // Try each column position
			if b.CanPlace(piece, row, col) { // Check if piece fits here
				newBoard := b.Copy()            // Create copy for immutable backtracking
				newBoard.Place(piece, row, col) // Place piece on copy

				if solve(ctx, newBoard, pieces, idx+1) { // Recursively place remaining pieces
					*b = *newBoard // Propagate successful solution back up the call stack
					return true
				}
			}
		}
	}

	return false // No valid placement found at this position
}
