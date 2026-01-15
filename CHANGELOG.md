# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0](https://github.com/terry-xyz/tetris-optimizer/compare/v0.0.1...v2.0.0) (2026-01-15)


### ⚠ BREAKING CHANGES

* add CLI entry point with signal handling

### Added

* add CLI entry point with signal handling ([e563fff](https://github.com/terry-xyz/tetris-optimizer/commit/e563fff069fbe36a214fe2c5b72c120781d8efeb))
* **board:** add 2D grid representation ([d712c0f](https://github.com/terry-xyz/tetris-optimizer/commit/d712c0ff7f8ba3a5a9bc6a4e1f26a19df0ee664f))
* **parser:** add input file validation ([fdaf7d6](https://github.com/terry-xyz/tetris-optimizer/commit/fdaf7d6dd4e9fce4f5287aa08f4b056d2aa842f5))
* **solver:** add backtracking algorithm ([5ec5a50](https://github.com/terry-xyz/tetris-optimizer/commit/5ec5a507a859e080e0754cf6d6a5c4f12268105d))
* **tetromino:** add 19 canonical shape definitions ([deb4b0a](https://github.com/terry-xyz/tetris-optimizer/commit/deb4b0ab262cac2e73a921923f63254f099802df))
* **timer:** add progress display and profiling ([fa6c1dd](https://github.com/terry-xyz/tetris-optimizer/commit/fa6c1dd51adb3ff5fc5cac697306cfce21d8dd45))


### Other

* add GitHub release workflow ([72f6f58](https://github.com/terry-xyz/tetris-optimizer/commit/72f6f58d5907c8965ed8b3421d1875b678d73bfe))
* add test and lint workflow ([8f3afb8](https://github.com/terry-xyz/tetris-optimizer/commit/8f3afb871b38a100149fce6bd9b482bf506e576a))
* replace release-please with standard-version automation ([a86b821](https://github.com/terry-xyz/tetris-optimizer/commit/a86b821c0d6e841c0ae7b18b264aea03a9bb0d3f))
* replace standard-version with release-please ([64b7e39](https://github.com/terry-xyz/tetris-optimizer/commit/64b7e39f23cfc8d1c4524d6cac73befe9d9333a6))

## [2.0.0](https://github.com/terry-xyz/tetris-optimizer/compare/v0.0.1...v2.0.0) (2026-01-15)


### ⚠ BREAKING CHANGES

* add CLI entry point with signal handling

### Added

* add CLI entry point with signal handling ([e563fff](https://github.com/terry-xyz/tetris-optimizer/commit/e563fff069fbe36a214fe2c5b72c120781d8efeb))
* **board:** add 2D grid representation ([d712c0f](https://github.com/terry-xyz/tetris-optimizer/commit/d712c0ff7f8ba3a5a9bc6a4e1f26a19df0ee664f))
* **parser:** add input file validation ([fdaf7d6](https://github.com/terry-xyz/tetris-optimizer/commit/fdaf7d6dd4e9fce4f5287aa08f4b056d2aa842f5))
* **solver:** add backtracking algorithm ([5ec5a50](https://github.com/terry-xyz/tetris-optimizer/commit/5ec5a507a859e080e0754cf6d6a5c4f12268105d))
* **tetromino:** add 19 canonical shape definitions ([deb4b0a](https://github.com/terry-xyz/tetris-optimizer/commit/deb4b0ab262cac2e73a921923f63254f099802df))
* **timer:** add progress display and profiling ([fa6c1dd](https://github.com/terry-xyz/tetris-optimizer/commit/fa6c1dd51adb3ff5fc5cac697306cfce21d8dd45))


### Other

* add GitHub release workflow ([72f6f58](https://github.com/terry-xyz/tetris-optimizer/commit/72f6f58d5907c8965ed8b3421d1875b678d73bfe))
* add test and lint workflow ([8f3afb8](https://github.com/terry-xyz/tetris-optimizer/commit/8f3afb871b38a100149fce6bd9b482bf506e576a))
* replace release-please with standard-version automation ([a86b821](https://github.com/terry-xyz/tetris-optimizer/commit/a86b821c0d6e841c0ae7b18b264aea03a9bb0d3f))
* replace standard-version with release-please ([64b7e39](https://github.com/terry-xyz/tetris-optimizer/commit/64b7e39f23cfc8d1c4524d6cac73befe9d9333a6))

## [1.0.0] - 2026-01-16

### Added

- CLI entry point with signal handling and 5-minute timeout
- Progress bar display with TTY detection
- Backtracking solver algorithm for optimal grid placement
- Input file parser with validation for 1-26 tetrominoes
- 2D board representation with placement and collision detection
- 19 canonical tetromino shape definitions

### Other

- CI workflow with tests and linting
- Automated release workflow with changelog generation

[Unreleased]: https://github.com/terry-xyz/tetris-optimizer/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/terry-xyz/tetris-optimizer/releases/tag/v1.0.0
