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
	"path/filepath"
)

//buildOutputFilname will try to figure out an output filename (for the case none was provided)
func buildOutputFilename(output string, input string, overwrite bool, newExtension string) (string, error) {

	//validate that input is populated (should never happen if properly called)
	if input == "" {
		return "", fmt.Errorf("Unexepected error: no input file provided")
	}

	//No output was provided, let's create one from the input file
	if output == "" {
		extension := filepath.Ext(input)
		outputRootPart := input[0 : len(input)-len(extension)]
		output = outputRootPart + newExtension
		fmt.Println("No output provided, defaulting to \"" + output + "\"")
	}

	//process the computed or user-provided output filename
	info, err := os.Stat(output)
	if os.IsNotExist(err) {
		//File doesn't exist, so we're good
		return output, nil
	}
	//It exisits but is a directory
	if info.IsDir() {
		return "", fmt.Errorf("Error: specified output exists and is a directory")
	}
	if overwrite {
		//user accepted to overwrite the file
		return output, nil
	}

	return "", fmt.Errorf("File already exists. Use --overwrite flag if necessary")
}
