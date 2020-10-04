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
func ProcessCsvCommand(inputFilename, outputFilename string, isInterpolateTime, isOverwriteCsv bool) error {

	//Validate of build the output filenaem
	var verifiedOutputFilename string
	var err error

	if verifiedOutputFilename, err = buildOutputFilename(outputFilename, inputFilename, isOverwriteCsv, ".csv"); err != nil {
		return err
	}

	//Load the input file
	var loadedLogFile []LogLine
	var isLoadedOK bool

	if loadedLogFile, isLoadedOK = LoadFile(inputFilename, isInterpolateTime); isLoadedOK == false {
		return fmt.Errorf("There were input file parsing errors. Could not generate CSV file")
	}

	//Check if we have all the necessary data
	if err := validateDataForSotaCsv(loadedLogFile); err != nil {
		return err
	}

	outputCsv(verifiedOutputFilename, loadedLogFile)

	return nil

}

//validateDataForSotaCsv checks whether all the requiered data is present in the supplied data
func validateDataForSotaCsv(loadedLogFile []LogLine) error {
	if len(loadedLogFile) == 0 {
		return fmt.Errorf("No QSO found")
	}

	isNoMySota := false
	//MySOTA and MyCall are header values. If missing on the first line, it will be missing at every line
	if loadedLogFile[0].MySOTA == "" {
		//if not set, we might be dealing with a chaser log
		isNoMySota = true
	}
	if loadedLogFile[0].MyCall == "" {
		return fmt.Errorf("Missing MyCall")
	}

	var errorsBuffer strings.Builder
	//We accumulate the errors messages
	for i := 0; i < len(loadedLogFile); i++ {

		//Compute the error location for a meaning full error
		var errorLocation string
		if loadedLogFile[i].Time == "" {
			errorLocation = fmt.Sprintf("for log entry #%d", i+1)
		} else {
			errorLocation = fmt.Sprintf("for log entry at %s (#%d)", loadedLogFile[i].Time, i+1)
		}

		if loadedLogFile[i].Date == "" {
			if errorsBuffer.String() != "" {
				errorsBuffer.WriteString(fmt.Sprintf(", "))
			}
			errorsBuffer.WriteString(fmt.Sprintf("missing date %s", errorLocation))
		}
		if loadedLogFile[i].Band == "" {
			if errorsBuffer.String() != "" {
				errorsBuffer.WriteString(fmt.Sprintf(", "))
			}
			errorsBuffer.WriteString(fmt.Sprintf("missing band %s", errorLocation))
		}
		if loadedLogFile[i].Mode == "" {
			if errorsBuffer.String() != "" {
				errorsBuffer.WriteString(fmt.Sprintf(", "))
			}
			errorsBuffer.WriteString(fmt.Sprintf("missing mode %s", errorLocation))
		}
		if loadedLogFile[i].Call == "" {
			if errorsBuffer.String() != "" {
				errorsBuffer.WriteString(fmt.Sprintf(", "))
			}
			errorsBuffer.WriteString(fmt.Sprintf("missing call %s", errorLocation))
		}
		if loadedLogFile[i].Time == "" {
			if errorsBuffer.String() != "" {
				errorsBuffer.WriteString(fmt.Sprintf(", "))
			}
			errorsBuffer.WriteString(fmt.Sprintf("missing QSO time %s", errorLocation))
		}
		//FIXME: if isNoMySota and MySota defined means that it was defined later in the log file
		if (isNoMySota && loadedLogFile[i].SOTA == "") {
			if errorsBuffer.String() != "" {
				errorsBuffer.WriteString(fmt.Sprintf(", "))
			}
			errorsBuffer.WriteString(fmt.Sprintf("missing SOTA reference while attempting to process chaser log %s", errorLocation))	
		}
	}
	if errorsBuffer.String() != "" {
		return fmt.Errorf(errorsBuffer.String())
	}
	return nil
}
