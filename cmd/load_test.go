package cmd

import (
	"reflect"
	"testing"
	"time"
)

func Test_convertDateTime(t *testing.T) {
	type args struct {
		dateStr string
		timeStr string
	}
	tests := []struct {
		name         string
		args         args
		wantFullDate time.Time
	}{
		{
			"case 1",
			args{dateStr: "2020-05-24 2312"},
			time.Date(2020, time.May, 24, 23, 12, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFullDate := convertDateTime(tt.args.dateStr); !reflect.DeepEqual(gotFullDate, tt.wantFullDate) {
				t.Errorf("convertDateTime() = %v, want %v", gotFullDate, tt.wantFullDate)
			}
		})
	}
}
