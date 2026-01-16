# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

### [1.0.2](https://github.com/terry-xyz/tetris-optimizer/compare/v1.0.1...v1.0.2) (2026-01-16)

### Documentation

* add contributing guidelines ([#35](https://github.com/terry-xyz/tetris-optimizer/commit/eac719384680ef6747aebcff19d7cd5733b021aa))
* **readme:** add release, godoc, and codecov badges ([#34](https://github.com/terry-xyz/tetris-optimizer/commit/53f5a304a0cd3dce296971dd8a995d2ca533bfdb))

### Other

* add code coverage reporting with Codecov ([#33](https://github.com/terry-xyz/tetris-optimizer/commit/02fb4af31ab2f2df849f7024c7557df0cacbbf9e))

### [1.0.1](https://github.com/terry-xyz/tetris-optimizer/compare/v1.0.0...v1.0.1) (2026-01-16)

### Fixed

* **changelog:** restore content stripped during v1.0.0 release ([#31](https://github.com/terry-xyz/tetris-optimizer/commit/558fa4fb81e355eab24072ba22df17fd732acf0c))
* **ci:** use split instead of buggy regex for changelog sections ([#32](https://github.com/terry-xyz/tetris-optimizer/commit/cd2fb5a38f59985daff2157fad7e5785e32a8ab7))

### Documentation

* **readme:** add project badges ([#27](https://github.com/terry-xyz/tetris-optimizer/commit/cda1894d0140ea9d91101a1a357fc5c4baec03fb))

## [1.0.0](https://github.com/terry-xyz/tetris-optimizer/compare/v0.2.1...v1.0.0) (2026-01-16)

### Added

* **docs:** add comprehensive README documentation ([#26](https://github.com/terry-xyz/tetris-optimizer/commit/fdbd2c5))

### Fixed

* **ci:** reorder changelog sections to follow Keep a Changelog format ([#20](https://github.com/terry-xyz/tetris-optimizer/commit/6a23655))

### Other

* add Makefile for build automation ([#23](https://github.com/terry-xyz/tetris-optimizer/commit/4aa75bb))
* add sample input file ([#24](https://github.com/terry-xyz/tetris-optimizer/commit/0f095c1))
* **cli:** move entry point to cmd/ package ([#21](https://github.com/terry-xyz/tetris-optimizer/commit/740a076))
* **internal:** add inline comments to implementation ([#22](https://github.com/terry-xyz/tetris-optimizer/commit/ddf6556))
* **internal:** add unit tests for all packages ([#25](https://github.com/terry-xyz/tetris-optimizer/commit/dc445dc))

### [0.2.1](https://github.com/terry-xyz/tetris-optimizer/compare/v0.2.0...v0.2.1) (2026-01-16)

### Fixed

* **ci:** update breaking change regex to match feat!(scope) format ([#17](https://github.com/terry-xyz/tetris-optimizer/commit/e29db5d0d43886aa6141102b98292aa7bfa6b473))

## [0.2.0](https://github.com/terry-xyz/tetris-optimizer/compare/v0.1.0...v0.2.0) (2026-01-16)

### Added

* **board:** add 2D board operations for placement ([#8](https://github.com/terry-xyz/tetris-optimizer/commit/2de5b7ad14419848d518a8b6d6e5c4279e008c63))
* **solver:** add backtracking solver with timeout ([#9](https://github.com/terry-xyz/tetris-optimizer/commit/43a686923729f658ec528b2de59a33930685885e))
* **timer:** add TTY-aware progress and timing display ([#10](https://github.com/terry-xyz/tetris-optimizer/commit/eefa7e3355ef3850d125db8a33b799cbe1752751))

## [0.1.0](https://github.com/terry-xyz/tetris-optimizer/compare/v0.0.1...v0.1.0) (2026-01-16)

### Added

* **parser:** add input file parser with validation ([#7](https://github.com/terry-xyz/tetris-optimizer/commit/77ed705067d7be556a4574df03162a098f48fd50))

### 0.0.1 (2026-01-16)

### Added

* **tetromino:** add 19 canonical shape definitions ([#4](https://github.com/terry-xyz/tetris-optimizer/commit/8c61721172f0d1f59e9c12ed31528bc9868ac3a6))

### Other

* add CI and release workflows ([#2](https://github.com/terry-xyz/tetris-optimizer/commit/9369aa1ba957c20142a211cfa5d443a6b3fd4c6a))
