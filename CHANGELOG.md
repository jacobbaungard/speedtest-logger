# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Calendar Versioning](https://calver.org/), with the
format YYYY.MM.MICRO.

## [Unreleased]

## [2023.03.1]
### Changed
- Docker: Update Go from 1.19.3 to 1.20.2
- Dependency: gocron updated from 1.17.1 to 1.18.1
- Dependency: influx-db-client updated from 2.12.0 to 2.12.2
- Dependency: commander-cli/cmd updated from 1.5.0 to 1.6.0
- Dependency: viper from 1.14.0 to 1.15.0 
- Dependency: zerolog from 1.28.0 to 1.29.0
- Dependency: x/net from 0.4.0 to 0.7.0 (CVE-2022-41723)

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
