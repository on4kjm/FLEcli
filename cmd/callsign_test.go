package cmd

import "testing"

func TestValidateCall(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name         string
		args 		 args
		//args         string
		wantCall     string
		wantErrorMsg string
	}{
		{"Good call (simple)", args.String("on4kjm"), "ON4KJM", ""},
		{"Good call (suffix)", "on4kjm/p", "ON4KJM/P", ""},
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
