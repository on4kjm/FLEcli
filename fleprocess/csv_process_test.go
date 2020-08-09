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
			"Happy Case",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
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
			"Missing MySota",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			fmt.Errorf("Missing MY-SOTA reference"),
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
			if tt.want == nil || got == nil {
				if tt.want == nil && got != nil {
					t.Errorf("validateDataForSotaCsv() = %v, want %v", got, nil)
				}
				if tt.want != nil && got == nil {
					t.Errorf("validateDataForSotaCsv() = %v, want %v", nil, tt.want)
				}
			} else {
				if got.Error() != tt.want.Error() {
					t.Errorf("validateDataForSotaCsv() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
