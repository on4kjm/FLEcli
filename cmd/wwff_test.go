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
			"Good ref (simple)", 
			args{ inputStr: "onff-0258", }, 
			"ONFF-0258", "",
		},
		{
			"Good ref (single digit country)", 
			args{ inputStr: "fff-0258", }, 
			"FFF-0258", "",
		},
		{
			"Good ref (Numerical country)", 
			args{ inputStr: "4xff-0258", }, 
			"4XFF-0258", "",
		},
		{
			"Bad ref (no country prefix)", 
			args{ inputStr: "ff-0258", }, 
			"*FF-0258", "Invalid WWFF reference",
		},
		{
			"Bad ref (wrong separator)", 
			args{ inputStr: "gff/0258", }, 
			"*GFF/0258", "Invalid WWFF reference",
		},
		{
			"Bad ref (reference too short)", 
			args{ inputStr: "onff-258", }, 
			"*ONFF-258", "Invalid WWFF reference",
		},
		{
			"Bad ref (no country prefix)", 
			args{ inputStr: "onff-02589", }, 
			"*ONFF-02589", "Invalid WWFF reference",
		},

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
