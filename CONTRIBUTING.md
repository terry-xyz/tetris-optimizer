# Contributing to Tetris Optimizer

Thank you for your interest in contributing to Tetris Optimizer! This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/tetris-optimizer.git
   cd tetris-optimizer
   ```
3. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Requirements

- Go 1.25 or higher
- Standard library only (no external dependencies)

## Code Standards

- Follow Go conventions and run `go fmt` before committing
- Write tests for new features
- Keep code simple and readable

## Testing

Run tests using the Makefile:

```bash
make test          # Run all tests
```

Or directly with Go:

```bash
go test -v ./...              # Verbose test output
go test -cover ./...          # Test with coverage
go test -run TestName ./...   # Run specific test
```

## Building

```bash
make build              # Compile binary
make run ARGS=file.txt  # Build and run with input
make fmt                # Format code
make clean              # Remove build artifacts
```

## Submitting Changes

1. Ensure all tests pass
2. Write clear, descriptive commit messages following [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation changes
   - `test:` for test changes
   - `chore:` for maintenance tasks
3. Push to your fork
4. Open a Pull Request against `main`

## Project Structure

When contributing, understand the architecture:

- `cmd/main.go` - Entry point, CLI args, signal handling
- `internal/parser.go` - File reading, validation, tetromino extraction
- `internal/shapes.go` - 19 canonical shapes as coordinate offsets
- `internal/solver.go` - Backtracking algorithm
- `internal/board.go` - 2D slice operations
- `internal/timer.go` - TTY-detected progress bar

## Code Review

All submissions go through code review. Requirements:

- All tests must pass
- Code follows Go best practices
- Documentation is updated where applicable

## Questions?

Open an issue in the [issue tracker](https://github.com/terry-xyz/tetris-optimizer/issues) for questions or discussion.
