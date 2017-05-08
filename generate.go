package adorn

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"errors"
)

// Config encapsulates parameters for code generation.
type Config struct {
	Package       string
	TypeName      string
	MethodName    string
	ArgumentTypes []string
	ReturnTypes   []string
}

// FuncTypeName returns the name of the plain function type that implements the
// interface. This will be TypeName with "Func" appended.
func (c Config) FuncTypeName() string {
	return c.TypeName + "Func"
}

// ArgumentsUnnamed returns the list of argument types joined with commas.
func (c Config) ArgumentsUnnamed() string {
	return strings.Join(c.ArgumentTypes, ", ")
}

// ArgumentNames returns the names of the arguments.
func (c Config) ArgumentNames() []string {
	n := len(c.ArgumentTypes)
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = fmt.Sprintf("a%d", i)
	}
	return names
}

// ArgumentTypesDeduped returns the list of argument types, where runs of the same types
// are collapsed.
func (c Config) ArgumentTypesDeduped() []string {
	n := len(c.ArgumentTypes)
	types := make([]string, n)
	if n == 0 {
		return types
	}
	types[n-1] = c.ArgumentTypes[n-1]
	for i := 0; i < n-1; i++ {
		if c.ArgumentTypes[i] != c.ArgumentTypes[i+1] {
			types[i] = c.ArgumentTypes[i]
		}
	}
	return types
}

// ArgumentsNamed returns the arguments signature with names.
func (c Config) ArgumentsNamed() (string, error) {
	names := c.ArgumentNames()
	types := c.ArgumentTypesDeduped()
	if len(names) != len(types) {
		return "", errors.New("mismatch in number of argument names and argument types")
	}

	n := len(names)
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = names[i]
		if types[i] != "" {
			args[i] += " " + types[i]
		}
	}

	return strings.Join(args, ", "), nil
}

// ArgumentsCalling returns the comma separated list of argument names, used when
// calling the function or method
func (c Config) ArgumentsCalling() string {
	return strings.Join(c.ArgumentNames(), ", ")
}

// ReturnSignature returns the specification of the return type.
func (c Config) ReturnSignature() string {
	switch len(c.ReturnTypes) {
	case 0:
		return ""
	case 1:
		return c.ReturnTypes[0]
	default:
		return "(" + strings.Join(c.ReturnTypes, ", ") + ")"
	}
}

// Generate generates code for the given type Config and writes it to the given
// Writer.
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

// GenerateString returns code for the given type Config.
func GenerateString(c Config) (string, error) {
	var b bytes.Buffer
	err := Generate(c, &b)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
