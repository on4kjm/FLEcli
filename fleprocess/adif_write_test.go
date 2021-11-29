package fleprocess

/*
Copyright © 2020 Jean-Marc Meessen, ON4KJM <on4kjm@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "2020-05-24", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", MyWWFF: "ONFF-0259", Operator: "ON4KJM", Nickname: "ONFF-0259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "2020-05-24", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MyWWFF: "ONFF-0259", Operator: "ON4KJM"},
	}

	expectedOutput1 := []string{
		"ADIF Export for Fast Log Entry by DF3CB",
		"<PROGRAMID:3>FLE",
		"<ADIF_VER:5>3.1.0",
		"<EOH>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>S57LC <QSO_DATE:8>20200524 <TIME_ON:4>1310 <BAND:3>20m <MODE:2>CW <FREQ:6>14.045 <RST_SENT:3>599 <RST_RCVD:3>599 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <APP_EQSL_QTH_NICKNAME:11>ONFF-0259-1 <EOR>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>ON4LY <QSO_DATE:8>20200524 <TIME_ON:4>1312 <BAND:3>20m <MODE:2>CW <RST_SENT:3>559 <RST_RCVD:3>599 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <EOR>",
	}

	sampleFilledLog2 := []LogLine{
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "2020-05-24", MyGrid: "JO40eu", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", GridLoc: "JO50", MyWWFF: "ONFF-0259", Operator: "ON4KJM", Nickname: "ONFF-0259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "2020-05-24", MyGrid: "JO40eu", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MyWWFF: "ONFF-0259", Operator: "ON4KJM"},
	}

	expectedOutput2 := []string{
		"ADIF Export for Fast Log Entry by DF3CB",
		"<PROGRAMID:3>FLE",
		"<ADIF_VER:5>3.1.0",
		"<EOH>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>S57LC <QSO_DATE:8>20200524 <TIME_ON:4>1310 <BAND:3>20m <MODE:2>CW <FREQ:6>14.045 <RST_SENT:3>599 <RST_RCVD:3>599 <GRIDSQUARE:4>JO50 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <MY_GRIDSQUARE:6>JO40eu <APP_EQSL_QTH_NICKNAME:11>ONFF-0259-1 <EOR>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>ON4LY <QSO_DATE:8>20200524 <TIME_ON:4>1312 <BAND:3>20m <MODE:2>CW <RST_SENT:3>559 <RST_RCVD:3>599 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <MY_GRIDSQUARE:6>JO40eu <EOR>",
	}

	sampleFilledLog3 := []LogLine{
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "2020-05-24", MyGrid: "JO40eu", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", GridLoc: "JO50", MyWWFF: "ONFF-0259", Operator: "ON4KJM", Nickname: "ONFF-0259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "2020-05-24", MyGrid: "JO40eu", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MyWWFF: "ONFF-0259", Operator: "ON4KJM", WWFF: "DLFF-0001"},
	}

	expectedOutput3 := []string{
		"ADIF Export for Fast Log Entry by DF3CB",
		"<PROGRAMID:3>FLE",
		"<ADIF_VER:5>3.1.0",
		"<EOH>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>S57LC <QSO_DATE:8>20200524 <TIME_ON:4>1310 <BAND:3>20m <MODE:2>CW <FREQ:6>14.045 <RST_SENT:3>599 <RST_RCVD:3>599 <GRIDSQUARE:4>JO50 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <OPERATOR:6>ON4KJM <MY_GRIDSQUARE:6>JO40eu <APP_EQSL_QTH_NICKNAME:11>ONFF-0259-1 <EOR>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>ON4LY <QSO_DATE:8>20200524 <TIME_ON:4>1312 <BAND:3>20m <MODE:2>CW <RST_SENT:3>559 <RST_RCVD:3>599 <MY_SIG:4>WWFF <MY_SIG_INFO:9>ONFF-0259 <SIG:4>WWFF <SIG_INFO:9>DLFF-0001 <OPERATOR:6>ON4KJM <MY_GRIDSQUARE:6>JO40eu <EOR>",
	}

	sampleFilledLog_POTA := []LogLine{
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "2020-05-24", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", MyPOTA: "ON-00259", Operator: "ON4KJM", Nickname: "ON-00259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "2020-05-24", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MyPOTA: "ON-00259", Operator: "ON4KJM"},
	}

	expectedOutput_POTA := []string{
		"ADIF Export for Fast Log Entry by DF3CB",
		"<PROGRAMID:3>FLE",
		"<ADIF_VER:5>3.1.0",
		"<EOH>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>S57LC <QSO_DATE:8>20200524 <TIME_ON:4>1310 <BAND:3>20m <MODE:2>CW <FREQ:6>14.045 <RST_SENT:3>599 <RST_RCVD:3>599 <MY_SIG:4>POTA <MY_SIG_INFO:8>ON-00259 <OPERATOR:6>ON4KJM <APP_EQSL_QTH_NICKNAME:10>ON-00259-1 <EOR>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>ON4LY <QSO_DATE:8>20200524 <TIME_ON:4>1312 <BAND:3>20m <MODE:2>CW <RST_SENT:3>559 <RST_RCVD:3>599 <MY_SIG:4>POTA <MY_SIG_INFO:8>ON-00259 <OPERATOR:6>ON4KJM <EOR>",
	}


	sampleFilledLog_POTA2 := []LogLine{
		{MyCall: "ON4KJM/P", Call: "S57LC", Date: "2020-05-24", MyGrid: "JO40eu", Time: "1310", Band: "20m", Frequency: "14.045", Mode: "CW", RSTsent: "599", RSTrcvd: "599", GridLoc: "JO50", MyPOTA: "ON-00259", Operator: "ON4KJM", Nickname: "ON-00259-1"},
		{MyCall: "ON4KJM/P", Call: "ON4LY", Date: "2020-05-24", MyGrid: "JO40eu", Time: "1312", Band: "20m", Mode: "CW", RSTsent: "559", RSTrcvd: "599", MyPOTA: "ON-00259", Operator: "ON4KJM", POTA: "DL-00001"},
	}

	expectedOutput_POTA2 := []string{
		"ADIF Export for Fast Log Entry by DF3CB",
		"<PROGRAMID:3>FLE",
		"<ADIF_VER:5>3.1.0",
		"<EOH>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>S57LC <QSO_DATE:8>20200524 <TIME_ON:4>1310 <BAND:3>20m <MODE:2>CW <FREQ:6>14.045 <RST_SENT:3>599 <RST_RCVD:3>599 <GRIDSQUARE:4>JO50 <MY_SIG:4>POTA <MY_SIG_INFO:8>ON-00259 <OPERATOR:6>ON4KJM <MY_GRIDSQUARE:6>JO40eu <APP_EQSL_QTH_NICKNAME:10>ON-00259-1 <EOR>",
		"<STATION_CALLSIGN:8>ON4KJM/P <CALL:5>ON4LY <QSO_DATE:8>20200524 <TIME_ON:4>1312 <BAND:3>20m <MODE:2>CW <RST_SENT:3>559 <RST_RCVD:3>599 <MY_SIG:4>POTA <MY_SIG_INFO:8>ON-00259 <SIG:4>POTA <SIG_INFO:8>DL-00001 <OPERATOR:6>ON4KJM <MY_GRIDSQUARE:6>JO40eu <EOR>",
	}

	type args struct {
		fullLog []LogLine
		adifParams AdifParams
	}
	tests := []struct {
		name         string
		args         args
		wantAdifList []string
	}{
		{
			"Happy case-WWFF",
			args{
				fullLog: sampleFilledLog1, 
				adifParams: AdifParams{IsWWFF: true, IsSOTA: false},
			},
			expectedOutput1,
		},
		{
			"Happy case-POTA", 
			args{
				fullLog: sampleFilledLog_POTA, 
				adifParams: AdifParams{IsPOTA: true},
			},
			expectedOutput_POTA,
		},
		{
			"Happy case-Grid",
			args{fullLog: sampleFilledLog2, 
			adifParams: AdifParams{IsWWFF: true, IsSOTA: false},
			},
			expectedOutput2,
		},
		{
			"Happy case-WWFF2WWFF",
			args{
				fullLog: sampleFilledLog3, 
				adifParams: AdifParams{IsWWFF: true},
			},
			expectedOutput3,
		},
		{
			"Happy case-POTA2POTA",
			args{
				fullLog: sampleFilledLog_POTA2, 
				adifParams: AdifParams{IsPOTA: true},
			},
			expectedOutput_POTA2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAdifList := buildAdif(tt.args.fullLog, tt.args.adifParams); !reflect.DeepEqual(gotAdifList, tt.wantAdifList) {
				t.Errorf("buildAdif() = %v, want %v", gotAdifList, tt.wantAdifList)
			}
		})
	}
}

func Test_adifDate(t *testing.T) {
	type args struct {
		inputDate string
	}
	tests := []struct {
		name           string
		args           args
		wantOutputDate string
	}{
		{
			"Happy case",
			args{inputDate: "2020-06-13"},
			"20200613",
		},
		//Panics as expected but I don't know how to test this.
		// {
		// 	"Bad format",
		// 	args{inputDate: "2020-13-06"},
		// 	"20200613",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutputDate := adifDate(tt.args.inputDate); gotOutputDate != tt.wantOutputDate {
				t.Errorf("adifDate() = %v, want %v", gotOutputDate, tt.wantOutputDate)
			}
		})
	}
}
