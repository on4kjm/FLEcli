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
	"fmt"
	"github.com/spf13/cobra"
	"bufio"
	"log"
	"os"
	"regexp"
	//"strings"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads and validates a FLE type logfile",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("load called")
		//fmt.Println("Inputfile: ",inputFilename)
		loadFile()
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loadFile() {
	file, err := os.Open(inputFilename)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()

	regexpLineComment, _ := regexp.Compile("^#")
	regexpOnlySpaces, _ := regexp.Compile("^\\s+$")
	regexpSingleMultiLineComment, _ := regexp.Compile("^{.+}$")
	regexpStartMultiLineComment, _ := regexp.Compile("^{")
	regexpEndMultiLineComment, _ := regexp.Compile("}$")
	regexpHeaderMyCall, _ := regexp.Compile("(?i)^mycall ")
	regexpHeaderOperator, _ := regexp.Compile("(?i)^operator ") 
	// regexpHeaderMyWwff, _ := regexp.Compile("(?i)^mywwff ")
	// regexpHeaderMySota, _ := regexp.Compile("(?i)^mysota ")
	// regexpHeaderQslMsg, _ := regexp.Compile("(?i)^qslmsg ")
	// regexpHeaderNickname, _ := regexp.Compile("(?i)^nickname ")
	// regexpHeaderDate, _ := regexp.Compile("(?i)^date ")

	headerMyCall := ""
	headerOperator := ""
	lineCount := 0

	var isInMultiLine = false 
	var cleanedInput []string
	var errorLog []string

	
 
	//Loop through all the stored lined
	for _, eachline := range txtlines {
		lineCount++

		// ****
		// ** Lets do some house keeping first by droping the unecessary lines
		// ****

		//Skip the line if it starts with "#"
		if(regexpLineComment.MatchString(eachline)) {
			continue
		}
		//Skip if line is empty or blank
		if((len(eachline) == 0) || (regexpOnlySpaces.MatchString(eachline))) {
			continue
		}

		// Process multi-line comments
		if(regexpStartMultiLineComment.MatchString(eachline)) {
			//Single-line "multi-line" coment
			if(regexpSingleMultiLineComment.MatchString(eachline)) {
			 	continue
			}
			isInMultiLine = true
			continue
		}
		if(isInMultiLine) {
			if(regexpEndMultiLineComment.MatchString(eachline)) {
				isInMultiLine = false
			}
			continue
		}

		// ****
		// ** Process the Header block
		// ****

		if(regexpHeaderMyCall.MatchString(eachline)) {
			errorMsg := ""
			myCallList := regexpHeaderMyCall.Split(eachline,-1)			
			if(len(myCallList[1]) > 0) {
				headerMyCall, errorMsg = ValidateCall(myCallList[1])
				cleanedInput = append(cleanedInput, fmt.Sprintf("My call: %s", headerMyCall))
				if(len(errorMsg) != 0) {
					errorLog = append(errorLog, fmt.Sprintf("Invalid myCall at line %d: %s (%s)",lineCount, myCallList[1], errorMsg))
				}
			} else {
				errorLog = append(errorLog, fmt.Sprintf("Undefined myCall at line %d",lineCount))
			}
			continue
		}

		if(regexpHeaderOperator.MatchString(eachline)) {
			errorMsg := ""
			myOperatorList := regexpHeaderOperator.Split(eachline,-1)
			if(len(myOperatorList[1]) > 0) {
				headerOperator, errorMsg = ValidateCall(myOperatorList[1])
				cleanedInput = append(cleanedInput, fmt.Sprintf("Operator: %s", headerOperator))
				if(len(errorMsg) != 0) {
					errorLog = append(errorLog, fmt.Sprintf("Invalid Operator at line %d: %s (%s)",lineCount, myOperatorList[1], errorMsg))	
				}
			} else {
				errorLog = append(errorLog, fmt.Sprintf("Undefined Operator at line %d",lineCount))		
			}
			continue
		}
		// ****
		// ** Process the data block
		// ****
		cleanedInput = append(cleanedInput,eachline)

	}

	for _, cleanedInputLine := range cleanedInput {
		fmt.Println(cleanedInputLine)
	}

	if(len(errorLog) != 0){
		fmt.Println("\nParsing errors:")
		for _, errorLogLine := range errorLog {
			fmt.Println(errorLogLine)
		}
	} else {
		fmt.Println("\nSuccesfuly parsed ",lineCount, " lines.")
	}


}
