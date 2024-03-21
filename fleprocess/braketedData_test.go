package fleprocess

import "testing"

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

func Test_getBraketedData(t *testing.T) {
	type args struct {
		inputLine  string
		braketType BraketType
	}
	tests := []struct {
		name             string
		args             args
		wantBraketedData string
		wantCleanedLine  string
	}{
		{
			"Happy case: comment",
			args{inputLine: "aaaa <bracketed text> bbbbb", braketType: COMMENT},
			"bracketed text",
			"aaaa  bbbbb",
		},
		{
			"Happy case: QSL",
			args{inputLine: "aaaa [bracketed text] bbbbb", braketType: QSL},
			"bracketed text",
			"aaaa  bbbbb",
		},
		{
			"Happy case: nothing",
			args{inputLine: "aaaa bbbbb cccccc", braketType: QSL},
			"",
			"aaaa bbbbb cccccc",
		},
		{
			"Empty brackets",
			args{inputLine: "aaaa <> bbbbb", braketType: COMMENT},
			"",
			"aaaa  bbbbb",
		},
		{
			"Brackets at right",
			args{inputLine: "aaaa bbbbb <bracketed text>", braketType: COMMENT},
			"bracketed text",
			"aaaa bbbbb ",
		},
		{
			"concatenated",
			args{inputLine: "aaaa<bracketed text>bbbbb", braketType: COMMENT},
			"bracketed text",
			"aaaabbbbb",
		},
		{
			"duplicated",
			args{inputLine: "aaaa <bracketed text> bbbbb < double > cccc", braketType: COMMENT},
			"bracketed text",
			"aaaa  bbbbb < double > cccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBraketedData, gotCleanedLine := getBracketedData(tt.args.inputLine, tt.args.braketType)
			if gotBraketedData != tt.wantBraketedData {
				t.Errorf("getBraketedData() gotBraketedData = %v, want %v", gotBraketedData, tt.wantBraketedData)
			}
			if gotCleanedLine != tt.wantCleanedLine {
				t.Errorf("getBraketedData() gotCleanedLine = %v, want %v", gotCleanedLine, tt.wantCleanedLine)
			}
		})
	}
}
