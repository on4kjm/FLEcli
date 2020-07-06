package cmd

import (
	"reflect"
	"testing"
	"time"
)

func TestInferTimeBlock_full_happyCase(t *testing.T) {
	//Given
	recordNumber := 4

	logLine1 := LogLine{}
	logLine1.Date = "2020-05-24"
	logLine1.Time = "1401"
	logLine1.ActualTime = "1401"

	logLine2 := LogLine{}
	logLine2.Date = "2020-05-24"
	logLine2.Time = "1401"

	logLine3 := LogLine{}
	logLine3.Date = "2020-05-24"
	logLine3.Time = "1410"
	logLine3.ActualTime = "1410"

	//When
	tb := InferTimeBlock{}
	isEndGap, err := tb.storeTimeGap(logLine1, recordNumber)
	if isEndGap == true || err != nil {
		t.Error("Unexpected results processing logline 1")
	}

	isEndGap, err = tb.storeTimeGap(logLine2, recordNumber+1)
	if isEndGap == true || err != nil {
		t.Error("Unexpected results processing logline 2")
	}

	isEndGap, err = tb.storeTimeGap(logLine3, recordNumber+2)
	if isEndGap == false || err != nil {
		t.Error("Unexpected results processing logline 3")
	}

	err = tb.finalizeTimeGap()
	if err != nil {
		t.Errorf("Unexpected error finalizing the timeGap")
	}


	//Then
	expectedCount := 1
	if tb.noTimeCount != expectedCount {
		t.Errorf("Unexpected number of missing records: %d, expected %d", tb.noTimeCount, expectedCount)
	}

	expectedInterval := 270
	if tb.deltatime != expectedInterval {
		t.Errorf("Unexpected interval: %d, expected %d", tb.deltatime, expectedInterval)
	}

	expectedLastRecordedTime := time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
	if tb.lastRecordedTime != expectedLastRecordedTime {
		t.Errorf("Unexpected last recorded time: %s, expected %s", tb.lastRecordedTime, expectedLastRecordedTime)
	}

	expectedNextValidTime := time.Date(2020, time.May, 24, 14, 10, 0, 0, time.UTC)
	if tb.nextValidTime != expectedNextValidTime {
		t.Errorf("Unexpected last recorded time: %s, expected %s", tb.nextValidTime, expectedNextValidTime)
	}
}

func TestInferTimeBlock_display_happyCase(t *testing.T){
		//Given
		tb := InferTimeBlock{}
		tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
		tb.nextValidTime = time.Date(2020, time.May, 24, 14, 10, 10, 0, time.UTC)
		tb.noTimeCount = 1

		//When
		buffer1 := tb.displayTimeGapInfo()

		tb.finalizeTimeGap()

		buffer2 := tb.displayTimeGapInfo()

		//Then
		expectedBuffer1 := "Last Recorded Time:                 2020-05-24 14:01\nnext Recorded Time:                 2020-05-24 14:10\nLog position of last recorded time: 0\nNbr of entries without time:        1\nComputed interval:                  0\n"
		expectedBuffer2 := "Last Recorded Time:                 2020-05-24 14:01\nnext Recorded Time:                 2020-05-24 14:10\nLog position of last recorded time: 0\nNbr of entries without time:        1\nComputed interval:                  275\n"

		if buffer1 != expectedBuffer1 {
			t.Errorf("Not the expected display: got: \n%s\n while expecting: \n%s\n",buffer1,expectedBuffer1)
		}
		if buffer2 != expectedBuffer2 {
			t.Errorf("Not the expected finalized display: got: \n%s\n while expecting: \n%s\n",buffer2,expectedBuffer2)
		}
}

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
	tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 10, 0, 0, time.UTC)
	tb.nextValidTime = time.Date(2020, time.May, 24, 14, 01, 10, 0, time.UTC)

	//When
	err := tb.finalizeTimeGap()

	//Then
	if err == nil {
		t.Error("Should have failed with an error")
	}
	if err.Error() != "Fatal error: Gap start time is later than the Gap end time" {
		t.Errorf("Did not not fail with the expected error. Failed with %s", err)
	}
}

func TestInferTimeBlock_computeGaps_noDifference(t *testing.T) {
	//Given
	tb := InferTimeBlock{}
	tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 00, 0, 0, time.UTC)
	tb.nextValidTime = time.Date(2020, time.May, 24, 14, 00, 00, 0, time.UTC)

	//When
	err := tb.finalizeTimeGap()

	//Then
	if err == nil {
		t.Error("Should have failed with an error")
	}
	if err.Error() != "Fatal error: the start and end gap times are equal" {
		t.Errorf("Did not not fail with the expected error. Failed with %s", err)
	}
}

func TestInferTimeBlock_computeGaps_happyCase(t *testing.T) {
	//Given
	tb := InferTimeBlock{}
	tb.lastRecordedTime = time.Date(2020, time.May, 24, 14, 01, 0, 0, time.UTC)
	tb.nextValidTime = time.Date(2020, time.May, 24, 14, 10, 10, 0, time.UTC)
	tb.noTimeCount = 1

	//When
	err := tb.finalizeTimeGap()

	//Then
	if err != nil {
		t.Error("Should not have failed")
	}

	//TODO: add some other validation
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
