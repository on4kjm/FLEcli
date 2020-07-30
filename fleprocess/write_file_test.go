package fleprocess

import (
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

	//TODO: check that the file contains what we want

	// //detete test file
	os.Remove(writeFileTestFname)

}


