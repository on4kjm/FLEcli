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

//Documentation of the SOTA CSV format: https://www.sotadata.org.uk/en/upload/activator/csv/info

import (
	"fmt"
)

//ProcessCsvCommand loads an FLE input to produce a SOTA CSV
func ProcessCsvCommand(inputFilename, outputCsvFilename string, isInterpolateTime, isOverwriteCsv bool) {

	if verifiedOutputFilename, filenameWasOK := buildOutputFilename(outputCsvFilename, inputFilename, isOverwriteCsv, ".csv"); filenameWasOK == true {
		if loadedLogFile, isLoadedOK := LoadFile(inputFilename, isInterpolateTime); isLoadedOK == true {
			if validateDataForSotaCsv(loadedLogFile) {
				outputCsv(verifiedOutputFilename, loadedLogFile)
			} else {
				//TODO: failed to validate
			}
		} else {
			//TODO: Parsing errors, aborting....
		}
	}
}

func validateDataForSotaCsv(loadedLogFile []LogLine) bool {
	if len(loadedLogFile) == 0 {
		fmt.Println("No useful data read. Aborting...")
		return false
	}

	for i := 0; i < len(loadedLogFile); i++ {
		if loadedLogFile[0].MySOTA == "" {
			fmt.Println("Missing MY-SOTA reference. Aborting...")
			return false
		}
	}
	// csvLine.WriteString(fmt.Sprintf("%s", logLine.MyCall))
	// csvLine.WriteString(fmt.Sprintf(",%s", logLine.MySOTA))
	// csvLine.WriteString(fmt.Sprintf(",%s", csvDate(logLine.Date)))
	// csvLine.WriteString(fmt.Sprintf(",%s", logLine.Time))
	// csvLine.WriteString(fmt.Sprintf(",%s", sotaBand))
	// csvLine.WriteString(fmt.Sprintf(",%s", logLine.Mode))
	// csvLine.WriteString(fmt.Sprintf(",%s", logLine.Call))

	//check if we have the necessary information for the type

	return true
}
