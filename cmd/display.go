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
)

// SprintLogRecord outputs the content of a logline
func SprintLogRecord(logLine LogLine) (output string) {
	output = ""
	output = output + "Date      " + logLine.Date + "\n"
	output = output + "MyCall    " + logLine.MyCall + "\n"
	output = output + "Operator  " + logLine.Operator + "\n"
	output = output + "MyWWFF    " + logLine.MyWWFF + "\n"
	output = output + "MySOTA    " + logLine.MySOTA + "\n"
	output = output + "QslMsg    " + logLine.QslMsg + "\n"
	output = output + "Nickname  " + logLine.Nickname + "\n"
	output = output + "Mode      " + logLine.Mode + "\n"
	output = output + "ModeType  " + logLine.ModeType + "\n"
	output = output + "Band      " + logLine.Band + "\n"
	output = output + "  Lower   " + fmt.Sprintf("%f", logLine.BandLowerLimit) + "\n"
	output = output + "  Upper   " + fmt.Sprintf("%f", logLine.BandUpperLimit) + "\n"
	output = output + "Frequency " + logLine.Frequency + "\n"
	output = output + "Time      " + logLine.Time + "\n"
	output = output + "Call      " + logLine.Call + "\n"
	output = output + "Comment   " + logLine.Comment + "\n"
	output = output + "QSLmsg    " + logLine.QSLmsg + "\n"
	output = output + "OMname    " + logLine.OMname + "\n"
	output = output + "GridLoc   " + logLine.GridLoc + "\n"
	output = output + "RSTsent   " + logLine.RSTsent + "\n"
	output = output + "RSTrcvd   " + logLine.RSTrcvd + "\n"
	output = output + "SOTA      " + logLine.SOTA + "\n"
	output = output + "WWFF      " + logLine.WWFF + "\n"

	return output
}

// SprintHeaderValues displays the header values
func SprintHeaderValues(logLine LogLine) (output string) {
	output = ""

	output = output + "MyCall    " + logLine.MyCall
	if logLine.Operator != "" {
		output = output + " (" + logLine.Operator + ")"
	}
	output = output + "\n"

	if logLine.MyWWFF != "" {
		output = output + "MyWWFF    " + logLine.MyWWFF + "\n"
	}

	if logLine.MySOTA != "" {
		output = output + "MySOTA    " + logLine.MySOTA + "\n"
	}

	return output
}

// Date, Time, band, mode, call, report sent, report rcvd, Notes
var logLineFormat = "%-10s %-4s %-4s %-4s %-10s %-4s %-4s %s \n"

// SprintColumnTitles displays the column titles for a log line
func SprintColumnTitles(logLine LogLine) (output string) {
	output = fmt.Sprintf(logLineFormat, "Date", "Time", "Band", "Mode", "Call", "Sent", "Rcvd", "Notes")
	output = output + fmt.Sprintf(logLineFormat, "----", "----", "----", "----", "----", "----", "----", "----")
	return output
}

// SprintLogInColumn displays the logLine in column mode
func SprintLogInColumn(logLine LogLine) (output string) {
	notes := ""
	if logLine.Frequency != "" {
		notes = notes + "QRG: " + logLine.Frequency + " "
	}
	if logLine.Comment != "" {
		notes = notes + "[" + logLine.Comment + "] "
	}
	if logLine.QSLmsg != "" {
		notes = notes + "[" + logLine.QSLmsg + "] "
	}
	if logLine.OMname != "" {
		notes = notes + logLine.OMname + " "
	}
	if logLine.GridLoc != "" {
		notes = notes + logLine.GridLoc + " "
	}
	if logLine.WWFF != "" {
		notes = notes + logLine.WWFF + " "
	}
	if logLine.SOTA != "" {
		notes = notes + logLine.SOTA + " "
	}

	output = fmt.Sprintf(logLineFormat, logLine.Date, logLine.Time, logLine.Band, logLine.Mode, logLine.Call, logLine.RSTsent, logLine.RSTrcvd, notes)

	return output
}
