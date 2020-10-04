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
	//"fmt"
	"io/ioutil"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
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
		t.Fatalf("expected \"%s\" got \"%s\"", "hi", string(out))
	}
}