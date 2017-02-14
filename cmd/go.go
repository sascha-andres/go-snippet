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

	"strings"

	"github.com/sascha-andres/go-snippet/app"
	"github.com/sascha-andres/go-snippet/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Preconfigured call for golang source files",
	Long: `This command runs the application preconfigured for 
Golang source files. It has to end with .go`,
	Run: func(cmd *cobra.Command, args []string) {
		if ok := strings.HasSuffix(viper.GetString("file"), ".go"); !ok {
			log.Fatal("File does not end with .go")
		}
		if err := app.Processor(data.SnippetMarkerGolang); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(goCmd)
}
