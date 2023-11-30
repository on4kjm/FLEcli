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

var pota_ref_list = []string{
	"K-0001", "CE9-0001", "VE-0001", "F-0001", "G-0001", "YB-0001", "VK-0001", "TI-0001", "I-0001",
	"9M-0001", "ZL-0001", "GW-0001", "9V-0001", "XE-0001", "HL-0001", "PA-0001", "SV-0001", "SP-0001",
	"DA-0001", "HS-0001", "CT-0001", "CU-0001", "HI-0001", "AR-0001", "9A-0001", "5B-0001", "HK-0001",
	"GM-0001", "JA-0001", "HP-0001", "CA-0001", "BV-0001", "OE-0001", "EA-0001", "HB-0001", "EK-0001",
	"V3-0001", "ZF-0001", "OA-0001", "LX-0001", "HR-0001", "TF-0001", "ON-0001", "OU-0001", "YV-0001",
	"8P-0001", "OX-0001", "OY-0001", "YT-0001", "DU-0001", "P4-0001", "LA-0001", "PY-0001", "VU-0001",
	"EI-0001", "6Y-0001", "LZ-0001", "TG-0001", "OH-0001", "HC-0001", "ZR-0001", "PZ-0001", "CV-0001",
	"8R-0001", "Z2-0001", "A2-0001", "CP-0001", "CO-0001", "9Y-0001", "ZP-0001", "C9-0001", "SM-0001",
	"4W-0001", "TA-0001", "V7-0001", "3D2-0001", "4X-0001", "S5-0001", "BY-0001", "P29-0001", "GI-0001",
	"T7-0001", "YO-0001", "HA-0001", "A3-0001", "XV2-0001", "J8-0001", "VP5-0001", "V5-0001", "HB0-0001",
	"YS-0001", "YJ8-0001", "UA-0001", "UR-0001", "V4-0001", "GD-0001", "V6-0001", "A6-0001", "Z3-0001",
	"5W-0001", "9G1-0001", "OK-0001", "C6-0001", "J6-0001", "J7-0001", "J3-0001", "T9-0001", "4L-0001",
	"7O-0001",
}

func Test_Pota_prefix(t *testing.T) {
	test_failed := false
	invalidRefs := ""
	for _, pota_ref := range  pota_ref_list {
		ref, _ := ValidatePota(pota_ref)
		if ref == "" {
			test_failed = true
			invalidRefs = invalidRefs + pota_ref + " validation failed \n"
		}
	}
	if test_failed {
		t.Error(invalidRefs)
	}
}

func TestValidatePota(t *testing.T) {
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
			args{inputStr: "on-0258"},
			"ON-0258", "",
		},
		{
			"Good ref (single digit country)",
			args{inputStr: "f-0258"},
			"F-0258", "",
		},
		{
			"Good ref (Numerical country)",
			args{inputStr: "4x-0258"},
			"4X-0258", "",
		},
		{
			"Good ref (country ref containing a digit)",
			args{inputStr: "HB0-0258"},
			"HB0-0258", "",
		},
		{
			"Good ref (country ref containing a digit)",
			args{inputStr: "E7-0258"},
			"E7-0258", "",
		},
		{
			"Good ref (5 digit park)",
			args{inputStr: "k-10177"},
			"K-10177", "",
		},
		{
			"Bad ref (no country prefix)",
			args{inputStr: "-0258"},
			"*-0258", "[-0258] is an invalid POTA reference",
		},
		{
			"Bad ref (wrong separator)",
			args{inputStr: "g/0258"},
			"*G/0258", "[G/0258] is an invalid POTA reference",
		},
		{
			"Bad ref (reference too short)",
			args{inputStr: "on-258"},
			"*ON-258", "[ON-258] is an invalid POTA reference",
		},
		{
			"Bad ref (park number too long)",
			args{inputStr: "on-123456"},
			"*ON-123456", "[ON-123456] is an invalid POTA reference",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRef, gotErrorMsg := ValidatePota(tt.args.inputStr)
			if gotRef != tt.wantRef {
				t.Errorf("ValidatePota() gotRef = %v, want %v", gotRef, tt.wantRef)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("ValidatePota() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
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
			"Good ref (american style) - Issue 78",
			args{inputStr: "W7A/YV-102"},
			"W7A/YV-102", "",
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
			"Valid call from activation",
			args{sign: "sm1/dl6jz/p"},
			"SM1/DL6JZ/P", "",
		},
		{
			"Valid call from activation (case 2)",
			args{sign: "Sm/dl8mf"},
			"SM/DL8MF", "",
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

func TestNormalizeDate(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name         string
		args         args
		wantDate     string
		wantErrorMsg string
	}{
		{
			"happy case",
			args{inputStr: "2020-09-04"},
			"2020-09-04", "",
		},
		{
			"alternate delimiter 1",
			args{inputStr: "2020/09/04"},
			"2020-09-04", "",
		},
		{
			"alternate delimiter 2",
			args{inputStr: "2020.09.04"},
			"2020-09-04", "",
		},
		{
			"alternate delimiter 3",
			args{inputStr: "2020 09 04"},
			"2020-09-04", "",
		},
		{
			"shortened date 1",
			args{inputStr: "20/09/04"},
			"2020-09-04", "",
		},
		{
			"shortened date 1",
			args{inputStr: "2020.9.4"},
			"2020-09-04", "",
		},
		{
			"Bad date",
			args{inputStr: "202009.04"},
			"*202009.04", "Bad date format: found 2 elements while expecting 3.",
		},
		{
			"Bad year length",
			args{inputStr: "202009.09.15"},
			"*202009.09.15", "Bad date format: first part doesn't look like a year",
		},
		{
			"Bad month length",
			args{inputStr: "2020.091.15"},
			"*2020.091.15", "Bad date format: second part doesn't look like a month",
		},
		{
			"Bad day length",
			args{inputStr: "2020.09.015"},
			"*2020.09.015", "Bad date format: third element doesn't look like a day",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, gotErrorMsg := NormalizeDate(tt.args.inputStr)
			if gotDate != tt.wantDate {
				t.Errorf("NormalizeDate() gotDate = %v, want %v", gotDate, tt.wantDate)
			}
			if gotErrorMsg != tt.wantErrorMsg {
				t.Errorf("NormalizeDate() gotErrorMsg = %v, want %v", gotErrorMsg, tt.wantErrorMsg)
			}
		})
	}
}

func TestIncrementDate(t *testing.T) {
	type args struct {
		date      string
		increment int
	}
	tests := []struct {
		name        string
		args        args
		wantNewdate string
		wantErr     string
	}{
		{
			"No date",
			args{date: "", increment: 2},
			"", "No date to increment",
		},
		{
			"increment below 0",
			args{date: "2020-09-05", increment: 0},
			"*2020-09-05", "Invalid day increment, expecting greater or equal to 1",
		},
		{
			"increment above 10",
			args{date: "2020-09-05", increment: 11},
			"*2020-09-05", "Invalid day increment, expecting smaller or equal to 10",
		},
		{
			"Invalid date",
			args{date: "2020-09-32", increment: 2},
			"*2020-09-32", "(Internal error) error parsing time \"2020-09-32\": day out of range",
		},
		{
			"happy case",
			args{date: "2020-09-05", increment: 2},
			"2020-09-07", "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewdate, gotErr := IncrementDate(tt.args.date, tt.args.increment)
			if gotNewdate != tt.wantNewdate {
				t.Errorf("IncrementDate() gotNewdate = %v, want %v", gotNewdate, tt.wantNewdate)
			}
			if gotErr != tt.wantErr {
				t.Errorf("IncrementDate() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
