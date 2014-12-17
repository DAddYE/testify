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

const assertions = `// auto genrated file, do not edit
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

// used to prefix assert types in the require package
var assertTypes = []string{
	"Comparison",
	"PanicTestFunc",
}

type Node struct {
	Doc           string
	DocRequire    string
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
		var docsAssert, docsRequire []string

		for _, doc := range fn.Doc.List {
			docsAssert = append(docsAssert, strings.Replace(doc.Text, "(t, ", "(", -1))
			docsRequire = append(docsRequire, strings.Replace(doc.Text, "assert.", "require.", -1))
		}

		node.Doc = strings.Join(docsAssert, "\n")
		node.DocRequire = strings.Join(docsRequire, "\n")

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

	{ // assertions_forward template
		out, err := os.OpenFile("./assert/assertions_forward.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("open assertions_forward.go: %s", err)
		}

		t, err := template.New("gen").Parse(assertions)
		if err != nil {
			log.Fatalf("assertions template error: %s", err)
		}

		t.Execute(out, nodes)

		if err := out.Close(); err != nil {
			log.Fatalf("error closing assertions_forward.go: %s", err)
		}
	}

	{ // requirements template
		out, err := os.OpenFile("./require/requirements.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("open requirements.go: %s", err)
		}

		t, err := template.New("gen").Parse(requirements)
		if err != nil {
			log.Fatalf("requirements template error: %s", err)
		}

		t.Execute(out, nodes)

		if err := out.Close(); err != nil {
			log.Fatalf("error closing requirements.go: %s", err)
		}
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
