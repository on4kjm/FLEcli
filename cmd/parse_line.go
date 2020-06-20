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
	"regexp"
	"strings"
	"fmt"
)

// LogLine is used to store all the data of a single log line
type LogLine struct {
    Date      string
	MyCall    string
	Operator  string
	MyWWFF    string
	MySOTA    string
	QslMsg    string
	Nickname  string
	Mode      string
	Band      string
	Frequency string
	Time      string
	Call      string
	Comment   string
	QSLmsg    string
	OMname    string
	GridLoc   string
	RSTsent   string
	RSTrcvd   string
}

var regexpIsBand = regexp.MustCompile("m$")
var regexpIsFullTime = regexp.MustCompile("^[0-2]{1}[0-9]{3}$")
var regexpIsTimePart = regexp.MustCompile("^[0-5]{1}[0-9]{1}$|^[1-9]{1}$")
var regexpIsOMname = regexp.MustCompile("^@")
var regexpIsGridLoc = regexp.MustCompile("^#")

// ParseLine cuts a FLE line into useful bits
func ParseLine(inputStr string, previousLine LogLine) (logLine LogLine, errorMsg string){
	//TODO: input null protection?

	//Flag telling that we are processing data to the right of the callsign
	isRightOfCall := false

	//TODO: Make something more intelligent
	//TODO: What happens if we have partial lines
	previousLine.Call = ""
	previousLine.RSTsent = ""
	previousLine.RSTrcvd = ""
	logLine = previousLine

	//TODO: what happens when we have <> or when there are multiple comments
	//TODO: Refactor this! it is ugly
	comment,inputStr := getBraketedData(inputStr, COMMENT)
	if comment != "" {
		logLine.Comment = comment
		fmt.Println("Cleaned input string: ", inputStr)
	}

	QSLmsg,inputStr := getBraketedData(inputStr, QSL)
	if QSLmsg != "" {
		logLine.QSLmsg = QSLmsg
		fmt.Println("Cleaned input string: ", inputStr)
	}


	elements := strings.Fields(inputStr)

	for _, element := range elements {

		//  Is it a mode?
		if lookupMode( strings.ToUpper(element)) {
			logLine.Mode = strings.ToUpper(element)
			// Set the default RST depending of the mode
			if (logLine.RSTsent == "") || (logLine.RSTrcvd == "") {
				switch logLine.Mode {
				case "SSB", "AM", "FM" :
					logLine.RSTsent = "59"
					logLine.RSTrcvd = "59"
				case "CW", "RTTY", "PSK":
					logLine.RSTsent = "599"
					logLine.RSTrcvd = "599"
				case "JT65", "JT9", "JT6M", "JT4", "JT44", "FSK441", "FT8", "ISCAT", "MSK144", "QRA64", "T10", "WSPR" :
					logLine.RSTsent = "-10"
					logLine.RSTrcvd = "-10"				
				}

			} else {
				errorMsg = errorMsg + "Double definitiion of RST"
			}
			continue
		}

		// Is it a band?
		if regexpIsBand.MatchString(element) {
			logLine.Band = element
			continue
		}

		// Is it a call sign ?
		if validCallRegexp.MatchString(strings.ToUpper(element)) {
			callErrorMsg := ""
			logLine.Call, callErrorMsg = ValidateCall(element)
			errorMsg = errorMsg + callErrorMsg
			isRightOfCall = true
			continue
		}

		// Is it a "full" time ?
		if isRightOfCall == false {
			if regexpIsFullTime.MatchString(element) {
				logLine.Time = element
				continue
			}

			// Is it a partial time ?
			if regexpIsTimePart.MatchString(element) {
				if logLine.Time == "" {
					logLine.Time = element
				} else {
					goodPart := logLine.Time[:len(logLine.Time)-len(element)]
					logLine.Time = goodPart + element
				}
				continue
			}
		}
		

		// Is it the OM's name (starting with "@")
		if regexpIsOMname.MatchString(element) {
			logLine.OMname = strings.TrimLeft(element, "@") 
			continue
		}


		// Is it the Grid Locator (starting with "#")
		if regexpIsGridLoc.MatchString(element) {
			logLine.GridLoc = strings.TrimLeft(element, "#") 
			continue
		}

		if isRightOfCall {
			//This is probably a RST
			//TODO: is it a number (or a data report)
			//TODO: it is sent or rcvd
		}

		//If we come here, we could not make sense of what we found
		errorMsg = errorMsg + "Unable to parse " + element + " "

	}

	fmt.Println(elements, len(elements))

	fmt.Println("\n", SprintLogRecord(logLine))

	return logLine, errorMsg
}

