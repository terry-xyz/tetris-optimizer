# Tetris Optimizer

Assembles tetrominoes into the smallest possible square grid using backtracking. Standard library only.

## Quick Start

```bash
# Build and run
make build
./tetris-optimizer sample.txt

# Or directly
go run ./cmd sample.txt
```

## Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile binary |
| `make run ARGS=file.txt` | Build and run |
| `make test` | Run all tests |
| `make test-v` | Verbose tests |
| `make fmt` | Format code |
| `make clean` | Remove binary |

## Input Format

Each tetromino is 4 lines of 4 characters (`#` = block, `.` = empty), separated by blank lines:

```
#...
#...
#...
#...

....
####
....
....
```

## Output

Solved grid with pieces labeled A-Z in input order, `.` for empty cells.

**Example** (2 tetrominoes → 3×3 grid):
```
Input:              Output:
....                ABB
.##.                ABB
.##.                A..
....                A..

..#.
..#.
..#.
..#.
```

**Special outputs:**
- `ERROR` — invalid input (malformed tetromino, wrong characters, etc.)
- `TIMEOUT - try with fewer tetrominoes` — solving exceeded 5 minutes
- `INTERRUPTED` — user pressed Ctrl+C
