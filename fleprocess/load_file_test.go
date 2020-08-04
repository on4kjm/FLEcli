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

func TestLoadFile_header_happyCase(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "{ Sample multi-line comment")
	dataArray = append(dataArray, "	( with quotes) Check: Logging > \"Contest Logging\"")
	dataArray = append(dataArray, "  - Data item1")
	dataArray = append(dataArray, "  - Data item2")
	dataArray = append(dataArray, "  }")
	dataArray = append(dataArray, "{ offset one liner comment }")
	dataArray = append(dataArray, "  { offset one liner comment }")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff onff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, "QslMsg This is a QSL message")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
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
		t.Errorf("Not the expected MyCall value: %s (expecting %s)", loadedLogFile[0].MyCall, expectedValue)
	}
		expectedValue = "ON4KJM"
	if loadedLogFile[0].Operator != expectedValue {
		t.Errorf("Not the expected Operator value: %s (expecting %s)", loadedLogFile[0].Operator, expectedValue)
	}
	expectedValue = "Portable"
	if loadedLogFile[0].Nickname != expectedValue {
		t.Errorf("Not the expected eQsl Nickname value: %s (expecting %s)", loadedLogFile[0].Nickname, expectedValue)
	}
	expectedValue = "ONFF-0258"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}
	expectedValue = "ON/ON-001"
	if loadedLogFile[0].MySOTA != expectedValue {
		t.Errorf("Not the expected MySOTA value: %s (expecting %s)", loadedLogFile[0].MySOTA, expectedValue)
	}
	expectedValue = "This is a QSL message"
	if loadedLogFile[0].QSLmsg != expectedValue {
		t.Errorf("Not the expected QSL Message from Header value: %s (expecting %s)", loadedLogFile[0].QSLmsg, expectedValue)
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
