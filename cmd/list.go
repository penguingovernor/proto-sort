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
	"fmt"
	"os"

	"github.com/penguingovernor/proto-sort/internal/pencode"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		inputFileString, err := cmd.Flags().GetString("file")
		if err != nil {
			logrus.Fatalf("internal error - could not get flag: %q", "file")
		}

		inputFilefd, err := os.Open(inputFileString)
		if err != nil {
			logrus.Fatalf("could not open file %s: %v", inputFileString, err)
		}

		defer inputFilefd.Close()

		tnums, err := pencode.GetNumbers(inputFilefd)
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(tnums)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("file", "f", "numbers.pb", "the file to read from")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
