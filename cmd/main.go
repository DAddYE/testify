package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

const assertionsForward = `// auto genrated file, do not edit
package assert

import "time"

type Assertions struct {
	t TestingT
}

func New(t TestingT) *Assertions {
	return &Assertions{t: t}
}
{{range .}}
{{.Doc}}
func (a *Assertions) {{.Name}}({{.Params}}) {{.Results}} {
	return {{.Name}}(a.t, {{.Values}})
}
{{end}}`

const requirements = `// auto generated file, do not edit
package require

import (
	"time"

	"github.com/stretchr/testify/assert"
)

type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}

// Fail reports a failure through
func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{}) {
	assert.Fail(t, failureMessage, msgAndArgs...)
	t.FailNow()
}
{{range .}}
{{.DocRequire}}
func {{.Name}}(t TestingT, {{.ParamsRequire}}) {
	if !assert.{{.Name}}(t, {{.Values}}) {
		t.FailNow()
	}
}
{{end}}`

const requirementsForward = `// auto generated file, do not edit
package require

import (
	"time"

	"github.com/stretchr/testify/assert"
)

type Requirements struct {
	t TestingT
}

func New(t TestingT) *Requirements {
	return &Requirements{t: t}
}
{{range .}}
{{.DocRequireFd}}
func (r *Requirements) {{.Name}}({{.ParamsRequire}}) {
	{{.Name}}(r.t, {{.Values}})
}
{{end}}`

// used to prefix assert types in the require package
var assertTypes = []string{
	"Comparison",
	"PanicTestFunc",
}

type Node struct {
	Doc           string
	DocRequire    string
	DocRequireFd  string
	Name          string
	Params        string
	ParamsRequire string
	Values        string
	Results       string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen: ")

	fset := token.NewFileSet()
	parser, err := parser.ParseFile(fset, "./assert/assertions.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("parsing error: %s", err)
	}

	var nodes []Node

	for _, decl := range parser.Decls {

		fn, ok := decl.(*ast.FuncDecl)
		if !ok || !ast.IsExported(fn.Name.Name) || fn.Type.Params.NumFields() == 0 {
			continue
		}

		if _, ok := fn.Type.Params.List[0].Type.(*ast.Ident); !ok {
			continue
		}

		node := Node{Name: fn.Name.Name}

		// add back the comment
		var docsAssert, docsRequire, docsRequireFd []string

		for _, doc := range fn.Doc.List {
			noT := strings.Replace(doc.Text, "(t, ", "(", -1)
			noA := strings.Replace(doc.Text, "assert.", "require.", -1)
			noR := strings.Replace(noT, "assert.", "require.", -1)

			docsAssert = append(docsAssert, noT)
			docsRequire = append(docsRequire, noA)
			docsRequireFd = append(docsRequireFd, noR)
		}

		node.Doc = strings.Join(docsAssert, "\n")
		node.DocRequire = strings.Join(docsRequire, "\n")
		node.DocRequireFd = strings.Join(docsRequireFd, "\n")

		// parse the parameters
		var params, paramsRequire, values []string

		for _, param := range fn.Type.Params.List {
			t := getType(param.Type)
			tr := t

			if t == "TestingT" {
				continue
			}

			// prefix assert types for the require package
			for _, match := range assertTypes {
				if match == t {
					tr = "assert." + t
				}
			}

			names := make([]string, len(param.Names))
			for i, name := range param.Names {
				names[i] = name.Name
			}

			params = append(params, strings.Join(names, ", ")+" "+t)
			paramsRequire = append(paramsRequire, strings.Join(names, ", ")+" "+tr)
			values = append(values, strings.Join(names, ", "))
		}

		node.Params = strings.Join(params, ", ")
		node.ParamsRequire = strings.Join(paramsRequire, ", ")
		node.Values = strings.Join(values, ", ")

		// parse the results
		var results []string

		for _, result := range fn.Type.Results.List {
			t, ok := result.Type.(*ast.Ident)
			if !ok {
				continue
			}

			names := make([]string, len(result.Names))
			for i, name := range result.Names {
				names[i] = name.Name
			}

			results = append(results, strings.TrimSpace(strings.Join(names, ", ")+" "+t.Name))
		}

		switch l := len(results); {
		case l == 0:
		case l == 1:
			node.Results = results[0]
		default:
			node.Results = "(" + strings.Join(results, ", ") + ")"
		}

		nodes = append(nodes, node)
	}

	// save templates
	executeTemplate(assertionsForward, "./assert/assertions_forward.go", nodes)
	executeTemplate(requirements, "./require/requirements.go", nodes)
	executeTemplate(requirementsForward, "./require/requirements_forward.go", nodes)
}

func executeTemplate(tpl, filename string, nodes []Node) {
	out, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("open %s: %s", filename, err)
	}

	t, err := template.New("gen").Parse(tpl)
	if err != nil {
		log.Fatalf("template error for %s: %s", filename, err)
	}

	t.Execute(out, nodes)

	if err := out.Close(); err != nil {
		log.Fatalf("error closing %s: %s", filename, err)
	}
}

func getType(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.Ellipsis:
		return "..." + getType(x.Elt)
	case *ast.ArrayType:
		return "[]" + getType(x.Elt)
	case *ast.StructType:
		return "struct{}"
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.SelectorExpr:
		return getType(x.X) + "." + getType(x.Sel)
	default:
		log.Fatalf("unknow type: %#+v", x)
		return "<unknow T>"
	}
}
