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
	"strings"
)

//ProcessCsvCommand loads an FLE input to produce a SOTA CSV
func ProcessCsvCommand(inputFilename, outputCsvFilename string, isInterpolateTime, isOverwriteCsv bool) error {

	if verifiedOutputFilename, filenameWasOK := buildOutputFilename(outputCsvFilename, inputFilename, isOverwriteCsv, ".csv"); filenameWasOK == true {
		if loadedLogFile, isLoadedOK := LoadFile(inputFilename, isInterpolateTime); isLoadedOK == true {
			if err := validateDataForSotaCsv(loadedLogFile); err != nil {
				return err
			}
			outputCsv(verifiedOutputFilename, loadedLogFile)
			return nil
		}
		return fmt.Errorf("There were input file parsing errors. Could not generate CSV file")
	}
	//TODO: we need something more explicit here
	return fmt.Errorf("Failed to compute or validate output file name")
}

//TODO: change return boolean to full err
func validateDataForSotaCsv(loadedLogFile []LogLine) error {
	if len(loadedLogFile) == 0 {
		return fmt.Errorf("No useful data read")
	}

	//MySOTA is a header value. If missing on the first line, it will be missing at every line
	if loadedLogFile[0].MySOTA == "" {
		return fmt.Errorf("Missing MY-SOTA reference")
	}

	var errorsBuffer strings.Builder
	//We accumulate the errors messages
	for i := 0; i < len(loadedLogFile); i++ {

		var errorLocation string
		if loadedLogFile[i].Time == "" {
			errorLocation = fmt.Sprintf("for log entry #%d", i+1)
		} else {
			errorLocation = fmt.Sprintf("for log entry at %s (#%d)", loadedLogFile[i].Time, i+1)
		}
		if loadedLogFile[i].MyCall == "" {
			errorsBuffer.WriteString(fmt.Sprintf("Missing MyCall %s\n", errorLocation))
		}
		if loadedLogFile[i].Date == "" {
			errorsBuffer.WriteString(fmt.Sprintf("Missing date %s\n", errorLocation))
		}
		if loadedLogFile[i].Band == "" {
			errorsBuffer.WriteString(fmt.Sprintf("Missing band %s\n", errorLocation))
		}
		if loadedLogFile[i].Mode == "" {
			errorsBuffer.WriteString(fmt.Sprintf("Missing mode %s\n", errorLocation))
		}
		if loadedLogFile[i].Call == "" {
			errorsBuffer.WriteString(fmt.Sprintf("Missing call %s\n", errorLocation))
		}
	}
	if errorsBuffer.String() != "" {
		return fmt.Errorf(errorsBuffer.String())
	}
	return nil
}
