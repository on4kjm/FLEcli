package cmd

import (
	"reflect"
	"testing"
	"time"
)

func Test_storeTimeGap(t *testing.T) {
	type args struct {
		logline   LogLine
		position  int
		timeBlock InferTimeBlock
	}
	tests := []struct {
		name             string
		args             args
		wantNewTimeBlock InferTimeBlock
	}{
		{
			"Time defined in FLE log",
			args{LogLine{Date: "2020-05-24", Time: "2312", ActualTime: "2312"}, 1, InferTimeBlock{}},
			InferTimeBlock{lastRecordedTime: time.Date(2020, time.May, 24, 23, 12, 0, 0, time.UTC), noTimeCount: 0, logFilePosition: 1},
		},
		{
			"Time is not defined in FLE log",
			args{LogLine{Date: "2020-05-24", Time: "2312", ActualTime: ""}, 1, InferTimeBlock{lastRecordedTime: time.Date(2020, time.May, 24, 23, 12, 0, 0, time.UTC), noTimeCount: 3, logFilePosition: 1}},
			InferTimeBlock{lastRecordedTime: time.Date(2020, time.May, 24, 23, 12, 0, 0, time.UTC), noTimeCount: 4, logFilePosition: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewTimeBlock := storeTimeGap(tt.args.logline, tt.args.position, tt.args.timeBlock); !reflect.DeepEqual(gotNewTimeBlock, tt.wantNewTimeBlock) {
				t.Errorf("storeTimeGap() = %v, want %v", gotNewTimeBlock, tt.wantNewTimeBlock)
			}
		})
	}
}

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
