package fleprocess

import (
	"bufio"
	"fmt"
	"os"
)

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

// writeFile writes the in-memory data (lines) to a file
func writeFile(outputFile string, adifData []string) {

	//TODO: check access rights
	f, err := os.Create(outputFile)
	checkFileError(err)

	defer f.Close()

	w := bufio.NewWriter(f)

	lineCount := 0
	for _, adifLine := range adifData {
		_, err := w.WriteString(adifLine + "\n")
		checkFileError(err)

		w.Flush()
		checkFileError(err)
		lineCount++
	}
	fmt.Printf("\nSuccessfully wrote %d lines to file \"%s\"\n", lineCount, outputFile)
}