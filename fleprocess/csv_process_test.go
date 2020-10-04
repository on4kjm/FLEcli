package fleprocess

import (
	"fmt"
	"testing"
)

func Test_validateDataForSotaCsv(t *testing.T) {
	type args struct {
		loadedLogFile []LogLine
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"Happy Case (activator)",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			nil,
		},
		{
			"Happy Case (chaser)",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call", SOTA: "Sota1"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call", SOTA: "Sota2"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call", SOTA: "Sota3"}},
			},
			nil,
		},
		{
			"Missing Date",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
				{Date: "", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"}},
			},
			fmt.Errorf("missing date for log entry at 12:02 (#2), missing date for log entry at 12:03 (#3)"),
		},
		{
			"Missing MyCall",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"}},
			},
			fmt.Errorf("Missing MyCall"),
		},
		{
			"Neither Activator nor Chaser",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:03", Call: "call"}},
			},
			fmt.Errorf("missing SOTA reference while attempting to process chaser log for log entry at 12:01 (#1), missing SOTA reference while attempting to process chaser log for log entry at 12:02 (#2), missing SOTA reference while attempting to process chaser log for log entry at 12:03 (#3)"),
		},
		{
			"SOTA Chaser log with one reference missing",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:01", Call: "call", SOTA: "Sota1"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:03", Call: "call", SOTA: "Sota3"}},
			},
			fmt.Errorf("missing SOTA reference while attempting to process chaser log for log entry at 12:02 (#2)"),
		},
		{
			"SOTA Chaser log with mySota redefinition",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:01", Call: "call", SOTA: "Sota1"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call", SOTA: "Sota2"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "12:03", Call: "call", SOTA: "Sota3"}},
			},
			fmt.Errorf("encountered an unexpexted MySota reference while processing what should be a chaser log for log entry at 12:02 (#2)"),
		},
		{
			"Misc. missing data (Band, Time, Mode, Call)",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "", Time: "", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: ""}},
			},
			fmt.Errorf("missing band for log entry #1, missing QSO time for log entry #1, missing mode for log entry at 12:02 (#2), missing call for log entry at 12:03 (#3)"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateDataForSotaCsv(tt.args.loadedLogFile)

			//Test the error message, if any
			if got != nil && tt.want != nil {
				if got.Error() != tt.want.Error() {
					t.Errorf("validateDataForSotaCsv() = %v, want %v", got, tt.want)
				}
			} else {
				if !(got == nil && tt.want == nil) {
					t.Errorf("validateDataForSotaCsv() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestProcessCsvCommand(t *testing.T) {
	type args struct {
		inputFilename     string
		outputCsvFilename string
		isInterpolateTime bool
		isOverwriteCsv    bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Bad output filename (directory)",
			args{inputFilename: "../test/data/fle-4-no-qso.txt", outputCsvFilename: "../test/data", isInterpolateTime: false, isOverwriteCsv: false},
			true,
		},
		{
			"input file parsing errors",
			args{inputFilename: "../test/data/fle-3-error.txt", outputCsvFilename: "", isInterpolateTime: false, isOverwriteCsv: false},
			true,
		},
		{
			"No QSO in loaded file",
			args{inputFilename: "../test/data/fle-4-no-qso.txt", outputCsvFilename: "", isInterpolateTime: false, isOverwriteCsv: false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessCsvCommand(tt.args.inputFilename, tt.args.outputCsvFilename, tt.args.isInterpolateTime, tt.args.isOverwriteCsv); (err != nil) != tt.wantErr {
				t.Errorf("ProcessCsvCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
