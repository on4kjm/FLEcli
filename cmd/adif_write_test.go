package cmd

import (
	"reflect"
	"testing"
)

func Test_adifElement(t *testing.T) {
	type args struct {
		elementName  string
		elementValue string
	}
	tests := []struct {
		name        string
		args        args
		wantElement string
	}{
		{
			"case 1",
			args{elementName: "station_callsign", elementValue: "ON4KJM/P"},
			"<STATION_CALLSIGN:8>ON4KJM/P ",
		},
		{
			"case 2",
			args{elementName: "time_ON", elementValue: "1310"},
			"<TIME_ON:4>1310 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotElement := adifElement(tt.args.elementName, tt.args.elementValue); gotElement != tt.wantElement {
				t.Errorf("adifElement() = %v, want %v", gotElement, tt.wantElement)
			}
		})
	}
}

func Test_buildAdif(t *testing.T) {
	sampleFilledLog1 := []LogLine{
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "20200524", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", MyWWFF: "ONFF-0259", Operator: "ON4KJM", Nickname: "ONFF-0259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "20200524", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MyWWFF: "ONFF-0259", Operator: "ON4KJM"},
	}

	expectedOutput1 := []string{
		"ADIF Export for Fast Log Entry by DF3CB",
		"<PROGRAMID:3>FLE",
		"<ADIF_VER:5>3.0.6",
		"<EOH>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>S57LC <QSO_DATE:8>20200524 <TIME_ON:4>1310 <BAND:3>20m <MODE:2>CW <FREQ:6>14.045 <RST_SENT:3>599 <RST_RCVD:3>599 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <APP_EQSL_QTH_NICKNAME:11>ONFF-0259-1 <EOR>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>ON4LY <QSO_DATE:8>20200524 <TIME_ON:4>1312 <BAND:3>20m <MODE:2>CW <RST_SENT:3>559 <RST_RCVD:3>599 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <EOR>",
	}

	type args struct {
		fullLog []LogLine
	}
	tests := []struct {
		name         string
		args         args
		wantAdifList []string
	}{
		{
			"Happy case",
			args{fullLog: sampleFilledLog1},
			expectedOutput1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdifList := buildAdif(tt.args.fullLog); !reflect.DeepEqual(gotAdifList, tt.wantAdifList) {
				t.Errorf("buildAdif() = %v, want %v", gotAdifList, tt.wantAdifList)
			}
		})
	}
}
