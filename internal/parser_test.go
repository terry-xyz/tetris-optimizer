package internal

import (
	"os"
	"testing"
)

func TestParseError(t *testing.T) {
	tests := []struct {
		name    string
		err     *ParseError
		wantStr string
	}{
		{
			name:    "without piece",
			err:     &ParseError{Message: "empty file"},
			wantStr: "empty file",
		},
		{
			name:    "with piece",
			err:     &ParseError{Message: "invalid shape", Piece: 3},
			wantStr: "invalid shape at piece 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.wantStr {
				t.Errorf("Error() = %q, want %q", got, tt.wantStr)
			}
		})
	}
}

func TestParseFile_FileErrors(t *testing.T) {
	_, err := ParseFile("nonexistent_file_12345.txt")
	if err == nil {
		t.Error("ParseFile() should return error for non-existent file")
	}
}

// TestParseFile_BadExamples tests against audit/ spec bad input cases.
func TestParseFile_BadExamples(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			name: "badexample00 - 5 cells instead of 4",
			content: `####
...#
....
....
`,
		},
		{
			name: "badexample01 - diagonal shape",
			content: `...#
..#.
.#..
#...
`,
		},
		{
			name: "badexample02 - disconnected pieces",
			content: `...#
...#
#...
#...
`,
		},
		{
			name: "badexample03 - empty tetromino (0 cells)",
			content: `....
....
....
....
`,
		},
		{
			name: "badexample04 - disconnected pieces",
			content: `..##
....
....
##..
`,
		},
		{
			name: "badformat - consecutive blank lines",
			content: `...#
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := createTempFile(t, tt.content)
			defer os.Remove(tmpFile)

			_, err := ParseFile(tmpFile)
			if err == nil {
				t.Errorf("ParseFile() should return error for %s", tt.name)
			}
		})
	}
}

func TestParseFile_ValidInput(t *testing.T) {
	content := `....
.##.
.##.
....
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	pieces, err := ParseFile(tmpFile)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	if len(pieces) != 1 {
		t.Errorf("ParseFile() got %d pieces, want 1", len(pieces))
	}
	if pieces[0].Label != 'A' {
		t.Errorf("First piece label = %c, want 'A'", pieces[0].Label)
	}
}

func TestParseFile_MultiplePieces(t *testing.T) {
	content := `....
.##.
.##.
....

...#
...#
...#
...#
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	pieces, err := ParseFile(tmpFile)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	if len(pieces) != 2 {
		t.Errorf("ParseFile() got %d pieces, want 2", len(pieces))
	}
	if pieces[0].Label != 'A' || pieces[1].Label != 'B' {
		t.Errorf("Piece labels = %c, %c, want 'A', 'B'", pieces[0].Label, pieces[1].Label)
	}
}

func TestParseFile_EmptyFile(t *testing.T) {
	tmpFile := createTempFile(t, "")
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for empty file")
	}
}

func TestParseFile_InvalidCharacter(t *testing.T) {
	content := `....
.X#.
.##.
....
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for invalid character")
	}
}

func TestParseFile_WrongLineLength(t *testing.T) {
	content := `...
.##.
.##.
....
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for wrong line length")
	}
}

func TestParseFile_IncompleteTetromino(t *testing.T) {
	content := `....
.##.
.##.
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for incomplete tetromino")
	}
}

func TestParseFile_MissingSeparator(t *testing.T) {
	content := `....
.##.
.##.
....
....
####
....
....
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for missing separator")
	}
}

func TestParseFile_ConsecutiveBlankLines(t *testing.T) {
	content := `....
.##.
.##.
....


....
####
....
....
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for consecutive blank lines")
	}
}

func TestParseFile_WrongCellCount(t *testing.T) {
	content := `....
.#..
.##.
....
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for wrong cell count")
	}
}

func TestParseFile_InvalidShape(t *testing.T) {
	content := `...#
..#.
.#..
#...
`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	_, err := ParseFile(tmpFile)
	if err == nil {
		t.Error("ParseFile() should return error for invalid shape")
	}
}

func TestParseFile_TrailingNewlines(t *testing.T) {
	content := `....
.##.
.##.
....


`
	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	pieces, err := ParseFile(tmpFile)
	if err != nil {
		t.Fatalf("ParseFile() error = %v", err)
	}
	if len(pieces) != 1 {
		t.Errorf("ParseFile() got %d pieces, want 1", len(pieces))
	}
}

// TestParseFile_AllShapeTypes validates parsing of all 7 base tetromino types.
func TestParseFile_AllShapeTypes(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			name: "I horizontal",
			content: `....
####
....
....
`,
		},
		{
			name: "I vertical",
			content: `#...
#...
#...
#...
`,
		},
		{
			name: "O piece",
			content: `....
.##.
.##.
....
`,
		},
		{
			name: "T piece",
			content: `###.
.#..
....
....
`,
		},
		{
			name: "S piece",
			content: `.##.
##..
....
....
`,
		},
		{
			name: "Z piece",
			content: `##..
.##.
....
....
`,
		},
		{
			name: "L piece",
			content: `#...
#...
##..
....
`,
		},
		{
			name: "J piece",
			content: `.#..
.#..
##..
....
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := createTempFile(t, tt.content)
			defer os.Remove(tmpFile)

			pieces, err := ParseFile(tmpFile)
			if err != nil {
				t.Errorf("ParseFile() error = %v", err)
			}
			if len(pieces) != 1 {
				t.Errorf("ParseFile() got %d pieces, want 1", len(pieces))
			}
		})
	}
}

// createTempFile creates a temporary file with the given content for test input.
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := f.WriteString(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		t.Fatalf("Failed to write temp file: %v", err)
	}
	f.Close()
	return f.Name()
}
