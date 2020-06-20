package cmd
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
	"regexp"
	"strings"
	"fmt"
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
		return wrongInputStr, "Invalid SOTA reference"
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
		return wrongInputStr, "Invalid WWFF reference"
}



var validCallRegexp = regexp.MustCompile(`[\d]{0,1}[A-Z]{1,2}\d([A-Z]{1,4}|\d{3,3}|\d{1,3}[A-Z])[A-Z]{0,5}`)
var validPrefixRegexp = regexp.MustCompile(`\A\d?[a-zA-Z]{1,2}$`)

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
		return wrongSign, "Invalid call"
	case 2:
		// some ambiguity here we need to resolve, could be a prefix or a suffix
		if validCallRegexp.MatchString(sp[0]) {
			//Callisign with suffix (unchecked)
			return sign, ""
		} 
		//else we are dealing with a prefixed Callsign
		//validate the part that should contain the call (sp[1]) 
		if !validCallRegexp.MatchString(sp[1]) {
			return wrongSign, "Invalid call"
		}
		//validate the prefix
		if !validPrefixRegexp.MatchString(sp[0]) {
			return wrongSign, "Invalid prefix"
		}
		return sign, ""
	case 3:
		//validate the part that should contain the call (sp[1]) 
		if !validCallRegexp.MatchString(sp[1]) {
			return wrongSign, "Invalid call"
		}
		//validate the prefix
		if !validPrefixRegexp.MatchString(sp[0]) {
			return wrongSign, "Invalid prefix"
		}
		//We don't check the suffix
		return sign, ""
	}
	return wrongSign, "Too many '/'"
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

//IsBand retuns true if the passed input string is a valid string
//TODO: return the frequencies
func IsBand(inputStr string) (result bool, lowerLimit, upperLimit float32) {
	switch strings.ToLower(inputStr) {
		case "2190m":	
			return true, 0.1357, 0.1378
		case "630m":	
			return true, 0.472,	0.479
		case "560m":	
			return true, 0.501,	0.504
		case "160m":	
			return true, 1.8, 2.0
		case "80m":		
			return true, 3.5, 4.0
		case "60m":		
			return true, 5.06, 5.45
		case "40m":		
			return true, 7.0, 7.3
		case "30m":		
			return true, 10.1, 10.15
		case "20m":		
			return true, 14.0, 14.35
		case "17m":		
			return true, 18.068, 18.168
		case "15m":		
			return true, 21.0, 21.45
		case "12m":		
			return true, 24.890, 24.99
		case "10m":		
			return true, 28.0, 29.7
		case "6m":		
			return true, 50, 54
		case "4m":		
			return true, 70, 71
		case "2m":		
			return true, 144, 148
		case "1.25m":	
			return true, 222, 225
		case "70cm":	
			return true, 420, 450
		case "33cm":	
			return true, 902, 928
		case "23cm":	
			return true, 1240, 1300
		case "13cm":	
			return true, 2300, 2450
		case "9cm":		
			return true, 3300, 3500
		case "6cm":		
			return true, 5650, 5925
		case "3cm":		
			return true, 10000, 10500
		case "1.25cm":	
			return true, 24000, 24250
		case "6mm":		
			return true, 47000, 47200
		case "4mm":		
			return true, 75500, 81000
		case "2.5mm":	
			return true, 119980, 120020
		case "2mm":		
			return true, 142000, 149000
		case "1mm":		
			return true, 241000, 250000
	}
	return false, 0, 0
}