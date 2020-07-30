package fleprocess

import (
	"fmt"
	"testing"
)

func TestSprintHeaderValues(t *testing.T) {
	type args struct {
		logLine LogLine
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Full Option",
			args{logLine: LogLine{MyCall: "on4kjm/p", Operator: "on4kjm", MyWWFF: "wwff", MySOTA: "sota"}},
			"MyCall    on4kjm/p (on4kjm)\nMyWWFF    wwff\nMySOTA    sota\n",
		},
		{
			"Minimal",
			args{logLine: LogLine{MyCall: "on4kjm/p"}},
			"MyCall    on4kjm/p\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SprintHeaderValues(tt.args.logLine); got != tt.want {
				t.Errorf("SprintHeaderValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSprintColumnTitles() {
	out := SprintColumnTitles()
	fmt.Print(out)
	//Output:
	//Date       Time Band Mode Call       Sent Rcvd Notes
	//----       ---- ---- ---- ----       ---- ---- ----
}

func ExampleSprintLogRecord() {
	logLine := LogLine{
		Date:             "date",
		MyCall:           "myCall",
		Operator:         "operator",
		MyWWFF:           "myWwff",
		MySOTA:           "mySota",
		QslMsgFromHeader: "QslMsgFromHeader",
		Nickname:         "nickname",
		Mode:             "mode",
		ModeType:         "modeType",
		Band:             "band",
		BandLowerLimit:   1.0,
		BandUpperLimit:   2.0,
		Frequency:        "frequency",
		Time:             "time",
		Call:             "call",
		Comment:          "comment",
		QSLmsg:           "qslMessage",
		OMname:           "omName",
		GridLoc:          "gridLoc",
		RSTsent:          "rstSent",
		RSTrcvd:          "rstRcvd",
		SOTA:             "sota",
		WWFF:             "wwff",
	}
	out := SprintLogRecord(logLine)
	fmt.Print(out)

	//output:
	//Date      date
	//MyCall    myCall
	//Operator  operator
	//MyWWFF    myWwff
	//MySOTA    mySota
	//QslMsg    QslMsgFromHeader
	//Nickname  nickname
	//Mode      mode
	//ModeType  modeType
	//Band      band
	//   Lower   1.000000
	//   Upper   2.000000
	//Frequency frequency
	//Time      time
	//Call      call
	//Comment   comment
	//QSLmsg    qslMessage
	//OMname    omName
	//GridLoc   gridLoc
	//RSTsent   rstSent
	//RSTrcvd   rstRcvd
	//SOTA      sota
	//WWFF      wwff

}

func TestSprintLogInColumn(t *testing.T) {
	type args struct {
		logLine LogLine
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			"Full Option",
			args{logLine: LogLine{
				Date: "date",
				MyCall:           "myCall",
				Operator:         "operator",
				MyWWFF:           "myWwff",
				MySOTA:           "mySota",
				QslMsgFromHeader: "QslMsgFromHeader",
				Nickname:         "nickname",
				Mode:             "mode",
				ModeType:         "modeType",
				Band:             "band",
				BandLowerLimit:   1.0,
				BandUpperLimit:   2.0,
				Frequency:        "frequency",
				Time:             "time",
				Call:             "call",
				Comment:          "comment",
				QSLmsg:           "qslMessage",
				OMname:           "omName",
				GridLoc:          "gridLoc",
				RSTsent:          "rstSent",
				RSTrcvd:          "rstRcvd",
				SOTA:             "sota",
				WWFF:             "wwff",},
			},
			"date       time band mode call       rstSent rstRcvd QRG: frequency [comment] [qslMessage] omName gridLoc wwff sota \n",
		},
		{
			"Minimal",
			args{logLine: LogLine{
				Date: "date",
				MyCall:           "myCall",
				Operator:         "operator",
				MyWWFF:           "myWwff",
				MySOTA:           "mySota",
				QslMsgFromHeader: "QslMsgFromHeader",
				Nickname:         "nickname",
				Mode:             "mode",
				ModeType:         "modeType",
				Band:             "band",
				BandLowerLimit:   1.0,
				BandUpperLimit:   2.0,
				Time:             "time",
				Call:             "call",
				RSTsent:          "rstSent",
				RSTrcvd:          "rstRcvd",},
			},
			"date       time band mode call       rstSent rstRcvd \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := SprintLogInColumn(tt.args.logLine); gotOutput != tt.wantOutput {
				t.Errorf("SprintLogInColumn() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
