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
	"strings"
)

//ProcessAdifCommand FIXME
func ProcessAdifCommand(inputFilename, outputFilename string, isInterpolateTime, isWWFFcli, isSOTAcli, isOverwrite bool) error {

	//Validate of build the output filenaem
	var verifiedOutputFilename string
	var err error

	if verifiedOutputFilename, err = buildOutputFilename(outputFilename, inputFilename, isOverwrite, ".adi"); err != nil {
		return err
	}

	//Load the input file
	var loadedLogFile []LogLine
	var isLoadedOK bool

	if loadedLogFile, isLoadedOK = LoadFile(inputFilename, isInterpolateTime); isLoadedOK == false {
		return fmt.Errorf("There were input file parsing errors. Could not generate ADIF file")
	}

	//Check if we have all the necessary data
	if err := validateDataforAdif(loadedLogFile, isWWFFcli, isSOTAcli); err != nil {
		return err
	}

	//Write the output file with the checked data
	OutputAdif(verifiedOutputFilename, loadedLogFile, isWWFFcli, isSOTAcli)

	//If we reached this point, everything was processed OK and the file generated
	return nil
}

//validateDataforAdif checks whether all the required data is present
//The details of the mandatory files can be found at http://wwff.co/rules-faq/confirming-and-sending-log/
func validateDataforAdif(loadedLogFile []LogLine, isWWFFcli, isSOTAcli bool) error {

	//do we have QSOs at all?
	if len(loadedLogFile) == 0 {
		return fmt.Errorf("No QSO found")
	}

	//MySOTA, MyWWFF and MyCall are header values. If missing on the first line, it will be missing at every line
	if loadedLogFile[0].MyCall == "" {
		return fmt.Errorf("Missing MyCall")
	}
	if isSOTAcli {
		if loadedLogFile[0].MySOTA == "" {
			return fmt.Errorf("Missing MY-SOTA reference")
		}
	}
	if isWWFFcli {
		if loadedLogFile[0].MyWWFF == "" {
			return fmt.Errorf("Missing MY-WWFF reference")
		}
		if loadedLogFile[0].Operator == "" {
			return fmt.Errorf("Missing Operator call sign")
		}
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
	}
	if errorsBuffer.String() != "" {
		return fmt.Errorf(errorsBuffer.String())
	}

	//If we reached here, all is ok
	return nil
}
