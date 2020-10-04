package flecmd

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

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"io/ioutil"
	"testing"
)

func Test_ExecuteCommand_help(t *testing.T) {
	cmd := loadCmdConstructor()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := "Loads and validates a FLE type shorthand logfile\n\nUsage:\n  load [flags] inputFile\n\nFlags:\n  -h, --help   help for load\n"
	if string(out) != expectedOutput {
		t.Fatalf("expected \"%s\" got \"%s\"", expectedOutput, string(out))
	}
}

func Test_ExecuteCommand_noArgs(t *testing.T) {
	cmd := loadCmdConstructor()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	//cmd.SetArgs([]string{""})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	//FIXME: doesn't work as espected
	expectedOutputStart := "Error: Missing input file \nUsage:"
	if !strings.HasPrefix(string(out), expectedOutputStart){
		t.Fatalf("expected to start with \"%s\" got \"%s\"", expectedOutputStart, string(out))
	}
}

func Test_ExecuteCommand_toManyArgs(t *testing.T) {
	cmd := loadCmdConstructor()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"blaah", "blaah", "blaah"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	
	expectedOutputStart := "Error: Too many arguments.\nUsage"
	if !strings.HasPrefix(string(out), expectedOutputStart){
		t.Fatalf("expected to start with \"%s\" got \"%s\"", expectedOutputStart, string(out))
	}
}

func Test_ExecuteCommand_happyCase(t *testing.T) {
	cmd := loadCmdConstructor()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	
	cmd.SetArgs([]string{"../test/data/fle-1.txt"})
	cmdErr := cmd.Execute()
	if cmdErr != nil {
		t.Fatalf("Unexpected error executing command: %s",cmdErr)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "" {
		t.Fatalf("No output at this level was expected. Got %s", string(out))
	}
}

func Test_ExecuteCommand_bigFile(t *testing.T) {
	cmd := loadCmdConstructor()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	
	cmd.SetArgs([]string{"../test/data/fle-6-bigFile.txt"})
	isInterpolateTime = true

	PrintMemUsage()
	cmdErr := cmd.Execute()
	if cmdErr != nil {
		t.Fatalf("Unexpected error executing command: %s",cmdErr)
	}
	PrintMemUsage()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "" {
		t.Fatalf("No output at this level was expected. Got %s", string(out))
	}

}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}