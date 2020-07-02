package cmd

import "time"

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

	noTimeCount     int
	logFilePosition int

	//deltatime
}

//TODO: reset record
//TODO: finalize record (case no end, negative difference)

func storeTimeGap(logline LogLine, position int, timeBlock InferTimeBlock) {
	//ActualTime is filled if a time could be found
	if logline.ActualTime != "" {
		//store the date and time
		timeBlock.lastRecordedTime = convertDateTime(logline.Date + " " + logline.ActualTime)
		timeBlock.noTimeCount = 0
	} else {
		//if the noTimeCount is 0, store the position in the logfile
		//increment the no time caount
	}
}
