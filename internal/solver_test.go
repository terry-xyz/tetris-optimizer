package internal

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSolve_EmptyInput(t *testing.T) {
	ctx := context.Background()
	result := Solve(ctx, []*Tetromino{})

	if result.Board == nil {
		t.Fatal("Solve() returned nil board for empty input")
	}
	if result.Board.Size != 0 {
		t.Errorf("Solve() board size = %d, want 0", result.Board.Size)
	}
	if result.Timeout {
		t.Error("Solve() timed out on empty input")
	}
}

func TestSolve_SinglePiece(t *testing.T) {
	ctx := context.Background()

	// Single O-piece
	pieces := []*Tetromino{
		{Label: 'A', Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
	}

	result := Solve(ctx, pieces)

	if result.Board == nil {
		t.Fatal("Solve() returned nil board")
	}
	if result.Timeout {
		t.Error("Solve() timed out")
	}
	if result.Board.Size != 2 {
		t.Errorf("Solve() board size = %d, want 2", result.Board.Size)
	}
	if result.Board.CountEmpty() != 0 {
		t.Errorf("Solve() empty cells = %d, want 0", result.Board.CountEmpty())
	}
}

func TestSolve_TwoPieces(t *testing.T) {
	ctx := context.Background()

	// Two O-pieces -> minimum 3x3 grid with 1 empty cell
	pieces := []*Tetromino{
		{Label: 'A', Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
		{Label: 'B', Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
	}

	result := Solve(ctx, pieces)

	if result.Board == nil {
		t.Fatal("Solve() returned nil board")
	}
	if result.Timeout {
		t.Error("Solve() timed out")
	}
	// 8 cells need at least 3x3=9 grid
	if result.Board.Size < 3 {
		t.Errorf("Solve() board size = %d, want >= 3", result.Board.Size)
	}
}

func TestSolve_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Create many pieces to ensure solve takes time
	pieces := make([]*Tetromino, 10)
	for i := range pieces {
		pieces[i] = &Tetromino{
			Label:  byte('A' + i),
			Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		}
	}

	// Cancel immediately
	cancel()

	result := Solve(ctx, pieces)

	if !result.Timeout {
		t.Error("Solve() should return timeout on cancelled context")
	}
}

func TestSolve_Timeout(t *testing.T) {
	// Use an already-cancelled context to guarantee timeout
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	// Create pieces that would take time to solve
	pieces := make([]*Tetromino, 12)
	for i := range pieces {
		pieces[i] = &Tetromino{
			Label:  byte('A' + i),
			Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		}
	}

	result := Solve(ctx, pieces)

	if !result.Timeout {
		t.Error("Solve() should return timeout on cancelled context")
	}
}

// TestSolve_GoodExamples tests against audit/ spec expected empty cell counts.
func TestSolve_GoodExamples(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantEmpty int // Expected empty cells from spec
	}{
		{
			name: "goodexample00 - single O piece",
			input: `....
.##.
.##.
....
`,
			wantEmpty: 0,
		},
		{
			name: "goodexample01 - 4 pieces",
			input: `...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....
`,
			wantEmpty: 9,
		},
		{
			name: "goodexample02 - 8 pieces",
			input: `...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
`,
			wantEmpty: 4,
		},
		{
			name: "goodexample03 - 11 pieces",
			input: `....
.##.
.##.
....

...#
...#
...#
...#

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..

.###
...#
....
....

##..
.#..
.#..
....

....
..##
.##.
....

##..
.#..
.#..
....

.#..
.##.
..#.
....

....
###.
.#..
....
`,
			wantEmpty: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pieces := parsePiecesFromString(t, tt.input)

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			result := Solve(ctx, pieces)

			if result.Timeout {
				t.Fatal("Solve() timed out")
			}
			if result.Board == nil {
				t.Fatal("Solve() returned nil board")
			}

			gotEmpty := result.Board.CountEmpty()
			if gotEmpty != tt.wantEmpty {
				t.Errorf("Solve() empty cells = %d, want %d\nBoard:\n%s", gotEmpty, tt.wantEmpty, result.Board.String())
			}

			// Verify all pieces are placed
			verifyAllPiecesPlaced(t, result.Board, pieces)
		})
	}
}

