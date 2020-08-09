package fleprocess

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
	"os"
	"testing"
)

const testDir string = "test_dir"
const testFile string = "test.adi"

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	//create test directory
	os.Mkdir(testDir, os.FileMode(0522))
	//create test file
	f, _ := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	return func(t *testing.T) {
		t.Log("teardown test case")
		//delete test directory
		os.Remove(testDir)
		//detete test file
		os.Remove(testFile)
	}
}

func Test_buildOutputFilename(t *testing.T) {
	type args struct {
		output    string
		input     string
		overwrite bool
		extension string
	}
	tests := []struct {
		name               string
		args               args
		wantOutputFilename string
		wantWasOK          bool
	}{
		{
			"input file not provided",
			args{input: "", output: "xxx", overwrite: false, extension: ".adi"},
			"", false,
		},
		{
			"Output file does not exist",
			args{input: "a file", output: "output.adi", overwrite: false, extension: ".adi"},
			"output.adi", true,
		},
		{
			"Output name is a directory",
			args{input: "a file", output: testDir, overwrite: false, extension: ".adi"},
			"", false,
		},
		{
			"Output exist but no overwrite",
			args{input: "a file", output: testFile, overwrite: false, extension: ".adi"},
			"", false,
		},
		{
			"no output, input provided with extention",
			args{input: "/test/data/file.txt", output: "", overwrite: false, extension: ".adi"},
			"/test/data/file.adi", true,
		},
		{
			"no output, input provided without extention",
			args{input: "/test/data/file", output: "", overwrite: false, extension: ".adi"},
			"/test/data/file.adi", true,
		},
		{
			"no output, input provided, enfing with a point",
			args{input: "/test/data/file.", output: "", overwrite: false, extension: ".adi"},
			"/test/data/file.adi", true,
		},
	}

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutputFilename, gotWasOK := buildOutputFilename(tt.args.output, tt.args.input, tt.args.overwrite, tt.args.extension)
			if gotOutputFilename != tt.wantOutputFilename {
				t.Errorf("buildOutputFilename() gotOutputFilename = %v, want %v", gotOutputFilename, tt.wantOutputFilename)
			}
			if gotWasOK != tt.wantWasOK {
				t.Errorf("buildOutputFilename() gotWasOK = %v, want %v", gotWasOK, tt.wantWasOK)
			}
		})
	}
}
