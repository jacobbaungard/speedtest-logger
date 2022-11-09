# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Calendar Versioning](https://calver.org/), with the
format YYYY.MM.MICRO.

## [Unreleased]

## [2022.11.1]
### Added
- Support InfluxDB 1.8+

### Changed
- Use Go 1.19.3

## [2022.10.2]
### Fixed
- Fixed a bug causing the speedtest to not run unless specifically selecting a
  librespeed server

## [2022.10.1]
### Added
- Initial release with basic functionality for running periodic speedtests and
  saving the results to InfluxDB v2.
