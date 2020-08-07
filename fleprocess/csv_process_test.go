package fleprocess

import "testing"

func Test_validateDataForSotaCsv(t *testing.T) {
	type args struct {
		loadedLogFile []LogLine
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Happy Case",
			args{loadedLogFile: []LogLine{
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"}},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateDataForSotaCsv(tt.args.loadedLogFile); got != tt.want {
				t.Errorf("validateDataForSotaCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}
