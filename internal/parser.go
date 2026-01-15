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
	if e.Piece > 0 {
		return fmt.Sprintf("%s at piece %d", e.Message, e.Piece)
	}
	return e.Message
}

// ParseFile reads and validates a tetromino input file.
// Returns a slice of tetrominoes or an error.
func ParseFile(filename string) ([]*Tetromino, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, &ParseError{Message: fmt.Sprintf("cannot open file: %s", err.Error())}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Normalize CRLF to LF (scanner already strips \n, but \r might remain)
		line = strings.TrimSuffix(line, "\r")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, &ParseError{Message: fmt.Sprintf("error reading file: %s", err.Error())}
	}

	return parseLines(lines)
}

// parseLines processes the file content and extracts tetrominoes.
func parseLines(lines []string) ([]*Tetromino, error) {
	// Remove trailing empty lines
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return nil, &ParseError{Message: "empty file"}
	}

	var tetrominoes []*Tetromino
	pieceNum := 0
	i := 0

	for i < len(lines) {
		pieceNum++

		// Check for maximum pieces (26 = A-Z)
		if pieceNum > 26 {
			return nil, &ParseError{Message: "too many tetrominoes (max 26)", Piece: pieceNum}
		}

		// Extract 4 lines for this tetromino
		if i+4 > len(lines) {
			return nil, &ParseError{Message: "incomplete tetromino (less than 4 lines)", Piece: pieceNum}
		}

		pieceLines := lines[i : i+4]
		i += 4

		// Validate the 4 lines
		for lineIdx, line := range pieceLines {
			if len(line) != 4 {
				return nil, &ParseError{
					Message: fmt.Sprintf("line %d has %d characters (expected 4)", lineIdx+1, len(line)),
					Piece:   pieceNum,
				}
			}
			for _, ch := range line {
				if ch != '#' && ch != '.' {
					return nil, &ParseError{
						Message: fmt.Sprintf("invalid character '%c'", ch),
						Piece:   pieceNum,
					}
				}
			}
		}

		// Parse coordinates and validate shape
		coords := ParseGrid(pieceLines)
		if len(coords) != 4 {
			return nil, &ParseError{
				Message: fmt.Sprintf("tetromino has %d cells (expected 4)", len(coords)),
				Piece:   pieceNum,
			}
		}

		if !MatchShape(coords) {
			return nil, &ParseError{Message: "invalid tetromino shape", Piece: pieceNum}
		}

		// Create tetromino with label
		label := byte('A' + pieceNum - 1)
		tetrominoes = append(tetrominoes, &Tetromino{
			Label:  label,
			Coords: coords,
		})

		// Check for blank line separator or end of file
		if i < len(lines) {
			if lines[i] != "" {
				return nil, &ParseError{
					Message: "missing blank line separator",
					Piece:   pieceNum,
				}
			}
			i++ // Skip the blank line

			// Check for double blank line (invalid)
			if i < len(lines) && lines[i] == "" {
				return nil, &ParseError{
					Message: "consecutive blank lines not allowed",
					Piece:   pieceNum + 1,
				}
			}
		}
	}

	if len(tetrominoes) == 0 {
		return nil, &ParseError{Message: "no tetrominoes found"}
	}

	return tetrominoes, nil
}
