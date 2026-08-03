package fleprocess

/*
Copyright © 2026 Jean-Marc Meessen, ON4KJM <on4kjm@gmail.com>

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

import "sort"

// countChronologicalReversals counts adjacent QSO transitions where the date
// and time move backwards.
func countChronologicalReversals(fullLog []LogLine) int {
	reversals := 0
	for i := 1; i < len(fullLog); i++ {
		if logLineChronologicallyLess(fullLog[i], fullLog[i-1]) {
			reversals++
		}
	}
	return reversals
}

// sortLogChronologically stably sorts QSOs by date and time. It returns the
// number of records whose position changed. Stable sorting preserves the input
// order of QSOs with identical dates and times.
func sortLogChronologically(fullLog []LogLine) int {
	type indexedLogLine struct {
		logLine       LogLine
		originalIndex int
	}

	indexedLog := make([]indexedLogLine, len(fullLog))
	for i, logLine := range fullLog {
		indexedLog[i] = indexedLogLine{logLine: logLine, originalIndex: i}
	}

	sort.SliceStable(indexedLog, func(i, j int) bool {
		return logLineChronologicallyLess(indexedLog[i].logLine, indexedLog[j].logLine)
	})

	reorderedRecords := 0
	for i, item := range indexedLog {
		if item.originalIndex != i {
			reorderedRecords++
		}
		fullLog[i] = item.logLine
	}

	return reorderedRecords
}

func logLineChronologicallyLess(left, right LogLine) bool {
	if left.Date != right.Date {
		return left.Date < right.Date
	}
	return left.Time < right.Time
}
