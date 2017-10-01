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
	"bytes"
	"strings"
	"io/ioutil"
	"text/template"
	"github.com/spf13/cobra"
	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"

	"github.com/srossross/template/lib"
)

var varValues []string
var varValueFiles []string

func render(filePath string, ctx lib.Context) (string, error) {

	var input []byte
	var err error
	if strings.TrimSpace(filePath) == "-" {
		input, err = ioutil.ReadAll(os.Stdin)
	} else {
		input, err = ioutil.ReadFile(filePath)
	}
	if err != nil {
		return "", err
	}

	tmpl, err := template.New(filePath).Funcs(sprig.TxtFuncMap()).Parse(string(input))

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, ctx)

	return tpl.String(), err

}

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render a template or set of templates",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "render requires an input template argument");
			os.Exit(1)
		}

		ValuesYAML, err := lib.BuildValues(varValueFiles, varValues)

		if err != nil {
			fmt.Fprintln(os.Stderr, err);
			os.Exit(1)
		}

		fmt.Fprintln(os.Stderr, ":Values:\n", string(ValuesYAML));

		outputPath := "-"
		for _, fileSpec := range args {
			fileSpecList := strings.Split(fileSpec, ":")
			inputPath := fileSpecList[0]
			if len(fileSpecList) > 1 {
				outputPath = fileSpecList[1]
			}

			ctx := lib.Context{  }
			yaml.Unmarshal(ValuesYAML, &ctx.Values)
			ctx.Env = lib.UnmarshalEnv()

			output, err := render(inputPath, ctx)

			if err != nil {
				fmt.Fprintln(os.Stderr, err);
				os.Exit(1)
			}

			if strings.TrimSpace(outputPath) == "-" {
				_, err = os.Stdout.Write([]byte(output))
			} else {
				err = ioutil.WriteFile(outputPath, []byte(output), 0644)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, err);
				os.Exit(1)
			}


		}

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
	renderCmd.Flags().BoolP("verbose", "v", false, "Print verbose output to stderr")
}
