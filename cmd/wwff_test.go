package cmd

import "testing"

func TestValidateWwff(t *testing.T) {
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
			"Good call (simple)", 
			args{ inputStr: "onff-0258", }, 
			"ONFF-0258", "",
		},
		{
			"Good call (simple)", 
			args{ inputStr: "fff-0258", }, 
			"FFF-0258", "",
		},
		{
			"Good call (simple)", 
			args{ inputStr: "4xff-0258", }, 
			"4XFF-0258", "",
		},
		//TODO: add the invalid cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRef, gotErrorMsg := ValidateWwff(tt.args.inputStr)
			if gotRef != tt.wantRef {
				t.Errorf("ValidateWwff() gotRef = %v, want %v", gotRef, tt.wantRef)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ValidateWwff() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
