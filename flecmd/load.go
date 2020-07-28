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

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads and validates a FLE type shorthand logfile",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("load called")
		//fmt.Println("Inputfile: ",inputFilename)
		fleprocess.LoadFile(inputFilename, isInterpolateTime)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)

	loadCmd.PersistentFlags().StringVarP(&inputFilename, "input", "i", "", "FLE formatted input file (mandatory)")
	loadCmd.MarkPersistentFlagRequired("input")

	loadCmd.PersistentFlags().BoolVarP(&isInterpolateTime, "interpolate", "", false, "Interpolates the missing time entries.")

}


