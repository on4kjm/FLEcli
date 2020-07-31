package fleprocess

import (
	"bufio"
	"os"
	"testing"
)

const writeFileTestDir string = "test2_dir"
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


