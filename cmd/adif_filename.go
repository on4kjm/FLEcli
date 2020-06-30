package cmd

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
	//"log"
	"fmt"
	"path/filepath"
)

// does the target file exist?
// is the file defined
//	remove the extention

//returning "" is considered as invalid
func buildOutputFilename(output string, input string, overwrite bool) (outputFilename string, wasOK bool) {
	outputFilename = ""

	//validate that input is populated (should never happen if properly called)
	if input == "" {
		return "", false
	}

	//No output was provided, let's create one from the input file
	if output == "" {
		extension := filepath.Ext(input)
		outputRootPart := input[0 : len(input)-len(extension)]
		output = outputRootPart + ".adi"
		fmt.Println("No output provided, defaulting to \"" + output + "\"")
	}

	//an output was provided by the user
	if output != "" {
		info, err := os.Stat(output)
		if os.IsNotExist(err) {
			return output, true
		}
		//It exisits but is a directory
		if info.IsDir() {
			fmt.Println("Error: specified output exists and is a directory")
			return "", false
		}
		if overwrite {
			//user accepted to overwrite the file
			return output, true
		} else {
			fmt.Println("File already exists. Use --overwrite flag if necessary.")
			return "", false
		}
	}

	return outputFilename, true
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
