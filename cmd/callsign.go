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
)


var validCallRegexp = regexp.MustCompile(`[\d]{0,1}[A-Z]{1,2}\d([A-Z]{1,4}|\d{3,3}|\d{1,3}[A-Z])[A-Z]{0,5}`)
var validPrefixRegexp = regexp.MustCompile(`\A\d?[a-zA-Z]{1,2}$`)

// ValidateCall veriffies whether the supplied string is a valid callsign.
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