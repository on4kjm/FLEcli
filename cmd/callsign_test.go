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
			"ON4KJM", "",
		},
		{
			"Good call (suffix)", 
			args{ sign: "on4kjm/p", }, 
			"ON4KJM/P", "",
		},
		{
			"Good call (prefix only)", 
			args{ sign: "DL/on4KJm", }, 
			"DL/ON4KJM", "",
		},
		{
			"Good call (prefix and suffix)", 
			args{ sign: "DL/on4KJm/p", }, 
			"DL/ON4KJM/P", "",
		},
		{
			"Good call (Numerical prefix)", 
			args{ sign: "4x/on4KJm/p", }, 
			"4X/ON4KJM/P", "",
		},
		{
			"Good call (prefix and long suffix)", 
			args{ sign: "DL/on4KJm/qrpp ", }, 
			"DL/ON4KJM/QRPP", "",
		},
		//Error cases
		{
			"Pure junk passed", 
			args{ sign: "aaaaaa", }, 
			"*AAAAAA", "Invalid call",
		},
		{
			"empty string", 
			args{ sign: "", }, 
			"*", "Invalid call",
		},
		{
			"string with spaces", 
			args{ sign: "  ", }, 
			"*", "Invalid call",
		},		
		{
			"invalid prefix", 
			args{ sign: "xyz/on4kjm", }, 
			"*XYZ/ON4KJM", "Invalid prefix",
		},
		{
			"Too many /", 
			args{ sign: "F/on4kjm/p/x", }, 
			"*F/ON4KJM/P/X", "Too many '/'",
		},
		{
			"signe /", 
			args{ sign: "/", }, 
			"*/", "Invalid call",
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
