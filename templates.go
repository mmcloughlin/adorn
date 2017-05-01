package adorn

import "text/template"

var packageTemplate = template.Must(template.New("package").Parse(
	"package {{ .Package }}\n\n",
))

var interfaceTemplate = template.Must(template.New("interface").Parse(`
type {{ .TypeName }} interface {
	{{ .MethodName }}({{ .ArgumentsUnnamed }}) {{ .ReturnType }}
}
`))

var funcTemplate = template.Must(template.New("func").Parse(`
// {{ .FuncTypeName }} is an adapter to allow ordinary functions to be used as {{ .TypeName }} implementations.
type {{ .FuncTypeName }} func({{ .ArgumentsUnnamed }}) {{ .ReturnType }}

// {{ .MethodName }} calls f.
func (f {{ .FuncTypeName }}) {{ .MethodName }}({{ .ArgumentsUnnamed }}) {{ .ReturnType }} {
	return f()
}
`))

var assertionTemplate = template.Must(template.New("assertion").Parse(`
// Compile time assertion that {{ .FuncTypeName }} implements {{ .TypeName }}.
var _ {{ .TypeName }} = new({{ .FuncTypeName }})
`))
