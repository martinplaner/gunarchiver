# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

Currently no changes or additions.

## [0.3.0] - 2017-06-07
### Added
- Implemented proper progress support for .rar and .tar.gz archives.
  * Thanks to *Kidswiss* for the `numFiles` implementation for rar archives (https://github.com/martinplaner/gunarchiver/pull/1).

### Changed
- Code cleanup, implemented linter recommendations and added CI pipeline to improve code quality.

### Fixed
- Properly handle error case when root directory cannot be created.


## [0.2.1] - 2017-03-11
### Fixed
- Fixed files getting trashed when extraction is canceled

## [0.2.0] - 2017-03-11
### Added
- Support for .rar files added
- Added change log (this file)

### Fixed
- Fixed broken .tar.gz support in binary (worked in tests)


## [0.1.0] - 2017-03-10
### Added
- Initial public release (only Windows UI in english implemented)
- Support for .zip and .tar.gz


[Unreleased]: https://github.com/martinplaner/gunarchiver/tree/develop
[0.3.0]: https://github.com/martinplaner/gunarchiver/releases/tag/v0.3.0
[0.2.1]: https://github.com/martinplaner/gunarchiver/releases/tag/v0.2.1
[0.2.0]: https://github.com/martinplaner/gunarchiver/releases/tag/v0.2.0
[0.1.0]: https://github.com/martinplaner/gunarchiver/releases/tag/v0.1.0
