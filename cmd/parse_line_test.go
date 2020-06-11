package cmd

import (
	"reflect"
	"testing"
)


func TestParseLine(t *testing.T) {
	type args struct {
		inputStr     string
		previousLine LogLine
	}
	tests := []struct {
		name         string
		args         args
		wantLogLine  LogLine
		wantErrorMsg string
	}{
		{
			"Parse band and mode only", 
			args{ inputStr: "40m cw", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Band: "40m", Mode: "CW",}, "",
		},
		{
			"Parse for comment", 
			args{ inputStr: "4 g3noh <PSE QSL Direct>", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Comment: "PSE QSL Direct", Call: "G3NOH", Time: "2"}, "",
		},
		{
			"Parse for QSL", 
			args{ inputStr: "4 g3noh <Custom QSL message>", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ QSLmsg: "Custom QSL message", Call: "G3NOH", Time: "2"}, "",
		},
		{
			"Wrong mode", 
			args{ inputStr: "cww", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Mode: "SSB",}, "Unable to parse cww ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogLine, gotErrorMsg := ParseLine(tt.args.inputStr, tt.args.previousLine)
			if !reflect.DeepEqual(gotLogLine, tt.wantLogLine) {
				t.Errorf("ParseLine() gotLogLine = %v, want %v", gotLogLine, tt.wantLogLine)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ParseLine() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
