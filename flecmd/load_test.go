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
	"FLEcli/fleprocess"
	"bytes"
	"fmt"
	"os"
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
	if !strings.HasPrefix(string(out), expectedOutputStart) {
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
	if !strings.HasPrefix(string(out), expectedOutputStart) {
		t.Fatalf("expected to start with \"%s\" got \"%s\"", expectedOutputStart, string(out))
	}
}

func Test_ExecuteCommand_happyCase(t *testing.T) {
	processLoadFile = mockLoadFile

	//Capture output
	rescueStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

	cmd := loadCmdConstructor()

	cmd.SetArgs([]string{"data.txt"})
	cmdErr := cmd.Execute()

	//Close the capture and get the data
	w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = rescueStdout

	if cmdErr != nil {
		t.Fatalf("Unexpected error executing command: %s", cmdErr)
	}

	if string(out) != "fileLoad via mock" {
		t.Fatalf("Expected \"fileLoad via mock\". Got \"%s\"", string(out))
	}
}

func mockLoadFile(inputFilename string, isInterpolateTime bool) (filleFullLog []fleprocess.LogLine, isProcessedOK bool) {
	fmt.Print("fileLoad via mock")
	return nil, true
}
