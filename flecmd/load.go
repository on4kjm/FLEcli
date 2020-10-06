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

var processLoadFile = fleprocess.LoadFile
var loadCmd = loadCmdConstructor()

// loadCmd represents the load command
func loadCmdConstructor() *cobra.Command {
	return &cobra.Command{
		Use:   "load [flags] inputFile",
		Short: "Loads and validates a FLE type shorthand logfile",

		RunE: func(cmd *cobra.Command, args []string) error {
			//if args is empty, throw an error
			if len(args) < 1 {
				//FIXME: Doesn't work as expected
				return fmt.Errorf("Missing input file %s", "")
			}
			if len(args) > 1 {
				return fmt.Errorf("Too many arguments.%s", "")
			}
			inputFilename = args[0]
			//FIXME: we should return the result of the call
			processLoadFile(inputFilename, isInterpolateTime)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(loadCmd)

	loadCmd.PersistentFlags().BoolVarP(&isInterpolateTime, "interpolate", "i", false, "Interpolates the missing time entries.")
}
