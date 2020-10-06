package fleprocess

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

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

//InferTimeBlock contains the information describing a time gap
type InferTimeBlock struct {
	lastRecordedTime time.Time
	nextValidTime    time.Time
	//Number of records without actual time
	noTimeCount int
	//Position in file of the first log entry with missing date
	logFilePosition int
	//Computed time interval
	deltatime time.Duration
}

//ADIFdateTimeFormat describes the ADIF date & time parsing and displaying format pattern
const ADIFdateTimeFormat = "2006-01-02 1504"

//displayTimeGapInfo will print the details stored in an InferTimeBlock
func (tb *InferTimeBlock) String() string {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("Last Recorded Time:                 %s\n", tb.lastRecordedTime.Format(ADIFdateTimeFormat)))
	buffer.WriteString(fmt.Sprintf("next Recorded Time:                 %s\n", tb.nextValidTime.Format(ADIFdateTimeFormat)))
	buffer.WriteString(fmt.Sprintf("Log position of last recorded time: %d\n", tb.logFilePosition))
	buffer.WriteString(fmt.Sprintf("Nbr of entries without time:        %d\n", tb.noTimeCount))
	buffer.WriteString(fmt.Sprintf("Computed interval:                  %ds\n", int(tb.deltatime.Seconds())))
	return buffer.String()
}

//finalizeTimeGap makes the necessary checks and computation
func (tb *InferTimeBlock) finalizeTimeGap() error {

	if err := tb.validateTimeGap(); err != nil {
		return err
	}

	//Compute the gap
	diff := tb.nextValidTime.Sub(tb.lastRecordedTime)
	tb.deltatime = time.Duration(diff / time.Duration(tb.noTimeCount+1))

	//Do we have a positive noTimeCount
	if tb.noTimeCount < 1 {
		return fmt.Errorf("Invalid number of records without time (%d)", tb.noTimeCount)
	}

	//TODO: What should we expect as logFilePosition?

	return nil
}

//validateTimeGap checks some important assumptions
func (tb *InferTimeBlock) validateTimeGap() error {
	//Check that lastRecordedTime and nextValidTime are not null
	if tb.lastRecordedTime.IsZero() {
		return errors.New("Gap start time is empty")
	}
	if tb.nextValidTime.IsZero() {
		return errors.New("Gap end time is empty")
	}

	//Fail if we have a negative time difference
	if tb.nextValidTime.Before(tb.lastRecordedTime) {
		return errors.New("Gap start time is later than the Gap end time")
	}
	return nil
}

//storeTimeGap updates an InferTimeBLock (last valid time, nbr of records without time). It returns true if we reached the end of the time gap.
func (tb *InferTimeBlock) storeTimeGap(logline LogLine, position int) (bool, error) {
	var err error

	//TODO: try to return fast and/or simpllify

	//ActualTime is filled if a time could be found in the FLE input
	if logline.ActualTime != "" {
		//Are we starting a new block
		if tb.noTimeCount == 0 {
			//File is bad: date not found or badly formated
			if logline.Date == "" {
				return false, errors.New("Date not defined or badly formated")
			}
			if tb.lastRecordedTime, err = time.Parse(ADIFdateTimeFormat, logline.Date+" "+logline.ActualTime); err != nil {
				log.Println("Fatal error during internal date conversion: ", err)
				os.Exit(1)
			}
			tb.logFilePosition = position
		} else {
			// We reached the end of the gap
			if tb.lastRecordedTime.IsZero() {
				return false, errors.New("Gap start time is empty")
			}
			if tb.nextValidTime, err = time.Parse(ADIFdateTimeFormat, logline.Date+" "+logline.ActualTime); err != nil {
				log.Println("Fatal error during internal date conversion: ", err)
				os.Exit(1)
			}
			return true, nil
		}
	} else {
		//Check the data is correct.
		if tb.lastRecordedTime.IsZero() {
			err = errors.New("Gap start time is empty")
			//TODO:  this smells
		}
		if !tb.nextValidTime.IsZero() {
			err = errors.New("Gap end time is not empty")
		}
		tb.noTimeCount++
	}
	return false, err
}
