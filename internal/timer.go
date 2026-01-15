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
	start     time.Time
	isTTY     bool
	showTime  bool
	durations map[string]time.Duration
}

// NewTimer creates a new Timer instance.
func NewTimer(showTime bool) *Timer {
	return &Timer{
		start:     time.Now(),
		isTTY:     isTTY(),
		showTime:  showTime,
		durations: make(map[string]time.Duration),
	}
}

// isTTY checks if stderr is a terminal.
func isTTY() bool {
	fi, err := os.Stderr.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// IsTTY returns whether stderr is a TTY.
func (t *Timer) IsTTY() bool {
	return t.isTTY
}

// Elapsed returns time since timer started.
func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.start)
}

// Remaining returns time remaining until timeout.
func (t *Timer) Remaining() time.Duration {
	remaining := Timeout - t.Elapsed()
	if remaining < 0 {
		return 0
	}
	return remaining
}

// IsTimedOut checks if the timeout has been exceeded.
func (t *Timer) IsTimedOut() bool {
	return t.Elapsed() >= Timeout
}

// Track records time for an operation.
func (t *Timer) Track(name string, start time.Time) {
	t.durations[name] = t.durations[name] + time.Since(start)
}

// AddDuration adds duration to an operation.
func (t *Timer) AddDuration(name string, d time.Duration) {
	t.durations[name] = t.durations[name] + d
}

// ShowProgress displays the progress bar if TTY is available.
func (t *Timer) ShowProgress() {
	if !t.isTTY {
		return
	}

	remaining := t.Remaining()
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60

	// Calculate progress percentage
	elapsed := t.Elapsed()
	progress := float64(elapsed) / float64(Timeout)
	if progress > 1.0 {
		progress = 1.0
	}

	filled := int(progress * ProgressWidth)
	empty := ProgressWidth - filled

	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := 0; i < empty; i++ {
		bar += "░"
	}

	fmt.Fprintf(os.Stderr, "\rTime remaining: %02d:%02d [%s]", minutes, seconds, bar)
}

// ClearProgress clears the progress line if TTY is available.
func (t *Timer) ClearProgress() {
	if !t.isTTY {
		return
	}
	// Clear the line with spaces and return to beginning
	fmt.Fprintf(os.Stderr, "\r%*s\r", 50, "")
}

// ShowCompletion displays completion message to stderr if TTY.
func (t *Timer) ShowCompletion(solveTime time.Duration) {
	if !t.isTTY {
		return
	}
	totalTime := t.Elapsed()
	fmt.Fprintf(os.Stderr, "Solved in %.2fs (total: %.2fs)\n",
		solveTime.Seconds(), totalTime.Seconds())
}

// ShowTimingBreakdown displays detailed timing if --time flag is set.
func (t *Timer) ShowTimingBreakdown() {
	if !t.showTime {
		return
	}

	fmt.Fprintln(os.Stderr, "=== Timing Breakdown ===")

	// Show tracked durations in a specific order
	order := []string{"Parse", "Shape match", "Bounds check", "Collision", "Board copy", "Total solve"}
	for _, name := range order {
		if d, ok := t.durations[name]; ok {
			fmt.Fprintf(os.Stderr, "%-16s %4dms\n", name+":", d.Milliseconds())
		}
	}

	fmt.Fprintf(os.Stderr, "%-16s %4dms\n", "Total:", t.Elapsed().Milliseconds())
}
