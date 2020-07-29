package fleprocess

import "testing"

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
			gotBraketedData, gotCleanedLine := getBraketedData(tt.args.inputLine, tt.args.braketType)
			if gotBraketedData != tt.wantBraketedData {
				t.Errorf("getBraketedData() gotBraketedData = %v, want %v", gotBraketedData, tt.wantBraketedData)
			}
			if gotCleanedLine != tt.wantCleanedLine {
				t.Errorf("getBraketedData() gotCleanedLine = %v, want %v", gotCleanedLine, tt.wantCleanedLine)
			}
		})
	}
}
