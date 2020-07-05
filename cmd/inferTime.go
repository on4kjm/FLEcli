package cmd

import (
	"fmt"
	"log"
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

	//First log entry with missing date
	logFilePosition int

	//Computed time interval
	deltatime int
}

//finalizeTimeGap makes the necessary checks and computation
func (tb *InferTimeBlock) finalizeTimeGap() error {
	var err error
	err = nil

	//Check that lastRecordedTime and nextValidTime are not null
	nullTime := time.Time{}
	if tb.lastRecordedTime == nullTime {
		errMsg := "Fatal error: gap start time is empty"
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
		return err
	}
	if tb.nextValidTime == nullTime {
		errMsg := "Fatal error: gap end time is empty"
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
		return err
	}

	//Are the two times equal?
	if tb.nextValidTime == tb.lastRecordedTime {
		errMsg := "Fatal error: the start and end gap times are equal"
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
		return err
	}

	//Compute the time difference
	startTime := tb.lastRecordedTime
	endTime := tb.nextValidTime
	fmt.Println(startTime)
	fmt.Println(endTime)

	diff := endTime.Sub(startTime)
	fmt.Printf("Diff: %f min \n", diff.Minutes())

	//Fail if we have a negative time difference
	if diff.Minutes() < 0 {
		errMsg := "Fatal error: Gap start time is later than the Gap end time"
		log.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	//Do we have a non null noTimeCount
	if tb.noTimeCount < 1 {
		errMsg := fmt.Sprintf("Fatal error: invalid number of records without time (%d)\n", tb.noTimeCount)
		log.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	//TODO: What should we expect as logFilePosition?

	//Compute the gap
	floatInterval := diff.Minutes() / float64(tb.noTimeCount+1)
	intInterval := int(floatInterval)

	fmt.Printf("Interval: %f, Float interval: %f, Int interval: %d\n", diff.Minutes(), floatInterval, intInterval)

	tb.deltatime = intInterval

	return nil
}

//storeTimeGap updates an InferTimeBLock (last valid time, nbr of records without time). It returns true if we reached the end of the time gap.
func (tb *InferTimeBlock) storeTimeGap(logline LogLine, position int) (bool, error) {
	var err error
	err = nil

	//ActualTime is filled if a time could be found in the FLE input
	if logline.ActualTime != "" {
		//Are we starting a new block
		if tb.noTimeCount == 0 {
			tb.lastRecordedTime = convertDateTime(logline.Date + " " + logline.ActualTime)
			tb.logFilePosition = position
		} else {
			// We reached the end of the gap
			nullTime := time.Time{}
			if tb.lastRecordedTime == nullTime {
				errMsg := "Fatal error: gap start time is empty"
				log.Println(errMsg)
				err = fmt.Errorf(errMsg)
				return false, err
			}

			tb.nextValidTime = convertDateTime(logline.Date + " " + logline.ActualTime)
			return true, err
		}
	} else {
		//Check the data is correct.
		nullTime := time.Time{}
		if tb.lastRecordedTime == nullTime {
			errMsg := "Fatal error: gap start time is empty"
			log.Println(errMsg)
			err = fmt.Errorf(errMsg)
		}
		if tb.nextValidTime != nullTime {
			errMsg := "Fatal error: gap end time is not empty"
			log.Println(errMsg)
			err = fmt.Errorf(errMsg)
		}

		tb.noTimeCount++
	}
	return false, err
}

//convertDateTime converts the FLE date and time into a Go time structure
func convertDateTime(dateStr string) (fullDate time.Time) {
	const RFC3339FullDate = "2006-01-02 1504"

	date, err := time.Parse(RFC3339FullDate, dateStr)
	//error should never happen
	if err != nil {
		panic(err)
	}

	return date
}
