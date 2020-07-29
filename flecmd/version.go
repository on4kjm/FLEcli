package flecmd

/*
Copyright © 2020 Jean-Marc Meessen, ON4KJM <on4kjm@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	detailed   = false
	version    = "private build"
	commit     = "none"
	date       = "unknown"
	builtBy    = ""
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "\"version\" will output the current build information",
		Long:  ``,
		Run: func(_ *cobra.Command, _ []string) {
			var response string
			if detailed {
				prettyPrintedDate := "Unknown"
				if date != "unknown" {
					buildDate, error := time.Parse(time.RFC3339, date)
					if error == nil {
						prettyPrintedDate = buildDate.Format("2006-01-02 15:04") + " (UTC)"
					} else {
						prettyPrintedDate = fmt.Sprint(error)
					}
				}
				response = fmt.Sprintf("FLEcli :\n- version:  %s\n- commit:   %s\n- date:     %s\n- built by: %s\n", version, commit, prettyPrintedDate, builtBy)
			} else {
				response = fmt.Sprintf("FLEcli version: %s\n", version)
			}

			fmt.Printf("%+v", response)
			return
		},
	}
)

func init() {
	versionCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Prints the detailed version information")
	rootCmd.AddCommand(versionCmd)
}
