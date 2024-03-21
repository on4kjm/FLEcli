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
	"os"

	"github.com/spf13/cobra"
)

var outputCsvFilename string
var isOverwriteCsv bool

var csvCmd = csvCmdConstructor()

// csvCmd is executed when choosing the csv option (load FLE file and generate csv file)
func csvCmdConstructor() *cobra.Command {
	return &cobra.Command{
		Use:   "csv [flags] inputFile [outputFile]",
		Short: "Generates a SOTA .csv file based on a FLE type shorthand logfile.",
		// 	Long: `A longer description that spans multiple lines and likely contains examples
		// and usage of using your command. For example:

		RunE: func(cmd *cobra.Command, args []string) error {
			//if args is empty, throw an error (Cobra will display the )
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

			// Verify given output directory exists. This check should be performed
			// Before running any long process so as to not make the user wait and
			// then be notified the file cannot be written.
			dirErr := CheckDir(outputCsvFilename)
			if dirErr != nil {
				return dirErr
			}

			if err := fleprocess.ProcessCsvCommand(inputFilename, outputCsvFilename, isInterpolateTime, isOverwriteCsv); err != nil {
				fmt.Println("\nUnable to generate CSV file:")
				fmt.Println(err)
				os.Exit(1)
			}
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(csvCmd)

	csvCmd.PersistentFlags().BoolVarP(&isInterpolateTime, "interpolate", "i", false, "Interpolates the missing time entries.")

	csvCmd.PersistentFlags().BoolVarP(&isOverwriteCsv, "overwrite", "o", false, "Overwrites the output file if it exisits")
}
