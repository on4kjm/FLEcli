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

// SprintLogRecord outputs the content of a logline
func SprintLogRecord(logLine LogLine) string {
	var output strings.Builder

	output.WriteString("Date      " + logLine.Date + "\n")
	output.WriteString("MyCall    " + logLine.MyCall + "\n")
	output.WriteString("Operator  " + logLine.Operator + "\n")
	output.WriteString("MyWWFF    " + logLine.MyWWFF + "\n")
	output.WriteString("MySOTA    " + logLine.MySOTA + "\n")
	output.WriteString("MyPOTA    " + logLine.MyPOTA + "\n")
	output.WriteString("MyGrid    " + logLine.MyGrid + "\n")
	output.WriteString("QslMsg    " + logLine.QslMsgFromHeader + "\n")
	output.WriteString("Nickname  " + logLine.Nickname + "\n")
	output.WriteString("Mode      " + logLine.Mode + "\n")
	output.WriteString("ModeType  " + logLine.ModeType + "\n")
	output.WriteString("Band      " + logLine.Band + "\n")
	output.WriteString("  Lower   " + fmt.Sprintf("%f", logLine.BandLowerLimit) + "\n")
	output.WriteString("  Upper   " + fmt.Sprintf("%f", logLine.BandUpperLimit) + "\n")
	output.WriteString("Frequency " + logLine.Frequency + "\n")
	output.WriteString("Time      " + logLine.Time + "\n")
	output.WriteString("Call      " + logLine.Call + "\n")
	output.WriteString("Comment   " + logLine.Comment + "\n")
	output.WriteString("QSLmsg    " + logLine.QSLmsg + "\n")
	output.WriteString("OMname    " + logLine.OMname + "\n")
	output.WriteString("GridLoc   " + logLine.GridLoc + "\n")
	output.WriteString("RSTsent   " + logLine.RSTsent + "\n")
	output.WriteString("RSTrcvd   " + logLine.RSTrcvd + "\n")
	output.WriteString("SOTA      " + logLine.SOTA + "\n")
	output.WriteString("WWFF      " + logLine.WWFF + "\n")

	return output.String()
}

// SprintHeaderValues displays the header values
func SprintHeaderValues(logLine LogLine) string {
	var output strings.Builder

	output.WriteString("MyCall    " + logLine.MyCall)
	if logLine.Operator != "" {
		output.WriteString(" (" + logLine.Operator + ")")
	}
	output.WriteString("\n")

	if logLine.MyWWFF != "" {
		output.WriteString("MyWWFF    " + logLine.MyWWFF + "\n")
	}

	if logLine.MySOTA != "" {
		output.WriteString("MySOTA    " + logLine.MySOTA + "\n")
	}

	if logLine.MyPOTA != "" {
		output.WriteString("MyPOTA    " + logLine.MyPOTA + "\n")
	}

	if logLine.MyGrid != "" {
		output.WriteString("MyGrid    " + logLine.MyGrid + "\n")
	}

	if logLine.MyLat != "" {
		output.WriteString("MyLat     " + logLine.MyLat + "\n")
	}

	if logLine.MyLon != "" {
		output.WriteString("MyLon     " + logLine.MyLon + "\n")
	}

	return output.String()
}

// Date, Time, band, mode, call, report sent, report rcvd, Notes
var logLineFormat = "%-10s %-4s %-4s %-4s %-12s %-4s %-4s %s\n"

// SprintColumnTitles displays the column titles for a log line
func SprintColumnTitles() string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf(logLineFormat, "Date", "Time", "Band", "Mode", "Call", "Sent", "Rcvd", "Notes"))
	output.WriteString(fmt.Sprintf(logLineFormat, "----", "----", "----", "----", "----", "----", "----", "----"))
	return output.String()
}

// SprintLogInColumn displays the logLine in column mode
func SprintLogInColumn(logLine LogLine) (output string) {
	var notes strings.Builder
	if logLine.Frequency != "" {
		notes.WriteString("QRG: " + logLine.Frequency + " ")
	}
	if logLine.Comment != "" {
		notes.WriteString("[" + logLine.Comment + "] ")
	}
	if logLine.QSLmsg != "" {
		notes.WriteString("[" + logLine.QSLmsg + "] ")
	}
	if logLine.OMname != "" {
		notes.WriteString(logLine.OMname + " ")
	}
	if logLine.GridLoc != "" {
		notes.WriteString(logLine.GridLoc + " ")
	}
	if logLine.WWFF != "" {
		notes.WriteString(logLine.WWFF + " ")
	}
	if logLine.SOTA != "" {
		notes.WriteString(logLine.SOTA + " ")
	}

	output = fmt.Sprintf(logLineFormat, logLine.Date, logLine.Time, logLine.Band, logLine.Mode, logLine.Call, logLine.RSTsent, logLine.RSTrcvd, notes.String())

	return output
}
