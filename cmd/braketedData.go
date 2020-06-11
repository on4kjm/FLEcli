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
)

func getBraketedData(inputLine, braketType string) (text, cleanedLine string) {
	// Get substring between two strings.
	a := ""
	b := ""

	//TODO: refactor that as a switch statement to exclude non supported bracket types
	if braketType == "COMMENT" {
		a = "<"
		b = ">"
	} 
	if braketType == "QSL" {
		a = "["
		b = "]"
	} 	

    posFirst := strings.Index(inputLine, a)
    if posFirst == -1 {
        return "",inputLine
    }
    posLast := strings.Index(inputLine, b)
    if posLast == -1 {
        return "",inputLine
    }
    posFirstAdjusted := posFirst + 1
    if posFirstAdjusted >= posLast {
        return "",inputLine
    }
    return inputLine[posFirstAdjusted:posLast], inputLine
}