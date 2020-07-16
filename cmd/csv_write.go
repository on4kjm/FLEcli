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
	"fmt"
	"strings"
	"time"
)

// outputAdif generates and writes data in ADIF format
func outputCsv(outputFile string, fullLog []LogLine) {

	//TODO: validate input data for combination

	//convert the log data to an in-memory ADIF file
	csvData := buildCsv(fullLog)

	//write to a file (re-using function defined to write adif file)
	writeFile(outputFile, csvData)
}

// buildAdif creates the adif file in memory ready to be printed
func buildCsv(fullLog []LogLine) (csvList []string) {

	// V2,ON4KJM/P,ON/ON-001,24/05/20,1310,14MHz,CW,S57LC

	for _, logLine := range fullLog {
		var csvLine strings.Builder
		csvLine.WriteString("V2,")
		csvLine.WriteString(fmt.Sprintf("%s", logLine.MyCall))
		csvLine.WriteString(fmt.Sprintf(",%s", logLine.MySOTA))
		csvLine.WriteString(fmt.Sprintf(",%s", csvDate(logLine.Date)))
		csvLine.WriteString(fmt.Sprintf(",%s", logLine.Time))
		//TODO: Should we test the result
		_, _, _, sotaBand := IsBand(logLine.Band)
		csvLine.WriteString(fmt.Sprintf(",%s",sotaBand ))
		csvLine.WriteString(fmt.Sprintf(",%s",logLine.Mode))
		csvLine.WriteString(fmt.Sprintf(",%s", logLine.Call))
		if logLine.SOTA != "" {
			csvLine.WriteString(fmt.Sprintf(",%s", logLine.SOTA))		
		} else {
			if logLine.Comment != "" {
				csvLine.WriteString(",")	
			}
		}
		if logLine.Comment != "" {
			csvLine.WriteString(fmt.Sprintf(",%s", logLine.Comment))
		}

		csvList = append(csvList, csvLine.String())
	}
	return csvList
}


//adifDate converts a date in YYYY-MM-DD format to YYYYMMDD
func csvDate(inputDate string) (outputDate string) {
	const FLEdateFormat = "2006-01-02"
	date, err := time.Parse(FLEdateFormat, inputDate)
	//error should never happen
	if err != nil {
		panic(err)
	}

	const CSVdateFormat = "02/01/06"
	return date.Format(CSVdateFormat)
}
