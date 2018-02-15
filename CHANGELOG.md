# Changelog

## [Unreleased](https://github.com/theneedyguy/yikesbot/compare/v1.7.2...HEAD)

## [1.7.2](https://github.com/theneedyguy/yikesbot/compare/v1.7.1...v1.7.2) - 2018-02-15

### Changes

- Fixes graph command

## [1.7.1](https://github.com/theneedyguy/yikesbot/compare/v1.7.0...v1.7.1) - 2018-02-15

### Added

- Command to link to graphical overview
- URL to graphical overview is configured via the config file

## [1.7](https://github.com/theneedyguy/yikesbot/compare/v1.6.1...v1.7.0) - 2018-02-15

### Added

- Added Prometheus metrics endpoint for time series data. Will be used for graphing yikes level over time.

## [1.6.1](https://github.com/theneedyguy/yikesbot/compare/v1.6.0...v1.6.1) - 2018-02-15

### Added

- "Yikerz" is now recognized by the bot.

### Changed

- Got rid of some excess code.
- Yikes are recognized regardless of the capitalization

## [1.6](https://github.com/theneedyguy/yikesbot/compare/v1.5.0...v1.6.0) - 2018-02-15

### Changed

- When outputting the yikes level the emote will vary depending on the amount of the current yikes.

## 1.5.1-2 - 2018-02-14

### Changed

- Fixed problems with ipban not getting detected.

## 1.5.1- 2018-02-14

### Added

- Different messages when Destiny bans someone. Prevents duplicate message error.

## [1.5.0](https://github.com/theneedyguy/yikesbot/tree/v1.5.0) - 2018-02-11

### Added

- Bot now counts "omegayikes" or "OMEGAYIKES" as 50 points. Cooldown is 1 minute.
- Added rosewood18 to list of admins

### Changed

- Changed a function to raise the amount of yikes

### Removed

- Removed some comments

## 1.4.0 - 2018-02-09

### Added

- New admin command: **!topyikes** (Shows the highest yikes level ever achieved since starting the bot)

## 1.3.0 - 2018-02-09

### Added

- Added all features
- Admin Commands
- Default Commands