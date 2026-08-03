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

import "testing"

func TestCountChronologicalReversals(t *testing.T) {
	fullLog := []LogLine{
		{Date: "2026-08-01", Time: "2011"},
		{Date: "2026-08-01", Time: "2010"},
		{Date: "2026-08-01", Time: "2013"},
		{Date: "2026-08-01", Time: "2119"},
		{Date: "2026-08-01", Time: "2115"},
	}

	if got, want := countChronologicalReversals(fullLog), 2; got != want {
		t.Fatalf("countChronologicalReversals() = %d, want %d", got, want)
	}
}

func TestSortLogChronologicallyIsStable(t *testing.T) {
	fullLog := []LogLine{
		{Date: "2026-08-01", Time: "2151", Call: "AE6Z"},
		{Date: "2026-08-01", Time: "2149", Call: "NW7E"},
		{Date: "2026-08-01", Time: "2151", Call: "SAME1"},
		{Date: "2026-07-31", Time: "2359", Call: "PREVIOUS"},
	}

	if got, want := sortLogChronologically(fullLog), 3; got != want {
		t.Fatalf("sortLogChronologically() reordered %d records, want %d", got, want)
	}

	wantCalls := []string{"PREVIOUS", "NW7E", "AE6Z", "SAME1"}
	for i, want := range wantCalls {
		if got := fullLog[i].Call; got != want {
			t.Errorf("fullLog[%d].Call = %q, want %q", i, got, want)
		}
	}
}
