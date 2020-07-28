package flecmd

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
	"FLEcli/fleprocess"
	"fmt"

	"github.com/spf13/cobra"
)

var outputCsvFilename string
var isOverwriteCsv bool

// csvCmd is executed when choosing the csv option (load FLE file and generate csv file)
var csvCmd = &cobra.Command{
	Use:   "csv [flags] inputFile [outputFile]",
	Short: "Generates a SOTA .csv file based on a FLE type shorthand logfile.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	RunE: func(cmd *cobra.Command, args []string) error {
		//if args is empty, throw an error
		if len(args) == 0 {
			//TODO: fix this ugly statement (because I am lazy)
			return fmt.Errorf("Missing input file %s", "")
		}
		inputFilename = args[0]
		if len(args) == 2 {
			outputCsvFilename = args[1]
		}
		if len(args) > 2 {
			return fmt.Errorf("Too many arguments.%s", "")
		}


		//TODO: should return an error
		fleprocess.ProcessCsvCommand(inputFilename, outputCsvFilename, isInterpolateTime, isOverwriteCsv)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)

	csvCmd.PersistentFlags().BoolVarP(&isInterpolateTime, "interpolate", "i", false, "Interpolates the missing time entries.")

	csvCmd.PersistentFlags().BoolVarP(&isOverwriteCsv, "overwrite", "o", false, "Overwrites the output file if it exisits")
}

