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
			args{inputStr: "40M cw", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Band: "40m", BandLowerLimit: 7, BandUpperLimit: 7.3, Mode: "CW", ModeType: "CW", RSTsent: "599", RSTrcvd: "599"}, "",
		},
		{
			"Parse for time",
			args{inputStr: "1314 g3noh", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Time: "1314", ActualTime: "1314", Call: "G3NOH", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse partial time - 1",
			args{inputStr: "4 g3noh", previousLine: LogLine{Time: "", Mode: "SSB"}},
			LogLine{Time: "4", ActualTime: "4", Call: "G3NOH", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "", //TODO: should fail
		},
		{
			"Parse partial time - 2",
			args{inputStr: "15 g3noh", previousLine: LogLine{Time: "1200", Mode: "SSB"}},
			LogLine{Time: "1215", ActualTime: "1215", Call: "G3NOH", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse partial time - 3",
			args{inputStr: "4 g3noh", previousLine: LogLine{Time: "1200", Mode: "SSB"}},
			LogLine{Time: "1204", ActualTime: "1204", Call: "G3NOH", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse for comment",
			args{inputStr: "4 g3noh <PSE QSL Direct>", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Time: "4", ActualTime: "4", Comment: "PSE QSL Direct", Call: "G3NOH", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse for QSL",
			args{inputStr: "g3noh [Custom QSL message]", previousLine: LogLine{Mode: "SSB"}},
			LogLine{QSLmsg: "Custom QSL message", Call: "G3NOH", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Wrong mode",
			args{inputStr: "cww", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "Unable to make sense of [cww]. ",
		},
		{
			"Parse OM name",
			args{inputStr: "1314 g3noh @Jean", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Time: "1314", ActualTime: "1314", Call: "G3NOH", OMname: "Jean", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse Grid locator OK",
			args{inputStr: "1314 g3noh #jo50eJ", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Time: "1314", ActualTime: "1314", Call: "G3NOH", GridLoc: "JO50ej", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse Grid locator NOK",
			args{inputStr: "#grid", previousLine: LogLine{Mode: "SSB"}},
			LogLine{GridLoc: "*grid", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "[grid] is an invalid grid reference",
		},
		{
			"Parse frequency",
			args{inputStr: "14.153 on4kjm", previousLine: LogLine{Mode: "SSB", Band: "20m", BandLowerLimit: 14.0, BandUpperLimit: 14.35}},
			LogLine{Band: "20m", BandLowerLimit: 14.0, BandUpperLimit: 14.35, Frequency: "14.153", Call: "ON4KJM", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"Parse frequency out of limit",
			args{inputStr: "14.453 on4kjm", previousLine: LogLine{Mode: "SSB", Band: "20m", BandLowerLimit: 14.0, BandUpperLimit: 14.35}},
			LogLine{Band: "20m", BandLowerLimit: 14.0, BandUpperLimit: 14.35, Call: "ON4KJM", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "Frequency [14.453] is invalid for 20m band.",
		},
		{
			"Parse frequency out of limit with no band defined",
			args{inputStr: "14.453 on4kjm", previousLine: LogLine{Mode: "SSB"}},
			LogLine{Call: "ON4KJM", Mode: "SSB", RSTsent: "59", RSTrcvd: "59"}, "Unable to load frequency [14.453]: no band defined for that frequency.",
		},
		{
			"parse partial RST (sent) - CW",
			args{inputStr: "1230 on4kjm 5", previousLine: LogLine{Mode: "CW", ModeType: "CW"}},
			LogLine{Call: "ON4KJM", Time: "1230", ActualTime: "1230", RSTsent: "559", RSTrcvd: "599", Mode: "CW", ModeType: "CW"}, "",
		},
		{
			"parse partial RST (received) - CW",
			args{inputStr: "1230 on4kjm 5 44", previousLine: LogLine{Mode: "CW", ModeType: "CW"}},
			LogLine{Call: "ON4KJM", Time: "1230", ActualTime: "1230", RSTsent: "559", RSTrcvd: "449", Mode: "CW", ModeType: "CW"}, "",
		},
		{
			"parse full RST (received) - CW",
			args{inputStr: "1230 on4kjm 5 448", previousLine: LogLine{Mode: "CW", ModeType: "CW"}},
			LogLine{Call: "ON4KJM", Time: "1230", ActualTime: "1230", RSTsent: "559", RSTrcvd: "448", Mode: "CW", ModeType: "CW"}, "",
		},
		{
			"parse partial report (sent) - FM",
			args{inputStr: "1230 on4kjm 5", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "ON4KJM", Time: "1230", ActualTime: "1230", RSTsent: "55", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE"}, "",
		},
		{
			"parse partial report (received) - FM",
			args{inputStr: "1230 on4kjm 5 44", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "ON4KJM", Time: "1230", ActualTime: "1230", RSTsent: "55", RSTrcvd: "44", Mode: "FM", ModeType: "PHONE"}, "",
		},
		{
			"Incompatible report",
			args{inputStr: "1230 on4kjm 5 599", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "ON4KJM", Time: "1230", ActualTime: "1230", RSTsent: "55", RSTrcvd: "*599", Mode: "FM", ModeType: "PHONE"}, "Invalid report [599] for PHONE mode.",
		},
		{
			"SOTA keywork ",
			args{inputStr: "1230 oe6cud/p sota oe/st-309", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", SOTA: "OE/ST-309"}, "",
		},
		{
			"implied SOTA keywork ",
			args{inputStr: "1230 oe6cud/p oe/st-309", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", SOTA: "OE/ST-309"}, "",
		},
		{
			"WWFF keywork ",
			args{inputStr: "1230 oe6cud/p wwff onff-0258", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", WWFF: "ONFF-0258"}, "",
		},
		{
			"implied WWFF keywork ",
			args{inputStr: "1230 oe6cud/p onff-0258", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", WWFF: "ONFF-0258"}, "",
		},
		{
			"date processing",
			args{inputStr: "20.09.7 1230 oe6cud/p onff-0258", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Date: "2020-09-07", Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", WWFF: "ONFF-0258"}, "",
		},
		{
			"date processing (with keyword)",
			args{inputStr: "Date 20.09.7 1230 oe6cud/p onff-0258", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Date: "2020-09-07", Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", WWFF: "ONFF-0258"}, "",
		},
		{
			"date processing - validation error",
			args{inputStr: "20.09.34 1230 oe6cud/p onff-0258", previousLine: LogLine{Mode: "FM", ModeType: "PHONE"}},
			LogLine{Date: "*2020-09-34", Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE", WWFF: "ONFF-0258"}, "Error parsing time \"2020-09-34\": day out of range",
		},
		{
			"date processing - day ",
			args{inputStr: "day ++ 1230 oe6cud/p ", previousLine: LogLine{Date: "2020-09-05", Mode: "FM", ModeType: "PHONE"}},
			LogLine{Date: "2020-09-07", Call: "OE6CUD/P", Time: "1230", ActualTime: "1230", RSTsent: "59", RSTrcvd: "59", Mode: "FM", ModeType: "PHONE"}, "",
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

func TestHappyParseLine(t *testing.T) {
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
			"test1",
			args{inputStr: "1202 g4elz",
				previousLine: LogLine{Mode: "CW", ModeType: "CW", Band: "40m", BandLowerLimit: 7, BandUpperLimit: 7.3}},
			LogLine{Time: "1202", ActualTime: "1202", Call: "G4ELZ", Band: "40m", BandLowerLimit: 7, BandUpperLimit: 7.3, Mode: "CW", ModeType: "CW", RSTsent: "599", RSTrcvd: "599"}, "",
		},
		{
			"test2",
			args{inputStr: "4 g3noh <PSE QSL Direct>",
				previousLine: LogLine{Time: "1202", Mode: "CW", ModeType: "CW", Band: "40m", BandLowerLimit: 7, BandUpperLimit: 7.3}},
			LogLine{Time: "1204", ActualTime: "1204", Call: "G3NOH", Band: "40m", BandLowerLimit: 7, BandUpperLimit: 7.3, Mode: "CW", ModeType: "CW", Comment: "PSE QSL Direct", RSTsent: "599", RSTrcvd: "599"}, "",
		},
		{
			"test3",
			args{inputStr: "1227 gw4gte <Dave>",
				previousLine: LogLine{Time: "1202", Mode: "FM", ModeType: "PHONE", Band: "2m", BandLowerLimit: 144, BandUpperLimit: 148}},
			LogLine{Time: "1227", ActualTime: "1227", Call: "GW4GTE", Band: "2m", BandLowerLimit: 144, BandUpperLimit: 148, Mode: "FM", ModeType: "PHONE", Comment: "Dave", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"test4",
			args{inputStr: "8 gw0tlk/m gwff-0021",
				previousLine: LogLine{Time: "1227", Mode: "FM", ModeType: "PHONE", Band: "2m", BandLowerLimit: 144, BandUpperLimit: 148}},
			LogLine{Time: "1228", ActualTime: "1228", Call: "GW0TLK/M", Band: "2m", BandLowerLimit: 144, BandUpperLimit: 148, Mode: "FM", ModeType: "PHONE", WWFF: "GWFF-0021", RSTsent: "59", RSTrcvd: "59"}, "",
		},
		{
			"test5",
			args{inputStr: "7 dl0dan/p dlff-0002 dl/al-044",
				previousLine: LogLine{Time: "1220", Mode: "FM", ModeType: "PHONE", Band: "2m", BandLowerLimit: 144, BandUpperLimit: 148}},
			LogLine{Time: "1227", ActualTime: "1227", Call: "DL0DAN/P", Band: "2m", BandLowerLimit: 144, BandUpperLimit: 148, Mode: "FM", ModeType: "PHONE", WWFF: "DLFF-0002", SOTA: "DL/AL-044", RSTsent: "59", RSTrcvd: "59"}, "",
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
