package cmd

import "testing"

func Test_getBraketedData(t *testing.T) {
	type args struct {
		inputLine  string
		braketType string
	}
	tests := []struct {
		name            string
		args            args
		wantText        string
		wantCleanedLine string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotText, gotCleanedLine := getBraketedData(tt.args.inputLine, tt.args.braketType)
			if gotText != tt.wantText {
				t.Errorf("getBraketedData() gotText = %v, want %v", gotText, tt.wantText)
			}
			if gotCleanedLine != tt.wantCleanedLine {
				t.Errorf("getBraketedData() gotCleanedLine = %v, want %v", gotCleanedLine, tt.wantCleanedLine)
			}
		})
	}
}
