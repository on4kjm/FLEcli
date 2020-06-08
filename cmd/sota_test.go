package cmd

import "testing"

func TestValidateSota(t *testing.T) {
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
			args{ inputStr: "on/ON-001", }, 
			"ON/ON-001", "",
		},
		{
			"Good ref (single digit prefix)", 
			args{ inputStr: "g/ON-001", }, 
			"G/ON-001", "",
		},
		{
			"Good ref (numerical prefix)", 
			args{ inputStr: "4x/ON-001", }, 
			"4X/ON-001", "",
		},	
		{
			"Good ref (american style)", 
			args{ inputStr: "w4z/ON-001", }, 
			"W4Z/ON-001", "",
		},		
		{
			"Bad ref (long prefix)", 
			args{ inputStr: "xxxx/ON-001", }, 
			"*XXXX/ON-001", "Invalid SOTA reference",
		},	
		{
			"Bad ref (missing slash)", 
			args{ inputStr: "on ON-001", }, 
			"*ON ON-001", "Invalid SOTA reference",
		},		
		{
			"Bad ref (numerical region)", 
			args{ inputStr: "on/9N-001", }, 
			"*ON/9N-001", "Invalid SOTA reference",
		},		
		{
			"Bad ref (too long region)", 
			args{ inputStr: "on/ONA-001", }, 
			"*ON/ONA-001", "Invalid SOTA reference",
		},		
		{
			"Bad ref (no dash)", 
			args{ inputStr: "on/ON/001", }, 
			"*ON/ON/001", "Invalid SOTA reference",
		},		
		{
			"Bad ref (number too short)", 
			args{ inputStr: "on/ON-01", }, 
			"*ON/ON-01", "Invalid SOTA reference",
		},		
		{
			"Bad ref (Number too long)", 
			args{ inputStr: "on/ON-9001", }, 
			"*ON/ON-9001", "Invalid SOTA reference",
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRef, gotErrorMsg := ValidateSota(tt.args.inputStr)
			if gotRef != tt.wantRef {
				t.Errorf("ValidateSota() gotRef = %v, want %v", gotRef, tt.wantRef)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ValidateSota() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
