package cmd

import (
	"testing"
)

// // Test<NomDeLaFocntion>_<CasTesté>. Le cas testé doit compléter la phrase "it...", dans mon cas ça donerais "it
// // returns values".
// func TestAnalyze_ReturnsValues(t *testing.T) {
func TestInferTimeBlock_storeTimeGap(t *testing.T) {
	// Given (mise en place du test)
	input, err := ioutil.ReadFile("./myinput")
	if err != nil {
		t.Fatal(err)
	}
	// When (appel a ta fonction testée)
	got, err := Analyze(string(input))

	// Then (Assertions sur tes valeurs de retours).
	if err != nil {
		t.Fail("unexpected error", err)
		return
	}

	if len(got) == 0 {
		t.Fail("should have items")
	}

	// and so on...
}
