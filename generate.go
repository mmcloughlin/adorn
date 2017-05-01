package adorn

import (
	"bytes"
	"io"
	"strings"
	"text/template"
)

type Config struct {
	Package       string
	TypeName      string
	MethodName    string
	ArgumentTypes []string
	ReturnType    string
}

func (c Config) FuncTypeName() string {
	return c.TypeName + "Func"
}

func (c Config) ArgumentsUnnamed() string {
	return strings.Join(c.ArgumentTypes, ", ")
}

func Generate(c Config, w io.Writer) error {
	templates := []*template.Template{
		packageTemplate,
		interfaceTemplate,
		funcTemplate,
		assertionTemplate,
	}
	for _, tmpl := range templates {
		err := tmpl.Execute(w, c)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateString(c Config) (string, error) {
	var b bytes.Buffer
	err := Generate(c, &b)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
