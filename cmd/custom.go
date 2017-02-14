// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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
	"log"

	"github.com/sascha-andres/go-snippet/app"
	"github.com/sascha-andres/go-snippet/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// customCmd represents the custom command
var customCmd = &cobra.Command{
	Use:   "custom",
	Short: "Place snippets using a custom marker",
	Long: `Use this command to use custom marker lines.

A snippetmarker must be a regex with two groups, the latter one
matching the snippet filename relative to the snippetdir`,
	Run: func(cmd *cobra.Command, args []string) {
		if "" == viper.GetString("custom.marker") {
			log.Fatal("No custom marker provided")
		}
		if err := app.Processor(data.SnippetMarker(viper.GetString("custom.marker"))); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(customCmd)
	customCmd.Flags().String("marker", "", "Custom marker")
	viper.BindPFlag("custom.marker", customCmd.Flags().Lookup("marker"))

}
