// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"strconv"

	"github.com/penguingovernor/proto-sort/internal/pencode"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		outputfileString, err := cmd.Flags().GetString("file")
		if err != nil {
			logrus.Fatalf("internal error - could not get flag %q: %v", "file", err)
		}

		ouputfd, err := os.OpenFile(outputfileString, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

		if err != nil {
			logrus.Fatal("could not open/create file %s: %v", outputfileString, err)
		}

		defer ouputfd.Close()

		// Convert the args to ints
		tints := []uint64{}
		for _, arg := range args {
			converted, err := strconv.ParseUint(arg, 10, 64)
			if err != nil {
				logrus.Warnf("skipping unknown symbol %s\n", arg)
				continue
			}
			tints = append(tints, converted)
		}

		// Add them to the file
		err = pencode.Append(ouputfd, tints)
		if err != nil {
			logrus.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("file", "f", "numbers.pb", "the file to save the numbers to")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
