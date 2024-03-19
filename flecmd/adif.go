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

var outputFilename string
var isWWFFcli bool
var isSOTAcli bool
var isPOTAcli bool
var isOverwrite bool

// adifCmd is executed when choosing the adif option (load and generate adif file)
var adifCmd = &cobra.Command{
	Use:   "adif [flags] inputFile [outputFile]",
	Short: "Generates an ADIF file based on a FLE type shorthand logfile.",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	RunE: func(cmd *cobra.Command, args []string) error {

		//if args is empty, throw an error
		if len(args) == 0 {
			//TODO: fix this ugly statement (because I am lazy)
			return fmt.Errorf("missing input file %s", "")
		}
		inputFilename = args[0]
		if len(args) == 2 {
			outputFilename = args[1]
		}
		if len(args) > 2 {
			return fmt.Errorf("Too many arguments.%s", "")
		}

		// Verify given output directory exists. This check should be performed
		// Before running any long process so as to not make the user wait and
		// then be notified the file cannot be written.
		CheckDir(outputFilename)

		var adifParam = new(fleprocess.AdifParams)
		adifParam.InputFilename = inputFilename
		adifParam.OutputFilename = outputFilename
		adifParam.IsInterpolateTime = isInterpolateTime
		adifParam.IsSOTA = isSOTAcli
		adifParam.IsPOTA = isPOTAcli
		adifParam.IsWWFF = isWWFFcli
		adifParam.IsOverwrite = isOverwrite

		err := fleprocess.ProcessAdifCommand(*adifParam)
		if err != nil {
			fmt.Println("\nUnable to generate ADIF file:")
			fmt.Println(err)
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(adifCmd)

	adifCmd.PersistentFlags().BoolVarP(&isInterpolateTime, "interpolate", "i", false, "Interpolates the missing time entries.")
	adifCmd.PersistentFlags().BoolVarP(&isWWFFcli, "wwff", "w", false, "Generates a WWFF ready ADIF file.")
	adifCmd.PersistentFlags().BoolVarP(&isSOTAcli, "sota", "s", false, "Generates a SOTA ready ADIF file.")
	adifCmd.PersistentFlags().BoolVarP(&isPOTAcli, "pota", "p", false, "Generates a POTA ready ADIF file.")
	adifCmd.PersistentFlags().BoolVarP(&isOverwrite, "overwrite", "o", false, "Overwrites the output file if it exisits")
}
