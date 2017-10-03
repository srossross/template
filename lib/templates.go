package lib

import (
  "strings"
  "os"
  "io/ioutil"
  "text/template"
  "github.com/Masterminds/sprig"
  "bytes"
)

// Render a template
func Render(filePath string, ctx Context) (string, error) {

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