// TestSolve_HardExample tests spec's hardexam (12 pieces, 1 empty space).
func TestSolve_HardExample(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow test in short mode") // Can take several seconds
	}
	input := `....
.##.
.##.
....

.#..
.##.
.#..
....

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..

.###
...#
....
....

##..
.#..
.#..
....

....
.##.
.##.
....

....
..##
.##.
....

##..
.#..
.#..
....

.#..
.##.
..#.
....

....
###.
.#..
....
`

	pieces := parsePiecesFromString(t, input)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result := Solve(ctx, pieces)

	if result.Timeout {
		t.Fatal("Solve() timed out")
	}
	if result.Board == nil {
		t.Fatal("Solve() returned nil board")
	}

	gotEmpty := result.Board.CountEmpty()
	if gotEmpty != 1 {
		t.Errorf("Solve() empty cells = %d, want 1\nBoard:\n%s", gotEmpty, result.Board.String())
	}

	verifyAllPiecesPlaced(t, result.Board, pieces)
}

func TestSolve_AllIPieces(t *testing.T) {
	ctx := context.Background()

	// 4 vertical I-pieces should fit in 4x4 perfectly
	pieces := []*Tetromino{
		{Label: 'A', Coords: []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
		{Label: 'B', Coords: []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
		{Label: 'C', Coords: []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
		{Label: 'D', Coords: []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	}

	result := Solve(ctx, pieces)

	if result.Board == nil {
		t.Fatal("Solve() returned nil board")
	}
	if result.Board.Size != 4 {
		t.Errorf("Solve() board size = %d, want 4", result.Board.Size)
	}
	if result.Board.CountEmpty() != 0 {
		t.Errorf("Solve() empty cells = %d, want 0", result.Board.CountEmpty())
	}
}

func TestSolve_MinimumBoardSize(t *testing.T) {
	tests := []struct {
		name      string
		numPieces int
		minSize   int
	}{
		{"1 piece", 1, 2},  // 4 cells -> 2x2
		{"2 pieces", 2, 3}, // 8 cells -> 3x3
		{"3 pieces", 3, 4}, // 12 cells -> 4x4
		{"4 pieces", 4, 4}, // 16 cells -> 4x4
		{"5 pieces", 5, 5}, // 20 cells -> 5x5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Use O-pieces for simplicity
			pieces := make([]*Tetromino, tt.numPieces)
			for i := range pieces {
				pieces[i] = &Tetromino{
					Label:  byte('A' + i),
					Coords: []Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
				}
			}

			result := Solve(ctx, pieces)

			if result.Board == nil {
				t.Fatal("Solve() returned nil board")
			}
			if result.Board.Size < tt.minSize {
				t.Errorf("Solve() board size = %d, want >= %d", result.Board.Size, tt.minSize)
			}
		})
	}
}

// parsePiecesFromString parses tetromino input via temp file to reuse ParseFile logic.
func parsePiecesFromString(t *testing.T, input string) []*Tetromino {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "test_solver_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(input); err != nil {
		tmpFile.Close()
		t.Fatalf("Failed to write temp file: %v", err)
	}
	tmpFile.Close()

	pieces, err := ParseFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	return pieces
}

// verifyAllPiecesPlaced ensures each piece label appears exactly 4 times on the board.
func verifyAllPiecesPlaced(t *testing.T, board *Board, pieces []*Tetromino) {
	t.Helper()

	for _, piece := range pieces {
		count := 0
		for _, row := range board.Grid {
			for _, cell := range row {
				if cell == piece.Label {
					count++
				}
			}
		}
		if count != 4 {
			t.Errorf("Piece %c appears %d times, want 4", piece.Label, count)
		}
	}
}

// verifyBoardContainsLabels checks that labels A through nth piece appear in output string.
func verifyBoardContainsLabels(t *testing.T, output string, numPieces int) {
	t.Helper()

	for i := 0; i < numPieces; i++ {
		label := string(byte('A' + i))
		if !strings.Contains(output, label) {
			t.Errorf("Output missing label %s", label)
		}
	}
}
