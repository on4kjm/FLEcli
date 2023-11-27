# Changelog

## [v0.1.7](https://github.com/on4kjm/FLEcli/tree/v0.1.7) (2023-11-24)

**Fixed bugs:**

- Support for POTA prefix ending with a digit [\#108](https://github.com/on4kjm/FLEcli/issues/108)

## [v0.1.6](https://github.com/on4kjm/FLEcli/tree/v0.1.6) (2023-11-14)

**Implemented enhancements:**

- Support parsing 5-digit POTA reference numbers [\#105](https://github.com/on4kjm/FLEcli/issues/105)
- POTA output filename convention  [\#77](https://github.com/on4kjm/FLEcli/issues/77)

**Closed issues:**

- Create release V 01.6 [\#86](https://github.com/on4kjm/FLEcli/issues/86)
- Fix goreleaser deprecation notices [\#84](https://github.com/on4kjm/FLEcli/issues/84)

## [v0.1.5](https://github.com/on4kjm/FLEcli/tree/v0.1.5) (2022-03-15)

**Implemented enhancements:**

- Implement HomeBrew installation [\#21](https://github.com/on4kjm/FLEcli/issues/21)

**Fixed bugs:**

- S2S contacts not recognized properly [\#78](https://github.com/on4kjm/FLEcli/issues/78)

## [v0.1.4](https://github.com/on4kjm/FLEcli/tree/v0.1.4) (2021-12-09)

**Implemented enhancements:**

- Frequency truncation [\#71](https://github.com/on4kjm/FLEcli/issues/71)
- Can POTA data be added? [\#69](https://github.com/on4kjm/FLEcli/issues/69)

**Fixed bugs:**

- Using a tab to delimit some keyword lines causes error [\#70](https://github.com/on4kjm/FLEcli/issues/70)

**Closed issues:**

- Document new POTA support [\#74](https://github.com/on4kjm/FLEcli/issues/74)
- Enable github dependabot [\#72](https://github.com/on4kjm/FLEcli/issues/72)

## [v0.1.3](https://github.com/on4kjm/FLEcli/tree/v0.1.3) (2020-10-10)

**Implemented enhancements:**

- Give the possibility to generate a SOTA chasser log [\#62](https://github.com/on4kjm/FLEcli/issues/62)
- Align summary display for long calls [\#61](https://github.com/on4kjm/FLEcli/issues/61)

**Closed issues:**

- Add a test to validate the parsing of a large file [\#65](https://github.com/on4kjm/FLEcli/issues/65)
- Bump golang to 1.15.2 [\#59](https://github.com/on4kjm/FLEcli/issues/59)

## [v0.1.2](https://github.com/on4kjm/FLEcli/tree/v0.1.2) (2020-09-25)

**Implemented enhancements:**

- Date, band and mode can be entered in a single line or in separate lines. [\#48](https://github.com/on4kjm/FLEcli/issues/48)
- Support WWFF keyword [\#38](https://github.com/on4kjm/FLEcli/issues/38)
- Support date increment \("day +" and "day ++"\) [\#37](https://github.com/on4kjm/FLEcli/issues/37)
- Support date not prefixed by "date" \(non header mode\) DATE keyword is now optional [\#36](https://github.com/on4kjm/FLEcli/issues/36)
- Support extrapolated date [\#35](https://github.com/on4kjm/FLEcli/issues/35)
- Support different date delimiter [\#34](https://github.com/on4kjm/FLEcli/issues/34)
- Support contact grid locator and MYGRID [\#33](https://github.com/on4kjm/FLEcli/issues/33)
- Add end-to-end testing with Bats [\#32](https://github.com/on4kjm/FLEcli/issues/32)
- Add a test to validate the docker image [\#18](https://github.com/on4kjm/FLEcli/issues/18)
- FLEcli should return a useful exit code [\#12](https://github.com/on4kjm/FLEcli/issues/12)

**Fixed bugs:**

- When inference time gap start time and end time are equal, parsing fails. [\#57](https://github.com/on4kjm/FLEcli/issues/57)
- Missing ADIF field for WWFF "Park to Park" [\#54](https://github.com/on4kjm/FLEcli/issues/54)
- Docker container doesn't seem to work anymore [\#50](https://github.com/on4kjm/FLEcli/issues/50)
- MYSOTA and MYWWFF cannot be overwritten later in the file [\#49](https://github.com/on4kjm/FLEcli/issues/49)
- \(interpolate time\) better error message when missing initial time [\#25](https://github.com/on4kjm/FLEcli/issues/25)

## [v0.1.1](https://github.com/on4kjm/FLEcli/tree/v0.1.1) (2020-08-12)

**Implemented enhancements:**

- Re-visit the release notes [\#22](https://github.com/on4kjm/FLEcli/issues/22)
- Create a Makefile to make the local build easier [\#19](https://github.com/on4kjm/FLEcli/issues/19)

**Fixed bugs:**

- Refactor ADIF error handling [\#28](https://github.com/on4kjm/FLEcli/issues/28)
- The output filename validation should return an error instead of a bool [\#27](https://github.com/on4kjm/FLEcli/issues/27)
- Improve automated tests and coverage [\#24](https://github.com/on4kjm/FLEcli/issues/24)

## [v0.1.0](https://github.com/on4kjm/FLEcli/tree/v0.1.0) (2020-07-29)

**Implemented enhancements:**

- Simplify the input and output specification \(-i and -o\) [\#17](https://github.com/on4kjm/FLEcli/issues/17)
- Create a docker container to run the app [\#15](https://github.com/on4kjm/FLEcli/issues/15)
- Separate the documentation part of the readme [\#13](https://github.com/on4kjm/FLEcli/issues/13)
- Improve version display [\#10](https://github.com/on4kjm/FLEcli/issues/10)
- Make element failing validation more visible in error message [\#6](https://github.com/on4kjm/FLEcli/issues/6)

**Fixed bugs:**

- "e7/z35m/p" is not considered a valid call [\#5](https://github.com/on4kjm/FLEcli/issues/5)
- "sota" keyword is not handled properly [\#4](https://github.com/on4kjm/FLEcli/issues/4)



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*
