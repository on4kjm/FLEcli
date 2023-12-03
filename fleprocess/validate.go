package fleprocess

/*
Copyright © 2020 Jean-Marc Meessen, ON4KJM <on4kjm@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var validSotaRegexp = regexp.MustCompile(`^[0-9A-Z]{1,3}/[A-Z]{2}-[\d]{3}$`)

// ValidateSota verifies whether the supplied string is a valid SOTA reference.
// The syntax is: AA/NN-CCC: Association/Name-3-digit numeric Code (e.g. G/CE-001).
func ValidateSota(inputStr string) (ref, errorMsg string) {
	inputStr = strings.ToUpper(strings.TrimSpace(inputStr))
	wrongInputStr := "*" + inputStr
	if validSotaRegexp.MatchString(inputStr) {
		return inputStr, ""
	}
	errorMsg = "[" + inputStr + "] is an invalid SOTA reference"
	return wrongInputStr, errorMsg
}

var validWwffRegexp = regexp.MustCompile(`^[\d]{0,1}[A-Z]{1,2}FF-[\d]{4}$`)

// ValidateWwff verifies whether the supplied string is a valid WWFF reference.
// The syntax is: AAFF-CCCC: AA = national prefix, CCCC = 4-digit numeric code (e.g. ONFF-0001).
func ValidateWwff(inputStr string) (ref, errorMsg string) {
	inputStr = strings.ToUpper(strings.TrimSpace(inputStr))
	wrongInputStr := "*" + inputStr
	if validWwffRegexp.MatchString(inputStr) {
		return inputStr, ""
	}
	errorMsg = "[" + inputStr + "] is an invalid WWFF reference"
	return wrongInputStr, errorMsg
}

// var validPotaRegexp = regexp.MustCompile(`^[\d]{0,1}[A-Z]{1,2}[\d]{0,1}-[\d]{4,5}$`)
var validPotaRegexp = regexp.MustCompile(`^[0-9A-Z]{1,5}-[\d]{4,5}$`)


// ValidatePota verifies whether the supplied string is a valid POTA reference. If valid, the error
// message is empty.
// The syntax is: AA-CCCCC: AA = national prefix, CCCCC = 4 or 5-digit numeric code (e.g. ON-00001).
// The national prefix is composed of letters and figures of at least 1 char and at most 5 char length
func ValidatePota(inputStr string) (ref, errorMsg string) {
	inputStr = strings.ToUpper(strings.TrimSpace(inputStr))
	wrongInputStr := "*" + inputStr
	if validPotaRegexp.MatchString(inputStr) {
		return inputStr, ""
	}
	errorMsg = "[" + inputStr + "] is an invalid POTA reference"
	return wrongInputStr, errorMsg
}

var validGridRegexp = regexp.MustCompile("(?i)^[a-z]{2}[0-9]{2}([a-z]{2})?$")

// ValidateGridLocator verifies that the supplied is a valid Maidenhead locator reference
// (either in 4 or 6 position). The returned grid case is normalized (first two letters
// in uppercase, last pair in lowercase). If the grid is not valid, the supicious string
// is prefixed with a * and an erroMsg is genrated.
func ValidateGridLocator(grid string) (processedGrid, errorMsg string) {
	if validGridRegexp.MatchString(grid) {
		var output strings.Builder
		for i, c := range grid {
			//The first pair of characters to be forced uppercase
			if (i == 0) || (i == 1) {
				output.WriteString(strings.ToUpper(string(c)))
			}
			//The second pair (numbers) are left alone
			if (i == 2) || (i == 3) {
				output.WriteString(string(c))
			}
			//The third pair of characters to be forced lowercase
			if (i == 4) || (i == 5) {
				output.WriteString(strings.ToLower(string(c)))
			}
		}
		return output.String(), ""
	}

	processedGrid = "*" + grid
	errorMsg = "[" + grid + "] is an invalid grid reference"
	return processedGrid, errorMsg
}

var validCallRegexp = regexp.MustCompile(`[\d]{0,1}[A-Z]{1,2}\d([A-Z]{1,4}|\d{3,3}|\d{1,3}[A-Z])[A-Z]{0,5}`)
var validPrefixRegexp = regexp.MustCompile(`\A[a-zA-Z0-9]{1,3}$`)

// ValidateCall verifies whether the supplied string is a valid callsign.
// prefix and suffix are not checked for validity
// If it is not valid, the supicious string is prefixed with a * and an erroMsg is genrated.
func ValidateCall(sign string) (call, errorMsg string) {
	sign = strings.ToUpper(strings.TrimSpace(sign))
	sp := strings.Split(sign, "/")
	wrongSign := "*" + sign
	switch len(sp) {
	case 1:
		if validCallRegexp.MatchString(sign) {
			return sign, ""
		}
		return wrongSign, "[" + sign + "] is an invalid call"
	case 2:
		// some ambiguity here we need to resolve, could be a prefix or a suffix
		if validCallRegexp.MatchString(sp[0]) {
			//Callisign with suffix (unchecked)
			return sign, ""
		}
		//else we are dealing with a prefixed Callsign
		//validate the part that should contain the call (sp[1])
		if !validCallRegexp.MatchString(sp[1]) {
			return wrongSign, "[" + sp[1] + "] is an invalid call"
		}
		//validate the prefix
		if !validPrefixRegexp.MatchString(sp[0]) {
			return wrongSign, "[" + sp[0] + "] is an invalid prefix"
		}
		return sign, ""
	case 3:
		//validate the part that should contain the call (sp[1])
		if !validCallRegexp.MatchString(sp[1]) {
			return wrongSign, "[" + sp[1] + "] is an invalid call"
		}
		//validate the prefix
		if !validPrefixRegexp.MatchString(sp[0]) {
			return wrongSign, "[" + sp[0] + "] is an invalid prefix"
		}
		//We don't check the suffix
		return sign, ""
	}
	return wrongSign, "[" + sign + "] is invalid: too many '/'"
}

var splitDateRegexp = regexp.MustCompile(`[-/ .]`)

// NormalizeDate takes what looks like a date and normalises it to "YYYY-MM-DD"
func NormalizeDate(inputStr string) (date, errorMsg string) {
	//Try to split the string
	s := splitDateRegexp.Split(inputStr, 4)

	//we should have three and only three elements
	if i := len(s); i != 3 {
		errorMsg = fmt.Sprintf("Bad date format: found %d elements while expecting 3.", i)
		return "*" + inputStr, errorMsg
	}

	//complete the numbers if shorter than expected ("20" for the first and "0" for the two next)
	year := s[0]
	if len(year) == 2 {
		year = "20" + year
	}
	//This test is not really necessary, but rather belt and suspenders
	if len(year) != 4 {
		errorMsg = "Bad date format: first part doesn't look like a year"
		return "*" + inputStr, errorMsg
	}

	month := s[1]
	if len(month) == 1 {
		month = "0" + month
	}
	if len(month) != 2 {
		errorMsg = "Bad date format: second part doesn't look like a month"
		return "*" + inputStr, errorMsg
	}

	day := s[2]
	if len(day) == 1 {
		day = "0" + day
	}
	if len(day) != 2 {
		errorMsg = "Bad date format: third element doesn't look like a day"
		return "*" + inputStr, errorMsg
	}

	//re-assemble the string with the correct delimiter
	date = year + "-" + month + "-" + day

	return date, ""
}

// ValidateDate verifies whether the string is a valid date (YYYY-MM-DD).
func ValidateDate(inputStr string) (ref, errorMsg string) {

	const RFC3339FullDate = "2006-01-02"

	inputStr = strings.ToUpper(strings.TrimSpace(inputStr))
	wrongInputStr := "*" + inputStr
	_, err := time.Parse(RFC3339FullDate, inputStr)

	if err == nil {
		return inputStr, ""
	}

	return wrongInputStr, fmt.Sprint(err)
}

// IncrementDate will increment the supplied date by the specified increment. It returns the new date.
func IncrementDate(date string, increment int) (newdate string, err string) {
	if date == "" {
		return "", "No date to increment"
	}
	if increment < 1 {
		return "*" + date, "Invalid day increment, expecting greater or equal to 1"
	}
	if 10 < increment {
		return "*" + date, "Invalid day increment, expecting smaller or equal to 10"
	}

	const RFC3339FullDate = "2006-01-02"
	convertedTime, timeErr := time.Parse(RFC3339FullDate, date)
	if timeErr != nil {
		return "*" + date, "(Internal error) error " + fmt.Sprint(timeErr)
	}
	// the number of days specified in increment
	newDate := convertedTime.AddDate(0, 0, increment)

	return newDate.Format(RFC3339FullDate), ""
}

// IsBand retuns true if the passed input string is a valid string
func IsBand(inputStr string) (result bool, lowerLimit, upperLimit float64, altBandName string) {
	switch strings.ToLower(inputStr) {
	case "2190m":
		return true, 0.1357, 0.1378, "VLF"
	case "630m":
		return true, 0.472, 0.479, "VLF"
	case "560m":
		return true, 0.501, 0.504, "VLF"
	case "160m":
		return true, 1.8, 2.0, "1.8MHz"
	case "80m":
		return true, 3.5, 4.0, "3.5MHz"
	case "60m":
		return true, 5.06, 5.45, "5MHz"
	case "40m":
		return true, 7.0, 7.3, "7MHz"
	case "30m":
		return true, 10.1, 10.15, "10MHz"
	case "20m":
		return true, 14.0, 14.35, "14MHz"
	case "17m":
		return true, 18.068, 18.168, "18MHz"
	case "15m":
		return true, 21.0, 21.45, "21MHz"
	case "12m":
		return true, 24.890, 24.99, "24MHz"
	case "10m":
		return true, 28.0, 29.7, "28MHz"
	case "6m":
		return true, 50, 54, "50MHz"
	case "4m":
		return true, 70, 71, "70MHz"
	case "2m":
		return true, 144, 148, "144MHz"
	case "1.25m":
		return true, 222, 225, "222MHz"
	case "70cm":
		return true, 420, 450, "432MHz"
	case "33cm":
		return true, 902, 928, "900MHz"
	case "23cm":
		return true, 1240, 1300, "1240MHz"
	case "13cm":
		return true, 2300, 2450, "2.3GHz"
	case "9cm":
		return true, 3300, 3500, "3.4GHz"
	case "6cm":
		return true, 5650, 5925, "5.6GHz"
	case "3cm":
		return true, 10000, 10500, "10GHz"
	case "1.25cm":
		return true, 24000, 24250, "24GHz"
	case "6mm":
		return true, 47000, 47200, "Microwave"
	case "4mm":
		return true, 75500, 81000, "Microwave"
	case "2.5mm":
		return true, 119980, 120020, "Microwave"
	case "2mm":
		return true, 142000, 149000, "Microwave"
	case "1mm":
		return true, 241000, 250000, "Microwave"
	}
	return false, 0, 0, ""
}

func getDefaultReport(mode string) (modeType, defaultReport string) {
	modeType = ""
	defaultReport = ""

	switch mode {
	case "SSB", "AM", "FM":
		modeType = "PHONE"
		defaultReport = "59"
	case "CW", "RTTY", "PSK":
		modeType = "CW"
		defaultReport = "599"
	case "JT65", "JT9", "JT6M", "JT4", "JT44", "FSK441", "FT8", "ISCAT", "MSK144", "QRA64", "T10", "WSPR":
		modeType = "DIGITAL"
		defaultReport = "-10"
	}
	return modeType, defaultReport
}
