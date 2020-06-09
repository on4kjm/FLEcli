package cmd

import "testing"

func TestValidateDate(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name         string
		args         args
		wantRef      string
		wantErrorMsg string
	}{
		{
			"Good date (simple)", 
			args{ inputStr: "2020-06-10", }, 
			"2020-06-10", "",
		},
		{
			"Bad date (simple)", 
			args{ inputStr: "2020-13-10", }, 
			"*2020-13-10", "parsing time \"2020-13-10\": month out of range",
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRef, gotErrorMsg := ValidateDate(tt.args.inputStr)
			if gotRef != tt.wantRef {
				t.Errorf("ValidateDate() gotRef = %v, want %v", gotRef, tt.wantRef)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ValidateDate() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
