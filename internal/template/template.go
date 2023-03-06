package template

import (
	"errors"
	"os"
	"text/template"
)

var name = "embed-evaluate"
var ErrTemplateExecuteFailed = errors.New("executing template is failed")

var embedTmpl = `// Code generated by generate-go-embed4assets. DO NOT EDIT.
// versions:
//   generate-go-embed4assets {{.ToolVersion}}
// source {{.File}}

package {{.GoPackageName}}

import (
	_ "embed"
)

//go:embed {{.Filename}}
var {{.GoVariableName}} []byte

`

type Data struct {
	ToolVersion    string
	File           string
	Filename       string
	GoPackageName  string
	GoVariableName string
}

func Process(filePath string, data Data) error {
	tmpl := template.Must(template.New(name).Parse(embedTmpl))

	file, _ := os.Create(filePath)
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return errors.Join(err, ErrTemplateExecuteFailed)
	}

	return nil
}
