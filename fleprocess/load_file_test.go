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
	"os"
	"testing"
)

func TestLoadFile_happyCase(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "{ Sample multi-line comment")
	dataArray = append(dataArray, "	( with quotes) Check: Logging > \"Contest Logging\"")
	dataArray = append(dataArray, "  - Data item1")
	dataArray = append(dataArray, "  - Data item2")
	dataArray = append(dataArray, "  }")
	dataArray = append(dataArray, "{ one liner comment }")
	dataArray = append(dataArray, "  { offset one liner comment }")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator\t on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff\tonff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, "myPota k-0802")
	dataArray = append(dataArray, "myGrid jo50")
	dataArray = append(dataArray, "myLat 55.5555555")
	dataArray = append(dataArray, "myLon -120.01234567")
	dataArray = append(dataArray, "myCounty Ham County")
	dataArray = append(dataArray, "QslMsg This is a QSL message")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "0954 on4do")

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
	expectedValue = "55.5555555"
	if loadedLogFile[0].MyLat != expectedValue {
		t.Errorf("Not the expected MyLat value: %s (expecting %s)", loadedLogFile[0].MyLat, expectedValue)
	}
	expectedValue = "-120.01234567"
	if loadedLogFile[0].MyLon != expectedValue {
		t.Errorf("Not the expected MyLon value: %s (expecting %s)", loadedLogFile[0].MyLon, expectedValue)
	}
	expectedValue = "This is a QSL message"
	if loadedLogFile[0].QSLmsg != expectedValue {
		t.Errorf("Not the expected QSL Message from Header value: %s (expecting %s)", loadedLogFile[0].QSLmsg, expectedValue)
	}
	expectedValue = "Ham County"
	if loadedLogFile[0].MyCounty != expectedValue {
		t.Errorf("Not the expected MyCounty from Header value: %s (expecting %s)", loadedLogFile[0].MyCounty, expectedValue)
	}
	expectedValue = "IK5ZVE/5"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_happyCase_date(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "{ Sample multi-line comment")
	dataArray = append(dataArray, "	( with quotes) Check: Logging > \"Contest Logging\"")
	dataArray = append(dataArray, "  - Data item1")
	dataArray = append(dataArray, "  - Data item2")
	dataArray = append(dataArray, "  }")
	dataArray = append(dataArray, "{ one liner comment }")
	dataArray = append(dataArray, "  { offset one liner comment }")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "0954 on4do")

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
	expectedValue = "IK5ZVE/5"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_happyCase_date2(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "0954 on4do")
	dataArray = append(dataArray, "20-05-25 20m ssb 1000 on4up")

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
	expectedValue = "IK5ZVE/5"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}

	//	 "20-05-25 20m 1000 on4up")
	expectedValue = "2020-05-25"
	if loadedLogFile[3].Date != expectedValue {
		t.Errorf("Not the expected Date[3] value: %s (expecting %s)", loadedLogFile[3].Time, expectedValue)
	}
	expectedValue = "1000"
	if loadedLogFile[3].Time != expectedValue {
		t.Errorf("Not the expected Time[3] value: %s (expecting %s)", loadedLogFile[3].Time, expectedValue)
	}
	expectedValue = "20m"
	if loadedLogFile[3].Band != expectedValue {
		t.Errorf("Not the expected Band[3] value: %s (expecting %s)", loadedLogFile[3].Band, expectedValue)
	}
	expectedValue = "ON4UP"
	if loadedLogFile[3].Call != expectedValue {
		t.Errorf("Not the expected Call[3] value: %s (expecting %s)", loadedLogFile[3].Call, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_happyCase_day(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "0954 on4do")
	dataArray = append(dataArray, "day ++ 20m 1000 on4up")

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
	expectedValue = "IK5ZVE/5"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}

	//	 "20-05-25 20m 1000 on4up")
	expectedValue = "2020-05-25"
	if loadedLogFile[3].Date != expectedValue {
		t.Errorf("Not the expected Date[3] value: %s (expecting %s)", loadedLogFile[3].Time, expectedValue)
	}
	expectedValue = "1000"
	if loadedLogFile[3].Time != expectedValue {
		t.Errorf("Not the expected Time[3] value: %s (expecting %s)", loadedLogFile[3].Time, expectedValue)
	}
	expectedValue = "20m"
	if loadedLogFile[3].Band != expectedValue {
		t.Errorf("Not the expected Band[3] value: %s (expecting %s)", loadedLogFile[3].Band, expectedValue)
	}
	expectedValue = "ON4UP"
	if loadedLogFile[3].Call != expectedValue {
		t.Errorf("Not the expected Call[3] value: %s (expecting %s)", loadedLogFile[3].Call, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_redefining_myCall_must_fail(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "myCall on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "ON4KJM/P"
	if loadedLogFile[0].MyCall != expectedValue {
		t.Errorf("Not the expected MyCall value: %s (expecting %s)", loadedLogFile[0].MyCall, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_redefining_myWWFF_must_fail(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff onff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "myWWFF onff-0001")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "ONFF-0258"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_redefining_mySOTA_must_fail(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff onff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "mySota on/on-111")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "ON/ON-001"
	if loadedLogFile[0].MySOTA != expectedValue {
		t.Errorf("Not the expected MySOTA value: %s (expecting %s)", loadedLogFile[0].MySOTA, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_redefining_myGRID_must_fail(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff onff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, "myGrid YO50")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "myGrid ZZ99")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "YO50"
	if loadedLogFile[0].MyGrid != expectedValue {
		t.Errorf("Not the expected MyGRID value: %s (expecting %s)", loadedLogFile[0].MyGrid, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_redefining_operator_must_fail(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff onff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "operator blaaahh")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "ON4KJM"
	if loadedLogFile[0].Operator != expectedValue {
		t.Errorf("Not the expected operator value: %s (expecting %s)", loadedLogFile[0].Operator, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_redefining_nickname_must_fail(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "myWwff onff-0258")
	dataArray = append(dataArray, "mySota on/on-001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "20/5/23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "nickname blaaahh")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "Portable"
	if loadedLogFile[0].Nickname != expectedValue {
		t.Errorf("Not the expected Nickname value: %s (expecting %s)", loadedLogFile[0].Nickname, expectedValue)
	}

	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_bad_date(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "{ Sample multi-line comment")
	dataArray = append(dataArray, "	( with quotes) Check: Logging > \"Contest Logging\"")
	dataArray = append(dataArray, "  - Data item1")
	dataArray = append(dataArray, "  - Data item2")
	dataArray = append(dataArray, "  }")
	dataArray = append(dataArray, "{ one liner comment }")
	dataArray = append(dataArray, "  { offset one liner comment }")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "nickname Portable")
	dataArray = append(dataArray, "date 2020-5-18")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "20/5/233")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "0954 on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
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
	expectedValue = "IK5ZVE/5"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-18"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_wrongHeader(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall fooBar")
	dataArray = append(dataArray, "operator")
	dataArray = append(dataArray, "myWwff  foobar")
	dataArray = append(dataArray, "mySota  ")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve/5 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "0954 on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "*FOOBAR"
	if loadedLogFile[0].MyCall != expectedValue {
		t.Errorf("Not the expected MyCall value: %s (expecting %s)", loadedLogFile[0].MyCall, expectedValue)
	}
	expectedValue = ""
	if loadedLogFile[0].Operator != expectedValue {
		t.Errorf("Not the expected Operator value: %s (expecting %s)", loadedLogFile[0].Operator, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}
	expectedValue = ""
	if loadedLogFile[0].MySOTA != expectedValue {
		t.Errorf("Not the expected MySOTA value: %s (expecting %s)", loadedLogFile[0].MySOTA, expectedValue)
	}

	expectedValue = "IK5ZVE/5"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

// if the first call is wrong the infertime doesn't work
func TestLoadFile_InferTime_missingStartTime(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "myWwff  onff-0001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw ik5zve 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "40m 0954 on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
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
	expectedValue = "ONFF-0001"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}

	expectedValue = "IK5ZVE"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = ""
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = ""
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_InferTime_missingEndTime(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "myWwff  onff-0001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "40m on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
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
	expectedValue = "ONFF-0001"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}

	expectedValue = "IK5ZVE"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

// FIXME: same time
func TestLoadFile_2_QSO_same_time(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall on4kjm/p")
	dataArray = append(dataArray, "operator on4kjm")
	dataArray = append(dataArray, "myWwff  onff-0001")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve 9 5")
	dataArray = append(dataArray, "0951 on6zq")
	dataArray = append(dataArray, "f6AA")
	dataArray = append(dataArray, "0951 on4do")
	dataArray = append(dataArray, "0952 on4bb")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if !isLoadedOK {
		t.Error("Test file should not return with an error")
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
	expectedValue = "ONFF-0001"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}

	expectedValue = "IK5ZVE"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0951"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "F6AA"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0951"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[3].Call != expectedValue {
		t.Errorf("Not the expected Call[3] value: %s (expecting %s)", loadedLogFile[3].Call, expectedValue)
	}
	expectedValue = "0951"
	if loadedLogFile[3].Time != expectedValue {
		t.Errorf("Not the expected Time[3] value: %s (expecting %s)", loadedLogFile[3].Time, expectedValue)
	}
	expectedValue = "ON4BB"
	if loadedLogFile[4].Call != expectedValue {
		t.Errorf("Not the expected Call[4] value: %s (expecting %s)", loadedLogFile[4].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[4].Time != expectedValue {
		t.Errorf("Not the expected Time[4] value: %s (expecting %s)", loadedLogFile[4].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_wrongData(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall fooBar")
	dataArray = append(dataArray, "operator foobar")
	dataArray = append(dataArray, "myWwff  foobar")
	dataArray = append(dataArray, "mySota foobar")
	dataArray = append(dataArray, "myGrid foobar")
	dataArray = append(dataArray, "myLat foobar")
	dataArray = append(dataArray, "myLon foobar")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 ik5zve 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "42m 0954 on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "*FOOBAR"
	if loadedLogFile[0].MyCall != expectedValue {
		t.Errorf("Not the expected MyCall value: %s (expecting %s)", loadedLogFile[0].MyCall, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].Operator != expectedValue {
		t.Errorf("Not the expected Operator value: %s (expecting %s)", loadedLogFile[0].Operator, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].MySOTA != expectedValue {
		t.Errorf("Not the expected MySOTA value: %s (expecting %s)", loadedLogFile[0].MySOTA, expectedValue)
	}
	expectedValue = "*foobar"
	if loadedLogFile[0].MyGrid != expectedValue {
		t.Errorf("Not the expected MyGrid value: %s (expecting %s)", loadedLogFile[0].MyGrid, expectedValue)
	}
	if loadedLogFile[0].MyLat != expectedValue {
		t.Errorf("Not the expected MyLat value: %s (expecting %s)", loadedLogFile[0].MyLat, expectedValue)
	}
	if loadedLogFile[0].MyLon != expectedValue {
		t.Errorf("Not the expected MyLon value: %s (expecting %s)", loadedLogFile[0].MyLon, expectedValue)
	}

	expectedValue = "IK5ZVE"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

func TestLoadFile_firstCallWrong(t *testing.T) {

	//Given
	dataArray := make([]string, 0)
	dataArray = append(dataArray, "# Header")
	dataArray = append(dataArray, "myCall fooBar")
	dataArray = append(dataArray, "operator foobar")
	dataArray = append(dataArray, "myWwff  foobar")
	dataArray = append(dataArray, "mySota foobar")
	dataArray = append(dataArray, " ")
	dataArray = append(dataArray, " #Log")
	dataArray = append(dataArray, "date 2020-05-23")
	dataArray = append(dataArray, "40m cw 0950 on4kjm/p/qrp 9 5")
	dataArray = append(dataArray, "on6zq")
	dataArray = append(dataArray, "42m 0954 on4do")

	temporaryDataFileName := createTestFile(dataArray)

	//When
	loadedLogFile, isLoadedOK := LoadFile(temporaryDataFileName, true)

	//Then
	if isLoadedOK {
		t.Error("Test file processing should return with an error")
	}
	if len(loadedLogFile) == 0 {
		t.Error("No data loaded")
	}

	expectedValue := "*FOOBAR"
	if loadedLogFile[0].MyCall != expectedValue {
		t.Errorf("Not the expected MyCall value: %s (expecting %s)", loadedLogFile[0].MyCall, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].Operator != expectedValue {
		t.Errorf("Not the expected Operator value: %s (expecting %s)", loadedLogFile[0].Operator, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].MyWWFF != expectedValue {
		t.Errorf("Not the expected MyWWFF value: %s (expecting %s)", loadedLogFile[0].MyWWFF, expectedValue)
	}
	expectedValue = "*FOOBAR"
	if loadedLogFile[0].MySOTA != expectedValue {
		t.Errorf("Not the expected MySOTA value: %s (expecting %s)", loadedLogFile[0].MySOTA, expectedValue)
	}

	expectedValue = "*ON4KJM/P/QRP"
	if loadedLogFile[0].Call != expectedValue {
		t.Errorf("Not the expected Call[0] value: %s (expecting %s)", loadedLogFile[0].Call, expectedValue)
	}
	expectedValue = "0950"
	if loadedLogFile[0].Time != expectedValue {
		t.Errorf("Not the expected Time[0] value: %s (expecting %s)", loadedLogFile[0].Time, expectedValue)
	}
	expectedValue = "2020-05-23"
	if loadedLogFile[0].Date != expectedValue {
		t.Errorf("Not the expected Date[0] value: %s (expecting %s)", loadedLogFile[0].Date, expectedValue)
	}
	expectedValue = "ON6ZQ"
	if loadedLogFile[1].Call != expectedValue {
		t.Errorf("Not the expected Call[1] value: %s (expecting %s)", loadedLogFile[1].Call, expectedValue)
	}
	expectedValue = "0952"
	if loadedLogFile[1].Time != expectedValue {
		t.Errorf("Not the expected Time[1] value: %s (expecting %s)", loadedLogFile[1].Time, expectedValue)
	}
	expectedValue = "ON4DO"
	if loadedLogFile[2].Call != expectedValue {
		t.Errorf("Not the expected Call[2] value: %s (expecting %s)", loadedLogFile[2].Call, expectedValue)
	}
	expectedValue = "0954"
	if loadedLogFile[2].Time != expectedValue {
		t.Errorf("Not the expected Time[2] value: %s (expecting %s)", loadedLogFile[2].Time, expectedValue)
	}
	//Clean Up
	os.Remove(temporaryDataFileName)
}

// createTestFile creates and populates a test FLE input file.
// Returns the created temporary filename.
func createTestFile(dataArray []string) (tempFileName string) {
	//create random file name
	tmpfile, err := os.CreateTemp("", "*.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Temporary file: %s", tmpfile.Name())

	//Write the passed data to the file
	writeFile(tmpfile.Name(), dataArray)

	//Return the temporary filename
	return tmpfile.Name()
}
