package fleprocess

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_validateDataforAdif2(t *testing.T) {
	type args struct {
		loadedLogFile []LogLine
		adifParams    AdifParams
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"Happy Case (no sota, pota or wwff)",
			args{
				adifParams: AdifParams{IsWWFF: false, IsSOTA: false, IsPOTA: false},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				},
			},
			nil,
		},
		{
			"No data",
			args{
				adifParams:    AdifParams{IsWWFF: false, IsSOTA: false, IsPOTA: false},
				loadedLogFile: []LogLine{},
			},
			fmt.Errorf("no QSO found"),
		},
		{
			"Missing Date",
			args{
				adifParams: AdifParams{IsWWFF: false, IsSOTA: false, IsPOTA: false},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
					{Date: "", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
					{Date: "", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"},
				},
			},
			fmt.Errorf("missing date for log entry at 12:02 (#2), missing date for log entry at 12:03 (#3)"),
		},
		{
			"Missing MyCall",
			args{
				adifParams: AdifParams{IsWWFF: true, IsSOTA: true, IsPOTA: true},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
					{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
					{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"},
				},
			},
			fmt.Errorf("missing MyCall"),
		},
		{
			"Missing MyCall (POTA)",
			args{
				adifParams: AdifParams{IsWWFF: false, IsSOTA: false, IsPOTA: true},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:01", Call: "call"},
					{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:02", Call: "call"},
					{Date: "date", MyCall: "", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: "call"},
				},
			},
			fmt.Errorf("missing MyCall"),
		},
		{
			"Missing MySota",
			args{
				adifParams: AdifParams{IsWWFF: false, IsSOTA: true, IsPOTA: false},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				},
			},
			fmt.Errorf("missing MY-SOTA reference"),
		},
		{
			"Missing MyPota",
			args{
				adifParams: AdifParams{IsWWFF: false, IsSOTA: false, IsPOTA: true},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				},
			},
			fmt.Errorf("missing MY-POTA reference"),
		},
		{
			"Misc. missing data (Band, Time, Mode, Call)",
			args{
				adifParams: AdifParams{IsWWFF: false, IsSOTA: false, IsPOTA: false},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "", Time: "", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "", Band: "band", Time: "12:02", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", Mode: "mode", Band: "band", Time: "12:03", Call: ""},
				},
			},
			fmt.Errorf("missing band for log entry #1, missing QSO time for log entry #1, missing mode for log entry at 12:02 (#2), missing call for log entry at 12:03 (#3)"),
		},
		{
			"Missing MY-WWFF",
			args{
				adifParams: AdifParams{IsWWFF: true, IsSOTA: false, IsPOTA: false},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", MyWWFF: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", MyWWFF: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MySOTA: "mySota", MyWWFF: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				},
			},
			fmt.Errorf("missing MY-WWFF reference"),
		},
		{
			"Missing Operator with isWWFF",
			args{
				adifParams: AdifParams{IsWWFF: true, IsSOTA: false, IsPOTA: false},
				loadedLogFile: []LogLine{
					{Date: "date", MyCall: "myCall", MyWWFF: "myWwff", Operator: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MyWWFF: "myWwff", Operator: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
					{Date: "date", MyCall: "myCall", MyWWFF: "myWwff", Operator: "", Mode: "mode", Band: "band", Time: "time", Call: "call"},
				},
			},
			fmt.Errorf("missing Operator call sign"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateDataforAdif(tt.args.loadedLogFile, tt.args.adifParams)

			//Test the error message, if any
			if got != nil && tt.want != nil {
				if got.Error() != tt.want.Error() {
					t.Errorf("validateDataforAdif() = %v, want %v", got, tt.want)
				}
			} else {
				if !(got == nil && tt.want == nil) { //nolint:staticcheck
					t.Errorf("validateDataforAdif() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestProcessAdifCommandSortsSOTAChronologicallyByDefault(t *testing.T) {
	outputFilename := filepath.Join(t.TempDir(), "sorted.adi")
	adifParams := AdifParams{
		InputFilename:  "../test/data/fle-adif-sort.txt",
		OutputFilename: outputFilename,
		IsSOTA:         true,
	}

	if err := ProcessAdifCommand(adifParams); err != nil {
		t.Fatalf("ProcessAdifCommand() error = %v", err)
	}

	output, err := os.ReadFile(outputFilename)
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}

	adif := string(output)
	positions := []int{
		strings.Index(adif, "<TIME_ON:4>2149"),
		strings.Index(adif, "<TIME_ON:4>2151"),
		strings.Index(adif, "<TIME_ON:4>2153"),
	}
	for i, position := range positions {
		if position == -1 {
			t.Fatalf("sorted ADIF is missing expected time at index %d", i)
		}
		if i > 0 && position <= positions[i-1] {
			t.Fatalf("ADIF times are not chronological: positions = %v", positions)
		}
	}
}

func TestProcessAdifCommand(t *testing.T) {
	type args struct {
		adifParams AdifParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Bad output filename (directory)",
			args{
				adifParams: AdifParams{
					InputFilename:     "../test/data/fle-4-no-qso.txt",
					OutputFilename:    "../test/data",
					IsInterpolateTime: false,
					IsOverwrite:       false,
				},
			},
			true,
		},
		{
			"input file parsing errors (missing band)",
			args{
				adifParams: AdifParams{
					InputFilename:     "../test/data/fle-3-error.txt",
					OutputFilename:    "",
					IsInterpolateTime: false,
					IsOverwrite:       false,
				},
			},
			true,
		},
		{
			"input file parsing errors (wrong call)",
			args{
				adifParams: AdifParams{
					InputFilename:     "../test/data/fle-5-wrong-call.txt",
					OutputFilename:    "",
					IsInterpolateTime: false,
					IsOverwrite:       false,
				},
			},
			true,
		},
		{
			"No QSO in loaded file",
			args{
				adifParams: AdifParams{
					InputFilename:     "../test/data/fle-4-no-qso.txt",
					OutputFilename:    "",
					IsInterpolateTime: false,
					IsOverwrite:       false,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ProcessAdifCommand(tt.args.adifParams); (err != nil) != tt.wantErr {
				t.Errorf("ProcessAdifCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
