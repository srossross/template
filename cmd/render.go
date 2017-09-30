// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"strings"
	"io/ioutil"
	"template/lib"
	"text/template"
	"github.com/spf13/cobra"
	// "github.com/imdario/mergo"
	// "github.com/ghodss/yaml"
)

var varValues []string
var varValueFiles []string

func render(filePath string, values []byte) ([]byte, error) {

	var bytes []byte
	var err error
	if strings.TrimSpace(filePath) == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filePath)
	}
	if err != nil {
		return []byte{}, err
	}

	tmpl, err := template.New(filePath).Parse(string(bytes))

	if err != nil {
		return []byte{}, err
	}

	err = tmpl.Execute(os.Stdout, c)

	return []byte{}, err

}
// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("render called")

		Values, err := lib.BuildValues(varValueFiles, varValues)

		if err != nil {
			fmt.Println(err);
			os.Exit(1)
		}
		fmt.Println("args", args);
		fmt.Println("{{Values", string(Values));

		filePath := args[0]
		output, err := render(filePath, Values)
		if err != nil {
			fmt.Println(err);
			os.Exit(1)
		}

		fmt.Println(output)


	},
}

func init() {
	RootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.
	renderCmd.Flags().StringArrayVarP(&varValueFiles, "values", "f", []string{}, "specify values in a YAML file (can specify multiple)")
	renderCmd.Flags().StringArrayVar(&varValues, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
