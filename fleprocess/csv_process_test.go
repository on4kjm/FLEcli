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
			fmt.Errorf("Missing date for log entry at 12:02 (#2)\nMissing date for log entry at 12:03 (#3)\n"),
		},
		{
			"Missing MyCall",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
				{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"}},
			},
			nil,
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
