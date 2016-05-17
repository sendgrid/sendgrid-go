# Change Log
All notable changes to this project will be documented in this file.

## [2.1.0] - 2015-6-09
### Changed
- Singular-verion functions in mail.go such as AddTo(), which used to take strings and mail.Address, have been switched to variadic parameters to allow users to pass single or multiple parameters using the same function, simplifying the API.

## [2.0.0] - 2015-5-02
### Changed
- Fixed a nasty bug with orphaned connections but drops support for Go versions < 1.3. Thanks [trinchan](https://github.com/sendgrid/sendgrid-go/pull/24)

## [1.2.0] - 2015-4-27
### Added
- Support for API keys

