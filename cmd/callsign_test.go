package cmd

import "testing"

func TestValidateCall(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name         string
		args 		 args
		wantCall     string
		wantErrorMsg string
	}{
		{
			"Good call (simple)", 
			args{ sign: "on4kjm", }, 
			"ON4KJM", 
			"",
		},
		{
			"Good call (suffix)", 
			args{ sign: "on4kjm/p", }, 
			"ON4KJM/P", 
			"",
		},
		{
			"Good call (prefix only)", 
			args{ sign: "DL/on4KJm", }, 
			"DL/ON4KJM", 
			"",
		},
		{
			"Good call (prefix and suffix)", 
			args{ sign: "DL/on4KJm/p", }, 
			"DL/ON4KJM/P", 
			"",
		},
		//Error cases
		{
			"Pure junk passed", 
			args{ sign: "aaaaaa", }, 
			"*AAAAAA", 
			"Invalid call",
		},
		{
			"empty string", 
			args{ sign: "", }, 
			"*", 
			"Invalid call",
		},
		{
			"string with spaces", 
			args{ sign: "  ", }, 
			"*", 
			"Invalid call",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCall, gotErrorMsg := ValidateCall(tt.args.sign)
			if gotCall != tt.wantCall {
				t.Errorf("ValidateCall() gotCall = %v, want %v", gotCall, tt.wantCall)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ValidateCall() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
