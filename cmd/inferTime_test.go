package cmd

import (
	"reflect"
	"testing"
	"time"
)

func TestInferTimeBlock_computeGaps_invalidData(t *testing.T) {
	//Given
	tb := InferTimeBlock{}

	//When
	err := tb.finalizeTimeGap()

	//Then
	if err == nil {
		t.Error("Should have failed with an error")
	}
	if err.Error() != "Fatal error: gap start time is empty" {
		t.Error("Did not not fail with the expected error.")
	}
}

func TestInferTimeBlock_computeGaps_missingEnTime(t *testing.T) {
	//Given
	tb := InferTimeBlock{}
	tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)

	//When
	err := tb.finalizeTimeGap()

	//Then
	if err == nil {
		t.Error("Should have failed with an error")
	}
	if err.Error() != "Fatal error: gap end time is empty" {
		t.Errorf("Did not not fail with the expected error. Failed with %s", err)
	}
}

func TestInferTimeBlock_computeGaps_negativeDifference(t *testing.T) {
	//Given
	tb := InferTimeBlock{}

	//When
	err := tb.finalizeTimeGap()

	//Then
	if err == nil {
		t.Error("Should have failed with an error")
	}
}

func TestInferTimeBlock_startsNewBlock(t *testing.T) {
	// Given
	logLine := LogLine{}
	logLine.Date = "2020-05-24"
	logLine.Time = "1401"
	logLine.ActualTime = "1401"

	recordNbr := 4

	tb := InferTimeBlock{}

	// When
	isEndGap, err := tb.storeTimeGap(logLine, recordNbr)

	// Then
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if isEndGap == true {
		t.Errorf("Result is true while expectig false")
	}
	if tb.lastRecordedTime != time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC) {
		t.Errorf("Not the expected lastRecordedTime")
	}
	if tb.noTimeCount != 0 {
		t.Errorf("nTimeCount should be 0, but is %d", tb.noTimeCount)
	}
	if tb.logFilePosition != recordNbr {
		t.Errorf("logFilePosition not set correctly: is %d while expecting %d", tb.logFilePosition, recordNbr)
	}
}

func TestInferTimeBlock_incrementCounter(t *testing.T) {
	// Given
	logLine := LogLine{}
	logLine.Date = "2020-05-24"
	logLine.Time = "1401"

	recordNbr := 4

	tb := InferTimeBlock{}
	tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
	tb.logFilePosition = recordNbr

	// When
	isEndGap, err := tb.storeTimeGap(logLine, recordNbr)

	// Then
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if isEndGap == true {
		t.Errorf("Result is true while expectig false")
	}
	if tb.lastRecordedTime != time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC) {
		t.Errorf("Not the expected lastRecordedTime")
	}
	if tb.noTimeCount != 1 {
		t.Errorf("nTimeCount should be 1, but is %d", tb.noTimeCount)
	}
	if tb.logFilePosition != recordNbr {
		t.Errorf("logFilePosition not set correctly: is %d while expecting %d", tb.logFilePosition, recordNbr)
	}
}

func TestInferTimeBlock_increment_missingLastTime(t *testing.T) {
	// Given
	logLine := LogLine{}
	logLine.Date = "2020-05-24"
	logLine.Time = "1401"

	recordNbr := 4

	tb := InferTimeBlock{}
	//tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
	tb.logFilePosition = recordNbr

	// When
	isEndGap, err := tb.storeTimeGap(logLine, recordNbr)

	// Then
	if err.Error() != "Fatal error: gap start time is empty" {
		t.Errorf("Unexpected error: %s", err)
	}
	if isEndGap == true {
		t.Errorf("Result is true while expectig false")
	}
}

func TestInferTimeBlock_increment_alreadyDefinedNewTime(t *testing.T) {
	// Given
	logLine := LogLine{}
	logLine.Date = "2020-05-24"
	logLine.Time = "1401"

	recordNbr := 4

	tb := InferTimeBlock{}
	tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
	tb.nextValidTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
	tb.logFilePosition = recordNbr

	// When
	isEndGap, err := tb.storeTimeGap(logLine, recordNbr)

	// Then
	if err.Error() != "Fatal error: gap end time is not empty" {
		t.Errorf("Unexpected error: %s", err)
	}
	if isEndGap == true {
		t.Errorf("Result is true while expectig false")
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
