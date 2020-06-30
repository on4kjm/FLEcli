package cmd

import (
	"fmt"
	"strings"
)

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

func writeAdif(outputFile string, fullLog []LogLine) {

	// TODO: create an array of strings first
	// TODO: write the array list to file
}

func buildAdif(fullLog []LogLine) (adifList []string) {
	//Print the fixed header
	adifList = append(adifList, "ADIF Export for Fast Log Entry by DF3CB")
	adifList = append(adifList, "<PROGRAMID:3>FLE")
	adifList = append(adifList, "<ADIF_VER:5>3.0.6")
	adifList = append(adifList, "<EOH>")

	for _, logLine := range fullLog {
		adifLine := ""
		adifLine = adifLine + adifElement("STATION_CALLSIGN", logLine.MyCall)
		adifLine = adifLine + adifElement("CALL", logLine.Call)
		//TODO: strip the delimiters of the date
		adifLine = adifLine + adifElement("QSO_DATE", logLine.Date)
		adifLine = adifLine + adifElement("TIME_ON", logLine.Time)
		adifLine = adifLine + adifElement("BAND", logLine.Band)
		adifLine = adifLine + adifElement("MODE", logLine.Mode)
		if logLine.Frequency != "" {
			adifLine = adifLine + adifElement("FREQ", logLine.Frequency)
		}
		adifLine = adifLine + adifElement("RST_SENT", logLine.RSTsent)
		adifLine = adifLine + adifElement("RST_RCVD", logLine.RSTrcvd)
		adifLine = adifLine + adifElement("MY_SIG", "WWFF")
		adifLine = adifLine + adifElement("MY_SIG_INFO", logLine.MyWWFF)
		adifLine = adifLine + adifElement("OPERATOR", logLine.Operator)
		if logLine.Nickname != "" {
			adifLine = adifLine + adifElement("APP_EQSL_QTH_NICKNAME", logLine.Nickname)
		}
		adifLine = adifLine + "<EOR>"

		adifList = append(adifList, adifLine)

	}

	return adifList
}

func adifElement(elementName, elementValue string) (element string) {
	return fmt.Sprintf("<%s:%d>%s ", strings.ToUpper(elementName), len(elementValue), elementValue)
}
