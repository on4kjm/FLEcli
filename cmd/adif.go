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
	//	"log"
	//"strings"
)

var outputFilename string
var isWWFFcli bool
var isSOTAcli bool
var isOverwrite bool

// adifCmd is executed when choosing the adif option (load and generate adif file)
var adifCmd = &cobra.Command{
	Use:   "adif",
	Short: "Generates an ADIF file based on a FLE type shorthand logfile.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	Run: func(cmd *cobra.Command, args []string) {
		processAdifCommand()
	},
}

func init() {
	rootCmd.AddCommand(adifCmd)

	adifCmd.PersistentFlags().BoolVarP(&isWWFFcli, "wwff", "w", false, "Generates a WWFF ready ADIF file.")
	adifCmd.PersistentFlags().BoolVarP(&isSOTAcli, "sota", "s", false, "Generates a SOTA ready ADIF file.")
	adifCmd.PersistentFlags().BoolVarP(&isOverwrite, "overwrite", "", false, "Overwrites the output file if it exisits")
	adifCmd.PersistentFlags().StringVarP(&outputFilename, "output", "o", "", "Output filename")
}

func processAdifCommand() {

	verifiedOutputFilename, filenameWasOK := buildOutputFilename(outputFilename, inputFilename, isOverwrite, ".adi")

	// if the output file could not be parsed correctly do noting
	if filenameWasOK {
		loadedLogFile, isLoadedOK := loadFile()

		//TODO: move this in a function so that it can be more easily tested
		if isLoadedOK {
			if len(loadedLogFile) == 0 {
				fmt.Println("No useful data read. Aborting...")
				return
			}

			//TODO: There are more tests required here
			//check if we have the necessary information for the type
			if isWWFFcli {
				if loadedLogFile[0].MyWWFF == "" {
					fmt.Println("Missing MY-WWFF reference. Aborting...")
					return
				}
				if loadedLogFile[0].Operator == "" {
					fmt.Println("Missing Operator. Aborting...")
					return
				}
			}
			if isSOTAcli {
				if loadedLogFile[0].MySOTA == "" {
					fmt.Println("Missing MY-SOTA reference. Aborting...")
					return
				}				
			}

			outputAdif(verifiedOutputFilename, loadedLogFile, isWWFFcli, isSOTAcli)
		}
	}
}
