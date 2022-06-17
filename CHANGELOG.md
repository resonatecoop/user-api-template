# Changelog
All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0-13] - 2022-06-17
### Security
- Update github.com/gogo/protobuf to v1.3.2

## [1.0.0-12] - 2022-06-17
### Added
- This CHANGELOG file

### Changed
- Refactored authorization package's user id extraction feat
- Go module renamed from `user-api` to `user-api-template`

### Removed
- Legacy id for wp compat
- Resonate API specific code (extra models, generated client)

**Note:** We renamed the repository from `resonatecoop/user-api` to `resonatecoop/user-api-template` and created a new repo `resonatecoop/user-api` from user-api-template.
