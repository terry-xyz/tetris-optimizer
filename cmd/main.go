// Tetris Optimizer - Assembles tetrominoes into the smallest possible square grid.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/terry-xyz/tetris-optimizer/internal"
)

func main() {
	os.Exit(run()) // Wrap in run() so defers execute before exit
}

func run() int {
	filename, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	tmr := internal.NewTimer()                                                 // Initialize timer for progress display
	ctx, cancel := context.WithTimeout(context.Background(), internal.Timeout) // 5-minute timeout from spec
	defer cancel()                                                             // Ensure context is cancelled on exit

	sigChan := make(chan os.Signal, 1)                    // Buffer of 1 ensures signal delivery even if not immediately received
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM) // Register for Ctrl+C and kill signals

	interrupted := make(chan struct{}) // Signals that user interrupted
	go func() {                        // Goroutine to handle interrupts asynchronously
		select {
		case <-sigChan: // User pressed Ctrl+C or sent SIGTERM
			cancel()           // Cancel the solve context
			close(interrupted) // Signal main that we were interrupted
		case <-ctx.Done(): // Context cancelled normally (timeout or completion)
		}
	}()

	parseStart := time.Now()                         // Start timing parse phase
	pieces, parseErr := internal.ParseFile(filename) // Parse and validate input file
	tmr.AddDuration("Parse", time.Since(parseStart)) // Record parse duration

	if parseErr != nil {
		fmt.Println("ERROR") // Spec requires "ERROR" on stdout for invalid input
		fmt.Fprintln(os.Stderr, parseErr)
		return 0 // Exit 0 per spec; error is communicated via stdout message
	}

	progressDone := make(chan struct{}) // Signals progress goroutine completion
	go func() {                         // Background goroutine for progress display
		defer close(progressDone)                        // Signal main when done
		ticker := time.NewTicker(100 * time.Millisecond) // Update progress bar 10 times per second
		defer ticker.Stop()                              // Clean up ticker on exit
		for {
			select {
			case <-ticker.C: // Ticker fired
				tmr.ShowProgress() // Update progress bar
			case <-ctx.Done(): // Solve completed or cancelled
				return
			}
		}
	}()

	solveStart := time.Now()                      // Start timing solve phase
	result := internal.Solve(ctx, pieces)         // Run backtracking solver
	solveDuration := time.Since(solveStart)       // Calculate solve duration
	tmr.AddDuration("Total solve", solveDuration) // Record solve duration

	cancel()            // Stop the context to terminate progress goroutine
	<-progressDone      // Wait for progress goroutine to finish
	tmr.ClearProgress() // Clear progress bar from terminal

	select {
	case <-interrupted: // Check if user interrupted (non-blocking)
		fmt.Println("INTERRUPTED")
		return 0
	default: // Not interrupted, continue
	}

	if result.Timeout || result.Board == nil { // Solver didn't find solution in time
		fmt.Println("TIMEOUT - try with fewer tetrominoes")
		return 0
	}

	fmt.Print(result.Board.String())  // Output solution grid to stdout
	tmr.ShowCompletion(solveDuration) // Show "Solved in X.XXs" if TTY

	return 0
}

// parseArgs parses command line arguments.
// Returns the input filename and any error.
func parseArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("Usage: tetris-optimizer <input-file>")
	}

	if len(args) > 1 {
		return "", fmt.Errorf("Usage: tetris-optimizer <input-file>")
	}

	return args[0], nil
}
