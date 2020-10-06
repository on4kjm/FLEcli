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
	"testing"
)

func Test_validateDataforAdif(t *testing.T) {
	type args struct {
		loadedLogFile []LogLine
		isWWFFcli     bool
		isSOTAcli     bool
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"Happy Case (no sota or wwff)",
			args{isWWFFcli: false, isSOTAcli: false, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			nil,
		},
		{
			"No data",
			args{isWWFFcli: false, isSOTAcli: false, loadedLogFile: []LogLine{}},
			fmt.Errorf("No QSO found"),
		},
		{
			"Missing Date",
			args{isWWFFcli: false, isSOTAcli: false, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
				{Date: "", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"}},
			},
			fmt.Errorf("missing date for log entry at 12:02 (#2), missing date for log entry at 12:03 (#3)"),
		},
		{
			"Missing MyCall",
			args{isWWFFcli: true, isSOTAcli: true, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"}},
			},
			fmt.Errorf("Missing MyCall"),
		},
		{
			"Missing MySota",
			args{isWWFFcli: false, isSOTAcli: true, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			fmt.Errorf("Missing MY-SOTA reference"),
		},
		{
			"Misc. missing data (Band, Time, Mode, Call)",
			args{isWWFFcli: false, isSOTAcli: false, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "", Time: "", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: ""}},
			},
			fmt.Errorf("missing band for log entry #1, missing QSO time for log entry #1, missing mode for log entry at 12:02 (#2), missing call for log entry at 12:03 (#3)"),
		},
		{
			"Missing MY-WWFF",
			args{isWWFFcli: true, isSOTAcli: false, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", MyWWFF: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", MyWWFF: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", MyWWFF: "", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			fmt.Errorf("Missing MY-WWFF reference"),
		},
		{
			"Missing MY-WWFF",
			args{isWWFFcli: true, isSOTAcli: false, loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MyWWFF: "myWwff", Operator: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MyWWFF: "myWwff", Operator: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MyWWFF: "myWwff", Operator: "", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			fmt.Errorf("Missing Operator call sign"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateDataforAdif(tt.args.loadedLogFile, tt.args.isWWFFcli, tt.args.isSOTAcli)

			//Test the error message, if any
			if got != nil && tt.want != nil {
				if got.Error() != tt.want.Error() {
					t.Errorf("validateDataforAdif() = %v, want %v", got, tt.want)
				}
			} else {
				if !(got == nil && tt.want == nil) {
					t.Errorf("validateDataforAdif() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestProcessAdifCommand(t *testing.T) {
	type args struct {
		inputFilename     string
		outputFilename    string
		isInterpolateTime bool
		isWWFFcli         bool
		isSOTAcli         bool
		isOverwrite       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Bad output filename (directory)",
			args{inputFilename: "../test/data/fle-4-no-qso.txt", outputFilename: "../test/data", isInterpolateTime: false, isOverwrite: false},
			true,
		},
		{
			"input file parsing errors (missing band)",
			args{inputFilename: "../test/data/fle-3-error.txt", outputFilename: "", isInterpolateTime: false, isOverwrite: false},
			true,
		},
		{
			"input file parsing errors (wrong call)",
			args{inputFilename: "../test/data/fle-5-wrong-call.txt", outputFilename: "", isInterpolateTime: false, isOverwrite: false},
			true,
		},
		{
			"No QSO in loaded file",
			args{inputFilename: "../test/data/fle-4-no-qso.txt", outputFilename: "", isInterpolateTime: false, isOverwrite: false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessAdifCommand(tt.args.inputFilename, tt.args.outputFilename, tt.args.isInterpolateTime, tt.args.isWWFFcli, tt.args.isSOTAcli, tt.args.isOverwrite); (err != nil) != tt.wantErr {
				t.Errorf("ProcessCsvCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
