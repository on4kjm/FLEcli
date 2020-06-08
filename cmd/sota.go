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