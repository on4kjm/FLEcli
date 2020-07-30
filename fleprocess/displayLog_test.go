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
	// ----       ---- ---- ---- ----       ---- ---- ----
}

