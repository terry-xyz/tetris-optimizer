// Package internal handles reading and validating tetromino input files.
package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ParseError represents a parsing error with details.
type ParseError struct {
	Message string
	Piece   int // 1-indexed piece number, 0 if not piece-specific
}

func (e *ParseError) Error() string {
	if e.Piece > 0 { // Error is associated with a specific piece
		return fmt.Sprintf("%s at piece %d", e.Message, e.Piece)
	}
	return e.Message
}

// ParseFile reads and validates a tetromino input file.
// Returns a slice of tetrominoes or an error.
func ParseFile(filename string) ([]*Tetromino, error) {
	file, err := os.Open(filename) // Open file for reading
	if err != nil {
		return nil, &ParseError{Message: fmt.Sprintf("cannot open file: %s", err.Error())}
	}
	defer file.Close() // Ensure file is closed on exit

	var lines []string                // Collect all lines from file
	scanner := bufio.NewScanner(file) // bufio.Scanner splits on \n but leaves \r on Windows files
	for scanner.Scan() {              // Read file line by line
		line := scanner.Text()                // Get line without trailing \n
		line = strings.TrimSuffix(line, "\r") // Handle Windows CRLF line endings
		lines = append(lines, line)           // Add line to collection
	}

	if err := scanner.Err(); err != nil { // Check for read errors
		return nil, &ParseError{Message: fmt.Sprintf("error reading file: %s", err.Error())}
	}

	return parseLines(lines) // Process collected lines
}

// parseLines processes the file content and extracts tetrominoes.
func parseLines(lines []string) ([]*Tetromino, error) {
	for len(lines) > 0 && lines[len(lines)-1] == "" { // Remove trailing empty lines
		lines = lines[:len(lines)-1] // Trim last element
	}

	if len(lines) == 0 { // File had no content
		return nil, &ParseError{Message: "empty file"}
	}

	var tetrominoes []*Tetromino // Collect parsed tetrominoes
	pieceNum := 0                // 1-indexed piece counter
	i := 0                       // Current line index

	for i < len(lines) { // Process each tetromino
		pieceNum++ // Increment piece counter

		if pieceNum > 26 { // Max 26 pieces limited by A-Z labeling scheme
			return nil, &ParseError{Message: "too many tetrominoes (max 26)", Piece: pieceNum}
		}

		if i+4 > len(lines) { // Need at least 4 lines for a tetromino
			return nil, &ParseError{Message: "incomplete tetromino (less than 4 lines)", Piece: pieceNum}
		}

		pieceLines := lines[i : i+4] // Extract 4 lines for this tetromino
		i += 4                       // Advance past the 4 lines

		for lineIdx, line := range pieceLines { // Validate each line
			if len(line) != 4 { // Each tetromino line must be exactly 4 chars per spec
				return nil, &ParseError{
					Message: fmt.Sprintf("line %d has %d characters (expected 4)", lineIdx+1, len(line)),
					Piece:   pieceNum,
				}
			}
			for _, ch := range line { // Validate each character
				if ch != '#' && ch != '.' { // Only '#' and '.' allowed
					return nil, &ParseError{
						Message: fmt.Sprintf("invalid character '%c'", ch),
						Piece:   pieceNum,
					}
				}
			}
		}

		coords := ParseGrid(pieceLines) // Extract and normalize coordinates
		if len(coords) != 4 {           // Tetrominoes must have exactly 4 filled cells
			return nil, &ParseError{
				Message: fmt.Sprintf("tetromino has %d cells (expected 4)", len(coords)),
				Piece:   pieceNum,
			}
		}

		if !MatchShape(coords) { // Validate against 19 canonical shapes
			return nil, &ParseError{Message: "invalid tetromino shape", Piece: pieceNum}
		}

		label := byte('A' + pieceNum - 1) // Labels A-Z assigned in input order (1st piece = 'A')
		tetrominoes = append(tetrominoes, &Tetromino{
			Label:  label,
			Coords: coords,
		})

		if i < len(lines) { // More content remains; check separator
			if lines[i] != "" { // Spec requires single blank line between pieces
				return nil, &ParseError{
					Message: "missing blank line separator",
					Piece:   pieceNum,
				}
			}
			i++ // Skip the blank line

			if i < len(lines) && lines[i] == "" { // Check for double blank line (invalid)
				return nil, &ParseError{
					Message: "consecutive blank lines not allowed",
					Piece:   pieceNum + 1,
				}
			}
		}
	}

	if len(tetrominoes) == 0 { // No valid tetrominoes found
		return nil, &ParseError{Message: "no tetrominoes found"}
	}

	return tetrominoes, nil
}
