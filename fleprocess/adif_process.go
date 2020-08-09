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
)

//ProcessAdifCommand FIXME
func ProcessAdifCommand(inputFilename, outputFilename string, isInterpolateTime, isWWFFcli, isSOTAcli, isOverwrite bool) {

	verifiedOutputFilename, filenameWasOK := buildOutputFilename(outputFilename, inputFilename, isOverwrite, ".adi")

	// if the output file could not be parsed correctly do noting
	if filenameWasOK {
		loadedLogFile, isLoadedOK := LoadFile(inputFilename, isInterpolateTime)

		if isLoadedOK {
			if len(loadedLogFile) == 0 {
				fmt.Println("No useful data read. Aborting...")
				return
			}

			//TODO: There are more tests required here
			//check if we have the necessary information for the type
			if isWWFFcli {
				if loadedLogFile[0].MyWWFF == "" {
					fmt.Println("Missing MY-WWFF reference. Aborting...")
					return
				}
				if loadedLogFile[0].Operator == "" {
					fmt.Println("Missing Operator. Aborting...")
					return
				}
			}
			if isSOTAcli {
				if loadedLogFile[0].MySOTA == "" {
					fmt.Println("Missing MY-SOTA reference. Aborting...")
					return
				}
			}

			OutputAdif(verifiedOutputFilename, loadedLogFile, isWWFFcli, isSOTAcli)
		}
	}
}
