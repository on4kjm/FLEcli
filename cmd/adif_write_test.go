package cmd

import "testing"

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
