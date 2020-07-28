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
	"github.com/spf13/cobra"
)

var outputCsvFilename string
var isOverwriteCsv bool

// csvCmd is executed when choosing the csv option (load FLE file and generate csv file)
var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Generates a SOTA .csv file based on a FLE type shorthand logfile.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	Run: func(cmd *cobra.Command, args []string) {
		fleprocess.ProcessCsvCommand(inputFilename, outputCsvFilename, isInterpolateTime, isOverwriteCsv)
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)

	csvCmd.PersistentFlags().StringVarP(&inputFilename, "input", "i", "", "FLE formatted input file (mandatory)")
	csvCmd.MarkPersistentFlagRequired("input")
	csvCmd.PersistentFlags().BoolVarP(&isInterpolateTime, "interpolate", "", false, "Interpolates the missing time entries.")

	csvCmd.PersistentFlags().BoolVarP(&isOverwriteCsv, "overwrite", "", false, "Overwrites the output file if it exisits")
	csvCmd.PersistentFlags().StringVarP(&outputCsvFilename, "output", "o", "", "Output filename")
}

