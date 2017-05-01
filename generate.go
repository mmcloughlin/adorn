package adorn

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"errors"
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

func (c Config) ArgumentNames() []string {
	n := len(c.ArgumentTypes)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = fmt.Sprintf("a%d", i)
	}
	return names
}

func (c Config) ArgumentsNamed() (string, error) {
	names := c.ArgumentNames()
	types := c.ArgumentTypes
	if len(names) != len(types) {
		return "", errors.New("mismatch in number of argument names and argument types")
	}

	n := len(names)
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = names[i] + " " + types[i]
	}

	return strings.Join(args, ", "), nil
}

func (c Config) ArgumentsCalling() string {
	return strings.Join(c.ArgumentNames(), ", ")
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
