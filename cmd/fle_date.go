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
	"strings"
	"fmt"
	"time"
)

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