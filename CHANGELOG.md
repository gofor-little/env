# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v1.0.3 - 2021-12-12
### Added
* Added better error handling when parsing an ```.env``` file.

### Fixed
* Fixed an issue where environment variables with ```=``` characters in them would not be parsed correctly.

## v1.0.2 - 2021-09-07
### Fixed
* Cleaned filepaths before use.
* Tightened permissions around file.

## v1.0.1 - 2021-08-21
### Added
* Added Go 1.17 support.
* Added a changelog.
* Added a code of conduct.

## v1.0.0 - 2021-06-22
### Added
* Added Semantic Pull Requests configuration file.
* Added a pull request template.

## v0.4.4 - 2021-02-20
### Added
* Added Go 1.16 support.

## v0.4.3 - 2020-12-31
### Changed
* Updated readme.

## v0.4.2 - 2020-12-30
### Added
* Added GitHub Stale action.

## v0.4.1 - 2020-12-30
### Changed
* Formatted several files.

## v0.4.0 - 2020-11-03
### Added
* Added Go 1.15 support.

### Fixed
* Fixed a bug not allowing quotes to be written to file when the ```Write``` function was called.

## v0.3.1 - 2020-09-03
### Changed
* Updated tests.

## v0.3.0 - 2020-09-03
### Changed
* **BREAKING**: Added ```setAfterWrite``` parameter to ```Write``` function.

## v0.2.0 - 2020-08-16
### Added
* Added ```Get```, ```MustGet``` and ```Set``` functions.

## v0.1.2 - 2020-06-15
### Fixed
* Fixed uncaught error.

## v0.1.1 - 2020-06-11
### Added
* Added tests.

## v0.1.0 - 2020-06-10
### Added
* Added ```Load``` and ```Write``` functions.