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
	"bufio"
	"os"
	"testing"
)

const WriteFileTestDir string = "test2_dir"
const writeFileTestFname string = "testFile.txt"

func Test_writeFile(t *testing.T) {

	dataArray := make([]string, 0)
	dataArray = append(dataArray, "foo")
	dataArray = append(dataArray, "bar")

	writeFile(writeFileTestFname, dataArray)

	//Open and read the file we have just created
	file, err := os.Open(writeFileTestFname)

	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var readLines []string
	for scanner.Scan() {
		readLines = append(readLines, scanner.Text())
	}
	if error := scanner.Err(); error != nil {
		t.Error(error)
	}
	file.Close()

	//Compare with what we have got
	if len(dataArray) != len(readLines) {
		t.Error("The number of lines read doesn't match the lines written")
	}
	for i, v := range readLines {
		if v != dataArray[i] {
			t.Error("Didn't read the expected data")
		}
	}

	// //detete test file
	os.Remove(writeFileTestFname)

}
