# What's new?

## v0.1.5

* Fix "S2S contacts not recognized properly" (issue #78)


## Previous releases

### v0.1.4

* Support POTA ADIF output
* Tabs can be used as delimiter in the header section (issue #70)
* Frequency can now be specified up to 500Hz without truncation (issue #71)
* Enabled Homebrew distribution



### v0.1.3

* Enable FLEcli to generate CSV chaser logs
* Fix the display for longer calls (issue #61)

### v0.1.2

* DATE keyword is now optional
* Date can have several delimiter ("-", "/", ".", or " ")
* Partial dates can be entered ("20-9-6" => "2020-09-06")
* The new (FLE v3) "DAY" keyword is now supported (increment is 10 max)
* Date, band, and mode can be specified on a same line, even within a QSO line
* Correct processing of optional WWFF keyword
* Time is now correctly inferred when start and end of gap is in the same minute 
* Correct some typos and bugs

### v0.1.1
* Improved test coverage
* Improved build automation
* Improved release notes publication

### v0.1.0
* First public MVP (Minimal Viable Product relase). Supports only SOTA and WWFF type log files. Some header keywords are missing as well as date increments.

### v0.0.0
* Initial release
