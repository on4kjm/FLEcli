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

func Test_csvDate(t *testing.T) {
	type args struct {
		inputDate string
	}
	tests := []struct {
		name           string
		args           args
		wantOutputDate string
	}{
		{
			"Happy case",
			args{inputDate: "2020-07-13"},
			"13/07/20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputDate := csvDate(tt.args.inputDate); gotOutputDate != tt.wantOutputDate {
				t.Errorf("csvDate() = %v, want %v", gotOutputDate, tt.wantOutputDate)
			}
		})
	}
}

func Test_buildCsv(t *testing.T) {
	sampleFilledLog1 := []LogLine{
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "2020-05-24", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", MySOTA: "ON/ON-001", Operator: "ON4KJM", Nickname: "ONFF-0259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "2020-05-24", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MySOTA: "ON/ON-001", Operator: "ON4KJM"},
	}
	//add case with no SOTA and with or no comment

	expectedOutput1 := []string{
		"V2,ON4KJM/P,ON/ON-001,24/05/20,1310,14MHz,CW,S57LC",
		"V2,ON4KJM/P,ON/ON-001,24/05/20,1312,14MHz,CW,ON4LY",
	}

	type args struct {
		fullLog []LogLine
	}
	tests := []struct {
		name        string
		args        args
		wantCsvList []string
	}{
		{
			"Happy case",
			args{fullLog: sampleFilledLog1},
			expectedOutput1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCsvList := buildCsv(tt.args.fullLog); !reflect.DeepEqual(gotCsvList, tt.wantCsvList) {
				t.Errorf("buildCsv() = %v, want %v", gotCsvList, tt.wantCsvList)
			}
		})
	}
}
