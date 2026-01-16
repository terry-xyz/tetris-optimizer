// Package internal provides progress display and timing utilities.
package internal

import (
	"fmt"
	"os"
	"time"
)

const (
	// Timeout is the maximum allowed solve time.
	Timeout = 5 * time.Minute
	// ProgressWidth is the width of the progress bar.
	ProgressWidth = 20
)

// Timer tracks elapsed time and provides progress display.
type Timer struct {
	start time.Time // When timer was created
	isTTY bool      // Whether stderr is a terminal
}

// NewTimer creates a new Timer instance.
func NewTimer() *Timer {
	return &Timer{
		start: time.Now(), // Record start time
		isTTY: isTTY(),    // Detect terminal
	}
}

// isTTY checks if stderr is a terminal.
func isTTY() bool {
	fi, err := os.Stderr.Stat() // Get stderr file info
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0 // ModeCharDevice indicates a terminal device
}

// IsTTY returns whether stderr is a TTY.
func (t *Timer) IsTTY() bool {
	return t.isTTY
}

// Elapsed returns time since timer started.
func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.start) // Calculate elapsed time
}

// Remaining returns time remaining until timeout.
func (t *Timer) Remaining() time.Duration {
	remaining := Timeout - t.Elapsed() // Calculate time left
	if remaining < 0 {                 // Clamp to zero
		return 0
	}
	return remaining
}

// IsTimedOut checks if the timeout has been exceeded.
func (t *Timer) IsTimedOut() bool {
	return t.Elapsed() >= Timeout // Compare elapsed to timeout
}

// AddDuration is a no-op kept for compatibility.
func (t *Timer) AddDuration(name string, d time.Duration) {}

// ShowProgress displays the progress bar if TTY is available.
func (t *Timer) ShowProgress() {
	if !t.isTTY { // Skip if not a terminal
		return
	}

	remaining := t.Remaining()               // Get time remaining
	minutes := int(remaining.Minutes())      // Extract minutes
	seconds := int(remaining.Seconds()) % 60 // Extract seconds component

	elapsed := t.Elapsed()                          // Get elapsed time
	progress := float64(elapsed) / float64(Timeout) // Calculate progress ratio
	if progress > 1.0 {                             // Clamp to 1.0
		progress = 1.0
	}

	filled := int(progress * ProgressWidth) // Calculate filled bar width
	empty := ProgressWidth - filled         // Calculate empty bar width

	bar := ""                     // Build progress bar string
	for i := 0; i < filled; i++ { // Add filled characters
		bar += "█"
	}
	for i := 0; i < empty; i++ { // Add empty characters
		bar += "░"
	}

	fmt.Fprintf(os.Stderr, "\rSolving... timeout in: %02d:%02d [%s]", minutes, seconds, bar) // Display progress
}

// ClearProgress clears the progress line if TTY is available.
func (t *Timer) ClearProgress() {
	if !t.isTTY { // Skip if not a terminal
		return
	}
	fmt.Fprintf(os.Stderr, "\r%*s\r", 31+ProgressWidth, "") // Overwrite with spaces (31 = prefix length)
}

// ShowCompletion displays completion message to stderr if TTY.
func (t *Timer) ShowCompletion(solveTime time.Duration) {
	if !t.isTTY { // Skip if not a terminal
		return
	}
	totalTime := t.Elapsed()                                         // Get total elapsed time
	fmt.Fprintf(os.Stderr, "Solved in %.2fs\n", totalTime.Seconds()) // Display completion message
}
