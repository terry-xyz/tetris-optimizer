package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantFile string
		wantErr  bool
	}{
		{
			name:     "single file",
			args:     []string{"input.txt"},
			wantFile: "input.txt",
			wantErr:  false,
		},
		{
			name:    "no arguments",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "too many arguments",
			args:    []string{"file1.txt", "file2.txt"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := parseArgs(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Error("parseArgs() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("parseArgs() unexpected error: %v", err)
				return
			}

			if file != tt.wantFile {
				t.Errorf("parseArgs() file = %s, want %s", file, tt.wantFile)
			}
		})
	}
}

// TestIntegration_ErrorCases tests bad input files against compiled binary.
func TestIntegration_ErrorCases(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode")
	}

	binary := buildBinary(t)
	defer os.Remove(binary)

	tests := []struct {
		name       string
		input      string
		wantOutput string
	}{
		{
			name: "badexample00 - 5 cells instead of 4",
			input: `####
...#
....
....
`,
			wantOutput: "ERROR",
		},
		{
			name: "badexample01 - diagonal shape",
			input: `...#
..#.
.#..
#...
`,
			wantOutput: "ERROR",
		},
		{
			name: "badexample02 - disconnected pieces",
			input: `...#
...#
#...
#...
`,
			wantOutput: "ERROR",
		},
		{
			name: "badexample03 - empty tetromino",
			input: `....
....
....
....
`,
			wantOutput: "ERROR",
		},
		{
			name: "badexample04 - disconnected pieces",
			input: `..##
....
....
##..
`,
			wantOutput: "ERROR",
		},
		{
			name: "badformat - consecutive blank lines",
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
			wantOutput: "ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := createTempFile(t, tt.input)
			defer os.Remove(inputFile)

			cmd := exec.Command(binary, inputFile)
			output, _ := cmd.Output()

			if !strings.Contains(string(output), tt.wantOutput) {
				t.Errorf("Output = %q, want to contain %q", string(output), tt.wantOutput)
			}
		})
	}
}

func TestIntegration_ValidCases(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode")
	}

	binary := buildBinary(t)
	defer os.Remove(binary)

	tests := []struct {
		name      string
		input     string
		wantEmpty int
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := createTempFile(t, tt.input)
			defer os.Remove(inputFile)

			cmd := exec.Command(binary, inputFile)
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Command failed: %v", err)
			}

			// Count empty cells in output
			emptyCount := strings.Count(string(output), ".")
			if emptyCount != tt.wantEmpty {
				t.Errorf("Empty cells = %d, want %d\nOutput:\n%s", emptyCount, tt.wantEmpty, output)
			}
		})
	}
}

func TestIntegration_NoFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode")
	}

	binary := buildBinary(t)
	defer os.Remove(binary)

	cmd := exec.Command(binary)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		t.Error("Expected error when no file provided")
	}
}

func TestIntegration_NonExistentFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode")
	}

	binary := buildBinary(t)
	defer os.Remove(binary)

	cmd := exec.Command(binary, "nonexistent_file_12345.txt")
	output, _ := cmd.Output()

	if !strings.Contains(string(output), "ERROR") {
		t.Errorf("Output = %q, want to contain ERROR", string(output))
	}
}

// buildBinary compiles the program to a temp file for integration tests.
func buildBinary(t *testing.T) string {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "tetris-optimizer-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tmpFile.Close()
	binaryPath := tmpFile.Name()

	// Build from current directory (cmd/)
	cmd := exec.Command("go", "build", "-o", binaryPath, ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build binary: %v\n%s", err, output)
	}

	return binaryPath
}

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "test_input_*.txt")
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
