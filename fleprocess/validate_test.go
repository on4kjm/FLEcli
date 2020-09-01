package fleprocess

import (
	"testing"
)

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
			args{inputStr: "onff-0258"},
			"ONFF-0258", "",
		},
		{
			"Good ref (single digit country)",
			args{inputStr: "fff-0258"},
			"FFF-0258", "",
		},
		{
			"Good ref (Numerical country)",
			args{inputStr: "4xff-0258"},
			"4XFF-0258", "",
		},
		{
			"Bad ref (no country prefix)",
			args{inputStr: "ff-0258"},
			"*FF-0258", "[FF-0258] is an invalid WWFF reference",
		},
		{
			"Bad ref (wrong separator)",
			args{inputStr: "gff/0258"},
			"*GFF/0258", "[GFF/0258] is an invalid WWFF reference",
		},
		{
			"Bad ref (reference too short)",
			args{inputStr: "onff-258"},
			"*ONFF-258", "[ONFF-258] is an invalid WWFF reference",
		},
		{
			"Bad ref (no country prefix)",
			args{inputStr: "onff-02589"},
			"*ONFF-02589", "[ONFF-02589] is an invalid WWFF reference",
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
			args{inputStr: "on/ON-001"},
			"ON/ON-001", "",
		},
		{
			"Good ref (single digit prefix)",
			args{inputStr: "g/ON-001"},
			"G/ON-001", "",
		},
		{
			"Good ref (numerical prefix)",
			args{inputStr: "4x/ON-001"},
			"4X/ON-001", "",
		},
		{
			"Good ref (american style)",
			args{inputStr: "w4z/ON-001"},
			"W4Z/ON-001", "",
		},
		{
			"Bad ref (long prefix)",
			args{inputStr: "xxxx/ON-001"},
			"*XXXX/ON-001", "[XXXX/ON-001] is an invalid SOTA reference",
		},
		{
			"Bad ref (missing slash)",
			args{inputStr: "on ON-001"},
			"*ON ON-001", "[ON ON-001] is an invalid SOTA reference",
		},
		{
			"Bad ref (numerical region)",
			args{inputStr: "on/9N-001"},
			"*ON/9N-001", "[ON/9N-001] is an invalid SOTA reference",
		},
		{
			"Bad ref (too long region)",
			args{inputStr: "on/ONA-001"},
			"*ON/ONA-001", "[ON/ONA-001] is an invalid SOTA reference",
		},
		{
			"Bad ref (no dash)",
			args{inputStr: "on/ON/001"},
			"*ON/ON/001", "[ON/ON/001] is an invalid SOTA reference",
		},
		{
			"Bad ref (number too short)",
			args{inputStr: "on/ON-01"},
			"*ON/ON-01", "[ON/ON-01] is an invalid SOTA reference",
		},
		{
			"Bad ref (Number too long)",
			args{inputStr: "on/ON-9001"},
			"*ON/ON-9001", "[ON/ON-9001] is an invalid SOTA reference",
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

func TestValidateCall(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name         string
		args         args
		wantCall     string
		wantErrorMsg string
	}{
		{
			"Good call (simple)",
			args{sign: "on4kjm"},
			"ON4KJM", "",
		},
		{
			"Good call (suffix)",
			args{sign: "on4kjm/p"},
			"ON4KJM/P", "",
		},
		{
			"Good call (prefix only)",
			args{sign: "DL/on4KJm"},
			"DL/ON4KJM", "",
		},
		{
			"Good call (prefix and suffix)",
			args{sign: "DL/on4KJm/p"},
			"DL/ON4KJM/P", "",
		},
		{
			"Good call (Numerical prefix)",
			args{sign: "4x/on4KJm/p"},
			"4X/ON4KJM/P", "",
		},
		{
			"Good call (prefix and long suffix)",
			args{sign: "DL/on4KJm/qrpp "},
			"DL/ON4KJM/QRPP", "",
		},
		{
			"Valid prefix (issue #2)",
			args{sign: "e7/z35m/p"},
			"E7/Z35M/P", "",
		},
		{
			"Valid E7 callsign (issue #2)",
			args{sign: "e7xyz"},
			"E7XYZ", "",
		},
		//*** Error cases *****
		{
			"Pure junk passed",
			args{sign: "aaaaaa"},
			"*AAAAAA", "[AAAAAA] is an invalid call",
		},
		{
			"empty string",
			args{sign: ""},
			"*", "[] is an invalid call",
		},
		{
			"string with spaces",
			args{sign: "  "},
			"*", "[] is an invalid call",
		},
		{
			"invalid prefix",
			args{sign: "xyz4/on4kjm"},
			"*XYZ4/ON4KJM", "[XYZ4] is an invalid prefix",
		},
		{
			"invalid prefix (when suffix is supplied)",
			args{sign: "xyz4/on4kjm/p"},
			"*XYZ4/ON4KJM/P", "[XYZ4] is an invalid prefix",
		},
		{
			"Too many /",
			args{sign: "F/on4kjm/p/x"},
			"*F/ON4KJM/P/X", "[F/ON4KJM/P/X] is invalid: too many '/'",
		},
		{
			"signe /",
			args{sign: "/"},
			"*/", "[] is an invalid call",
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
			args{inputStr: "2020-06-10"},
			"2020-06-10", "",
		},
		// {
		// 	"Good date (extrapolate, different delimiter)",
		// 	args{ inputStr: "16-2-1", },
		// 	"2020-06-10", "",
		// },
		{
			"Bad date (simple)",
			args{inputStr: "2020-13-10"},
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

func TestIsBand(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name            string
		args            args
		wantResult      bool
		wantLowerLimit  float64
		wantUpperLimit  float64
		wantAltBandName string
	}{
		{
			"invalid band",
			args{inputStr: "zzzz"},
			false, 0, 0, "",
		},
		{
			"valid band",
			args{inputStr: "40m"},
			true, 7.0, 7.3, "7MHz",
		},
		{
			"valid band but uppercase",
			args{inputStr: "60M"},
			true, 5.06, 5.45, "5MHz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotLowerLimit, gotUpperLimit, gotAltBandName := IsBand(tt.args.inputStr)
			if gotResult != tt.wantResult {
				t.Errorf("IsBand() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotLowerLimit != tt.wantLowerLimit {
				t.Errorf("IsBand() gotLowerLimit = %v, want %v", gotLowerLimit, tt.wantLowerLimit)
			}
			if gotUpperLimit != tt.wantUpperLimit {
				t.Errorf("IsBand() gotUpperLimit = %v, want %v", gotUpperLimit, tt.wantUpperLimit)
			}
			if gotAltBandName != tt.wantAltBandName {
				t.Errorf("IsBand() gotAltBandName = %v, want %v", gotAltBandName, tt.wantAltBandName)
			}
		})
	}
}

func TestValidateGridLocator(t *testing.T) {
	type args struct {
		grid string
	}
	tests := []struct {
		name              string
		args              args
		wantProcessedGrid string
		wantErrorMsg      string
	}{
		{
			"invalid grid",
			args{grid: "zzzz"},
			"*zzzz", "[zzzz] is an invalid grid reference",
		},
		{
			"Valid 4 pos grid",
			args{grid: "JO20"},
			"JO20", "",
		},
		{
			"Valid 4 pos grid (mixed case)",
			args{grid: "Jo20"},
			"JO20", "",
		},
		{
			"Valid 6 pos grid",
			args{grid: "JO20ec"},
			"JO20ec", "",
		},
		{
			"Valid 6 pos grid (mixed case)",
			args{grid: "Jo20Ec"},
			"JO20ec", "",
		},
		{
			"Valid grid but over 6 pos",
			args{grid: "JO20ec16"},
			"*JO20ec16", "[JO20ec16] is an invalid grid reference",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProcessedGrid, gotErrorMsg := ValidateGridLocator(tt.args.grid)
			if gotProcessedGrid != tt.wantProcessedGrid {
				t.Errorf("ValidateGridLocator() gotProcessedGrid = %v, want %v", gotProcessedGrid, tt.wantProcessedGrid)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ValidateGridLocator() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}
