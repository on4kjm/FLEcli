/*
Copyright © 2024 Jean-Marc Meessen jean-marc@meessen-web.org

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package flecmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CsvWithoutParmMustFail(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"csv"})
	error := rootCmd.Execute()

	assert.Error(t, error, "Function call should have failed")

	// Error is expected
	expectedMsg := "Error: Missing input file "
	lines := strings.Split(actual.String(), "\n")
	assert.Equal(t, expectedMsg, lines[0], "Function did not fail for the expected cause")
}

func Test_CsvWithToManyParmMustFail(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"csv", "param1", "param2", "param3"})
	error := rootCmd.Execute()

	assert.Error(t, error, "Function call should have failed")

	// Error is expected
	expectedMsg := "Error: Too many arguments."
	lines := strings.Split(actual.String(), "\n")
	assert.Equal(t, expectedMsg, lines[0], "Function did not fail for the expected cause")
}

func Test_CsvBadOutpoutDirMustFail(t *testing.T) {
	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"csv", "../test/data/fle-1.txt", "badDirectory/outputfile.adi"})
	error := rootCmd.Execute()

	assert.Error(t, error, "Function call should have failed")

	// Error is expected
	expectedMsg := "Error: The directory of specified output file (badDirectory) does not exist."
	lines := strings.Split(actual.String(), "\n")
	assert.Equal(t, expectedMsg, lines[0], "Function did not fail for the expected cause")
}