// SprintLogRecord outputs the content of a logline
func SprintLogRecord(logLine LogLine) (output string){
	output = ""
	output = output + "Date      " + logLine.Date + "\n"
	output = output + "MyCall    " + logLine.MyCall + "\n"
	output = output + "Operator  " + logLine.Operator + "\n"
	output = output + "MyWWFF    " + logLine.MyWWFF + "\n"
	output = output + "MySOTA    " + logLine.MySOTA + "\n"
	output = output + "QslMsg    " + logLine.QslMsg + "\n"
	output = output + "Nickname  " + logLine.Nickname + "\n"
	output = output + "Mode      " + logLine.Mode + "\n"
	output = output + "Band      " + logLine.Band + "\n"
	output = output + "Frequency " + logLine.Frequency + "\n"
	output = output + "Time      " + logLine.Time + "\n"
	output = output + "Call      " + logLine.Call + "\n"
	output = output + "Comment   " + logLine.Comment + "\n"
	output = output + "QSLmsg    " + logLine.QSLmsg + "\n"
	output = output + "OMname    " + logLine.OMname + "\n"
	output = output + "GridLoc   " + logLine.GridLoc + "\n"
	output = output + "RSTsent   " + logLine.RSTsent + "\n"
	output = output + "RSTrcvd   " + logLine.RSTrcvd + "\n"

	return output
}



func lookupMode(lookup string) bool {
	switch lookup {
	case
		"CW", 
		"SSB", 
		"AM",
		"FM",
		"RTTY", 
		"FT8",
		"PSK",
		"JT65",
		"JT9",
		"FT4",
		"JS8",
		"ARDOP",
		"ATV",
		"C4FM",
		"CHIP",
		"CLO",
		"CONTESTI",
		"DIGITALVOICE",
		"DOMINO",
		"DSTAR",
		"FAX",
		"FSK441",
		"HELL",
		"ISCAT",
		"JT4",
		"JT6M",
		"JT44",
		"MFSK",
		"MSK144",
		"MT63",
		"OLIVIA",
		"OPERA",
		"PAC",
		"PAX",
		"PKT",
		"PSK2K",
		"Q15",
		"QRA64",
		"ROS",
		"RTTYM",
		"SSTV",
		"T10",
		"THOR",
		"THRB",
		"TOR",
		"V4",
		"VOI",
		"WINMOR",
		"WSPR":	
		return true
	}
	return false
}



// func lookupBand(lookup string) bool {
// 	switch lookup {
// 	case
// 		"900898296857",
// 		"900898302052",
// 		"900898296492",
// 		"900898296850",
// 		"900898296703",
// 		"900898296633",
// 		"900898296613",
// 		"900898296615",
// 		"900898296620",
// 		"900898296636":
// 		return true
// 	}
// 	return false
// }

// 2190m	.1357	.1378
// 630m	.472	.479
// 560m	.501	.504
// 160m	1.8	2.0
// 80m	3.5	4.0
// 60m	5.06	5.45
// 40m	7.0	7.3
// 30m	10.1	10.15
// 20m	14.0	14.35
// 17m	18.068	18.168
// 15m	21.0	21.45
// 12m	24.890	24.99
// 10m	28.0	29.7
// 6m	50	54
// 4m	70	71
// 2m	144	148
// 1.25m	222	225
// 70cm	420	450
// 33cm	902	928
// 23cm	1240	1300
// 13cm	2300	2450
// 9cm	3300	3500
// 6cm	5650	5925
// 3cm	10000	10500
// 1.25cm	24000	24250
// 6mm	47000	47200
// 4mm	75500	81000
// 2.5mm	119980	120020
// 2mm	142000	149000
// 1mm	241000	250000