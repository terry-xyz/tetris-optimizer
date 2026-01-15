// Tetris Optimizer - Assembles tetrominoes into the smallest possible square grid.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tetris-optimizer/internal"
)

func main() {
	os.Exit(run())
}

func run() int {
	// Parse command line arguments
	args, showTime, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	filename := args[0]
	tmr := internal.NewTimer(showTime)

	// Set up signal handling for graceful interrupt
	ctx, cancel := context.WithTimeout(context.Background(), internal.Timeout)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	interrupted := make(chan struct{})
	go func() {
		select {
		case <-sigChan:
			cancel()
			close(interrupted)
		case <-ctx.Done():
		}
	}()

	// Parse input file
	parseStart := time.Now()
	pieces, parseErr := internal.ParseFile(filename)
	tmr.AddDuration("Parse", time.Since(parseStart))

	if parseErr != nil {
		fmt.Println("ERROR")
		fmt.Fprintln(os.Stderr, parseErr)
		tmr.ShowTimingBreakdown()
		return 0
	}

	// Start progress display goroutine
	progressDone := make(chan struct{})
	go func() {
		defer close(progressDone)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				tmr.ShowProgress()
			case <-ctx.Done():
				return
			}
		}
	}()

	// Solve
	solveStart := time.Now()
	result := internal.Solve(ctx, pieces)
	solveDuration := time.Since(solveStart)
	tmr.AddDuration("Total solve", solveDuration)

	// Stop progress and clear line
	cancel()
	<-progressDone
	tmr.ClearProgress()

	// Check for interrupt
	select {
	case <-interrupted:
		fmt.Println("INTERRUPTED")
		tmr.ShowTimingBreakdown()
		return 0
	default:
	}

	// Check for timeout
	if result.Timeout || result.Board == nil {
		fmt.Println("TIMEOUT - try with fewer tetrominoes")
		tmr.ShowTimingBreakdown()
		return 0
	}

	// Output solution
	fmt.Print(result.Board.String())
	tmr.ShowCompletion(solveDuration)
	tmr.ShowTimingBreakdown()

	return 0
}

// parseArgs parses command line arguments.
// Returns the file argument, --time flag, and any error.
func parseArgs(args []string) ([]string, bool, error) {
	var files []string
	showTime := false

	for _, arg := range args {
		if arg == "--time" {
			showTime = true
		} else if arg == "-h" || arg == "--help" {
			return nil, false, fmt.Errorf("Usage: tetris-optimizer <input-file> [--time]")
		} else if len(arg) > 0 && arg[0] == '-' {
			return nil, false, fmt.Errorf("unknown option: %s\nUsage: tetris-optimizer <input-file> [--time]", arg)
		} else {
			files = append(files, arg)
		}
	}

	if len(files) == 0 {
		return nil, false, fmt.Errorf("Usage: tetris-optimizer <input-file> [--time]\nError: no input file specified")
	}

	if len(files) > 1 {
		return nil, false, fmt.Errorf("Usage: tetris-optimizer <input-file> [--time]\nError: too many arguments")
	}

	return files, showTime, nil
}
