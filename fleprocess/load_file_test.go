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
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadFile_happyCase(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if !isLoadedOK {
		t.Error("Test file could not be correctly processed")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "ON4KJM/P"
	if loadedLogFile[0].MyCall != expectedValue {
		t.Errorf("Not the expected value: %s (expecting %s)", loadedLogFile[0].MyCall, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

//createTestFile creates and populates a test FLE input file. 
//Returns the created temporary filename. 
func createTestFile(dataArray []string) (tempFileName string) {
	//create random file name
	tmpfile, err := ioutil.TempFile("", "*.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Temporary file: %s", tmpfile.Name())

	//Write the passed data to the file
	writeFile(tmpfile.Name(), dataArray)

	//Return the temporaty filename
	return tmpfile.Name()
}
