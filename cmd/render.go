// Copyright © 2017 SEAN ROSS-ROSS srossross@gmail.com
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
	"github.com/spf13/cobra"
	"github.com/ghodss/yaml"

	"github.com/srossross/template/lib"
)

var varValues []string
var varValueFiles []string
var verbose bool
var outputArg string
var sep string

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render a template or set of templates",
	Long: `Render a template or set of templates

Exanple:

  template render ./template.tpl

Environment variables can be accessed inside of a template with {{ .Env.VALUE }}

Values files: given on the command line with '-f' or '--values'.

Are plain YAML files. They can be used inside of a template as {{ .Values.VALUE }}

Each new values file specified gets merged into the '.Values' object

`,
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
		if verbose {
			fmt.Fprintln(os.Stderr, ":Values:\n", string(ValuesYAML))
		}

		ctx := lib.Context{  }
		yaml.Unmarshal(ValuesYAML, &ctx.Values)
		ctx.Env = lib.UnmarshalEnv()

		for fileSpec := range lib.InputIterator(args, false) {

			output, err := lib.Render(fileSpec.In, ctx)

			if err != nil {
				fmt.Fprintln(os.Stderr, err);
				os.Exit(1)
			}

			if strings.TrimSpace(fileSpec.Out) == "-" {
				_, err = os.Stdout.Write([]byte(output))
			} else {
				err = ioutil.WriteFile(fileSpec.Out, []byte(output), 0644)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, err);
				os.Exit(1)
			}
			if len(sep) > 0 {
				fmt.Fprintln(os.Stdout, sep);
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.
	renderCmd.Flags().StringArrayVarP(&varValueFiles, "values", "f", []string{}, "specify values in a YAML file (can specify multiple)")
	renderCmd.Flags().StringArrayVar(&varValues, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	renderCmd.Flags().StringVarP(&outputArg, "output", "o", "-", "Print template output to a file")
	renderCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose output to stderr")
	renderCmd.Flags().StringVar(&sep, "sep", "", "When printing to stdout, seperate files with sep")
}
