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
			args{ inputStr: "40M cw", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Band: "40m", BandLowerLimit: 7, BandUpperLimit: 7.3, Mode: "CW", ModeType: "CW", RSTsent: "599", RSTrcvd: "599"}, "",
		},
		{
			"Parse for time", 
			args{ inputStr: "1314 g3noh", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Time: "1314", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse partial time - 1", 
			args{ inputStr: "4 g3noh", previousLine: LogLine{ Time: "", Mode: "SSB", }}, 
			LogLine{ Time: "4", Call: "G3NOH", Mode: "SSB",}, "", //TODO: should fail
		},
		{
			"Parse partial time - 2", 
			args{ inputStr: "15 g3noh", previousLine: LogLine{ Time: "1200", Mode: "SSB", }}, 
			LogLine{ Time: "1215", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse partial time - 3", 
			args{ inputStr: "4 g3noh", previousLine: LogLine{ Time: "1200", Mode: "SSB", }}, 
			LogLine{ Time: "1204", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse for comment", 
			args{ inputStr: "4 g3noh <PSE QSL Direct>", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Time: "4", Comment: "PSE QSL Direct", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Parse for QSL", 
			args{ inputStr: "g3noh [Custom QSL message]", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ QSLmsg: "Custom QSL message", Call: "G3NOH", Mode: "SSB",}, "",
		},
		{
			"Wrong mode", 
			args{ inputStr: "cww", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ Mode: "SSB",}, "Unable to parse cww ",
		},
		{
			"Parse OM name", 
			args{ inputStr: "@Jean", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ OMname: "Jean", Mode: "SSB",}, "",
		},
		{
			"Parse Grid locator", 
			args{ inputStr: "#grid", previousLine: LogLine{ Mode: "SSB", }}, 
			LogLine{ GridLoc: "grid", Mode: "SSB",}, "",
		},
		{
			"parse partial RST (sent) - CW", 
			args{ inputStr: "1230 on4kjm 5", previousLine: LogLine{ Mode: "CW", ModeType: "CW"}}, 
			LogLine{ Call: "ON4KJM", Time: "1230", RSTsent: "559", Mode: "CW", ModeType: "CW"}, "",
		},
		{
			"parse partial RST (received) - CW", 
			args{ inputStr: "1230 on4kjm 5 44", previousLine: LogLine{ Mode: "CW", ModeType: "CW"}}, 
			LogLine{ Call: "ON4KJM", Time: "1230", RSTsent: "559", RSTrcvd: "449", Mode: "CW", ModeType: "CW"}, "",
		},
		{
			"parse full RST (received) - CW", 
			args{ inputStr: "1230 on4kjm 5 448", previousLine: LogLine{ Mode: "CW", ModeType: "CW"}}, 
			LogLine{ Call: "ON4KJM", Time: "1230", RSTsent: "559", RSTrcvd: "448", Mode: "CW", ModeType: "CW"}, "",
		},
		{
			"parse partial report (sent) - FM", 
			args{ inputStr: "1230 on4kjm 5", previousLine: LogLine{ Mode: "FM", ModeType: "PHONE"}}, 
			LogLine{ Call: "ON4KJM", Time: "1230", RSTsent: "55", Mode: "FM", ModeType: "PHONE"}, "",
		},
		{
			"parse partial report (received) - FM", 
			args{ inputStr: "1230 on4kjm 5 44", previousLine: LogLine{ Mode: "FM", ModeType: "PHONE"}}, 
			LogLine{ Call: "ON4KJM", Time: "1230", RSTsent: "55", RSTrcvd: "44", Mode: "FM", ModeType: "PHONE"}, "",
		},
		{
			"Incompatible report", 
			args{ inputStr: "1230 on4kjm 5 599", previousLine: LogLine{ Mode: "FM", ModeType: "PHONE"}}, 
			LogLine{ Call: "ON4KJM", Time: "1230", RSTsent: "55", RSTrcvd: "*599", Mode: "FM", ModeType: "PHONE"}, "Invalid report (599) for PHONE mode ",
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
