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
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

//LoadFile FIXME
//returns nill if failure to process
func LoadFile(inputFilename string, isInterpolateTime bool) (filleFullLog []LogLine, isProcessedOK bool) {
	file, err := os.Open(inputFilename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	file.Close()

	regexpLineComment := regexp.MustCompile("^[[:blank:]]*#")
	regexpOnlySpaces := regexp.MustCompile("^\\s+$")
	regexpSingleMultiLineComment := regexp.MustCompile("^[[:blank:]]*{.+}$")
	regexpStartMultiLineComment := regexp.MustCompile("^[[:blank:]]*{")
	regexpEndMultiLineComment := regexp.MustCompile("}$")
	regexpHeaderMyCall := regexp.MustCompile("(?i)^mycall ")
	regexpHeaderOperator := regexp.MustCompile("(?i)^operator ")
	regexpHeaderMyWwff := regexp.MustCompile("(?i)^mywwff ")
	regexpHeaderMySota := regexp.MustCompile("(?i)^mysota ")
	regexpHeaderMyGrid := regexp.MustCompile("(?i)^mygrid ")
	regexpHeaderQslMsg := regexp.MustCompile("(?i)^qslmsg ")
	regexpHeaderNickname := regexp.MustCompile("(?i)^nickname ")

	headerMyCall := ""
	headerOperator := ""
	headerMyWWFF := ""
	headerMySOTA := ""
	headerMyGrid := ""
	headerQslMsg := ""
	headerNickname := ""
	//headerDate := ""
	lineCount := 0

	wrkTimeBlock := InferTimeBlock{}
	missingTimeBlockList := []InferTimeBlock{}

	var isInMultiLine = false
	var cleanedInput []string
	var errorLog []string

	var previousLogLine LogLine
	fullLog := []LogLine{}

	//Loop through all the stored lined
	for _, eachline := range txtlines {
		lineCount++

		// ****
		// ** Lets do some house keeping first by droping the unecessary lines
		// ****

		//Skip the line if it starts with "#"
		if regexpLineComment.MatchString(eachline) {
			continue
		}
		//Skip if line is empty or blank
		if (len(eachline) == 0) || (regexpOnlySpaces.MatchString(eachline)) {
			continue
		}

		// Process multi-line comments
		if regexpStartMultiLineComment.MatchString(eachline) {
			//Single-line "multi-line" coment
			if regexpSingleMultiLineComment.MatchString(eachline) {
				continue
			}
			isInMultiLine = true
			continue
		}
		if isInMultiLine {
			if regexpEndMultiLineComment.MatchString(eachline) {
				isInMultiLine = false
			}
			continue
		}

		// ****
		// ** Process the Header block
		// ****

		//My Call
		if regexpHeaderMyCall.MatchString(eachline) {
			//Attempt to redefine value
			if headerMyCall != "" {
				errorLog = append(errorLog, fmt.Sprintf("Attempt to redefine MyCall at line %d", lineCount))
				continue
			}
			errorMsg := ""
			myCallList := regexpHeaderMyCall.Split(eachline, -1)
			if len(strings.TrimSpace(myCallList[1])) > 0 {
				headerMyCall, errorMsg = ValidateCall(strings.TrimSpace(myCallList[1]))
				cleanedInput = append(cleanedInput, fmt.Sprintf("My call: %s", headerMyCall))
				if len(errorMsg) != 0 {
					errorLog = append(errorLog, fmt.Sprintf("Invalid myCall at line %d: %s (%s)", lineCount, myCallList[1], errorMsg))
				}
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		//Operator
		if regexpHeaderOperator.MatchString(eachline) {
			//Attempt to redefine value
			if headerOperator != "" {
				errorLog = append(errorLog, fmt.Sprintf("Attempt to redefine Operator at line %d", lineCount))
				continue
			}
			errorMsg := ""
			myOperatorList := regexpHeaderOperator.Split(eachline, -1)
			if len(strings.TrimSpace(myOperatorList[1])) > 0 {
				headerOperator, errorMsg = ValidateCall(strings.TrimSpace(myOperatorList[1]))
				cleanedInput = append(cleanedInput, fmt.Sprintf("Operator: %s", headerOperator))
				if len(errorMsg) != 0 {
					errorLog = append(errorLog, fmt.Sprintf("Invalid Operator at line %d: %s (%s)", lineCount, myOperatorList[1], errorMsg))
				}
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		// My WWFF
		if regexpHeaderMyWwff.MatchString(eachline) {
			//Attempt to redefine value
			if headerMyWWFF != "" {
				errorLog = append(errorLog, fmt.Sprintf("Attempt to redefine MyWWFF at line %d", lineCount))
				continue
			}
			errorMsg := ""
			myWwffList := regexpHeaderMyWwff.Split(eachline, -1)
			if len(strings.TrimSpace(myWwffList[1])) > 0 {
				headerMyWWFF, errorMsg = ValidateWwff(strings.TrimSpace(myWwffList[1]))
				cleanedInput = append(cleanedInput, fmt.Sprintf("My WWFF: %s", headerMyWWFF))
				if len(errorMsg) != 0 {
					errorLog = append(errorLog, fmt.Sprintf("Invalid \"My WWFF\" at line %d: %s (%s)", lineCount, myWwffList[1], errorMsg))
				}
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		//My Sota
		if regexpHeaderMySota.MatchString(eachline) {
			//Attempt to redefine value
			if headerMySOTA != "" {
				errorLog = append(errorLog, fmt.Sprintf("Attempt to redefine MySOTA at line %d", lineCount))
				continue
			}
			errorMsg := ""
			mySotaList := regexpHeaderMySota.Split(eachline, -1)
			if len(strings.TrimSpace(mySotaList[1])) > 0 {
				headerMySOTA, errorMsg = ValidateSota(strings.TrimSpace(mySotaList[1]))
				cleanedInput = append(cleanedInput, fmt.Sprintf("My Sota: %s", headerMySOTA))
				if len(errorMsg) != 0 {
					errorLog = append(errorLog, fmt.Sprintf("Invalid \"My SOTA\" at line %d: %s (%s)", lineCount, mySotaList[1], errorMsg))
				}
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		//My Grid
		if regexpHeaderMyGrid.MatchString(eachline) {
			//Attempt to redefine value
			if headerMyGrid != "" {
				errorLog = append(errorLog, fmt.Sprintf("Attempt to redefine MyGrid at line %d", lineCount))
				continue
			}
			errorMsg := ""
			myGridList := regexpHeaderMyGrid.Split(eachline, -1)
			if len(strings.TrimSpace(myGridList[1])) > 0 {
				headerMyGrid, errorMsg = ValidateGridLocator(strings.TrimSpace(myGridList[1]))
				cleanedInput = append(cleanedInput, fmt.Sprintf("My Grid: %s", headerMyGrid))
				if len(errorMsg) != 0 {
					errorLog = append(errorLog, fmt.Sprintf("Invalid \"My Grid\" at line %d: %s (%s)", lineCount, myGridList[1], errorMsg))
				}
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		//QSL Message
		if regexpHeaderQslMsg.MatchString(eachline) {
			myQslMsgList := regexpHeaderQslMsg.Split(eachline, -1)
			if len(myQslMsgList[1]) > 0 {
				headerQslMsg = myQslMsgList[1]
				cleanedInput = append(cleanedInput, fmt.Sprintf("QSL Message: %s", headerQslMsg))
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		//Nickname
		if regexpHeaderNickname.MatchString(eachline) {
			//Attempt to redefine value
			if headerNickname != "" {
				errorLog = append(errorLog, fmt.Sprintf("Attempt to redefine eQSL Nickname at line %d", lineCount))
				continue
			}
			myNicknameList := regexpHeaderNickname.Split(eachline, -1)
			if len(strings.TrimSpace(myNicknameList[1])) > 0 {
				headerNickname = strings.TrimSpace(myNicknameList[1])
				cleanedInput = append(cleanedInput, fmt.Sprintf("eQSL Nickmane: %s", headerNickname))
			}
			//If there is no data after the marker, we just skip the data.
			continue
		}

		// ****
		// ** Process the data block
		// ****

		// Load the header values in the previousLogLine
		previousLogLine.MyCall = headerMyCall
		previousLogLine.Operator = headerOperator
		previousLogLine.MyWWFF = headerMyWWFF
		previousLogLine.MySOTA = headerMySOTA
		previousLogLine.MyGrid = headerMyGrid
		previousLogLine.QSLmsg = headerQslMsg //previousLogLine.QslMsg is redundant
		previousLogLine.Nickname = headerNickname
		//previousLogLine.Date = headerDate

		//parse a line
		logline, errorLine := ParseLine(eachline, previousLogLine)

		//we have a valid line (contains a call)
		if logline.Call != "" {
			fullLog = append(fullLog, logline)

			//store time inference data
			if isInterpolateTime {
				var isEndOfGap bool
				if isEndOfGap, err = wrkTimeBlock.storeTimeGap(logline, len(fullLog)); err != nil {
					fmt.Println("\nProcessing errors:")
					for _, errorLogLine := range errorLog {
						fmt.Println(errorLogLine)
					}
					log.Println("Fatal error: ", err)
					os.Exit(1)
				}
				//If we reached the end of the time gap, we make the necessary checks and make our gap calculation
				if isEndOfGap {
					if err := wrkTimeBlock.finalizeTimeGap(); err != nil {
						//If an error occured it is a fatal error
						fmt.Println("\nProcessing errors:")
						for _, errorLogLine := range errorLog {
							fmt.Println(errorLogLine)
						}
						log.Println("Fatal error: ", err)
						os.Exit(1)
					}

					//add it to the gap collection
					missingTimeBlockList = append(missingTimeBlockList, wrkTimeBlock)

					//create a new block
					wrkTimeBlock = InferTimeBlock{}

					//Store this record in the new block as a new gap might be following
					//no error or endOfGap processing as it has already been successfully processed
					wrkTimeBlock.storeTimeGap(logline, len(fullLog))
				}
			}
		}

		//Store append the accumulated soft parsing errors into the global parsing error log file
		if errorLine != "" {
			errorLog = append(errorLog, fmt.Sprintf("Parsing error at line %d: %s ", lineCount, errorLine))
		}

		//store the current logline so that it can be used as a model when parsing the next line
		previousLogLine = logline

		//We go back to the top to process the next loaded log line (Continue not necessary here)
	}

	//***
	//*** We have done processing the log file, so let's post process it
	//***

	//if asked to infer the date, lets update the loaded logfile accordingly
	if isInterpolateTime {
		for _, timeBlock := range missingTimeBlockList {
			for i := 0; i < timeBlock.noTimeCount; i++ {
				position := timeBlock.logFilePosition + i
				pLogLine := &fullLog[position]

				// durationOffset := time.Second * time.Duration(timeBlock.deltatime*(i+1))
				durationOffset := timeBlock.deltatime * time.Duration(i+1)
				newTime := timeBlock.lastRecordedTime.Add(durationOffset)
				updatedTimeString := newTime.Format("1504")
				pLogLine.Time = updatedTimeString
			}
		}
	}

	displayLogSimple(fullLog)

	//Display parsing errors, if any
	if len(errorLog) != 0 {
		fmt.Println("\nProcessing errors:")
		for _, errorLogLine := range errorLog {
			fmt.Println(errorLogLine)
		}
		isProcessedOK = false
	} else {
		fmt.Println("\nSuccessfully parsed ", lineCount, " lines.")
		isProcessedOK = true
	}

	return fullLog, isProcessedOK

}

//displayLogSimple will print to stdout a simplified dump of a full log
func displayLogSimple(fullLog []LogLine) {
	firstLine := true
	for _, filledLogLine := range fullLog {
		if firstLine {
			fmt.Println(SprintHeaderValues(filledLogLine))
			fmt.Print(SprintColumnTitles())
			firstLine = false
		}
		fmt.Print(SprintLogInColumn(filledLogLine))
	}

}
