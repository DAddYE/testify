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

type Node struct {
	Doc     string
	Name    string
	Params  string
	Values  string
	Results string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen: ")

	fset := token.NewFileSet()
	parser, err := parser.ParseFile(fset, "./assert/assertions.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("parsing error: %s", err)
	}

	out, err := os.OpenFile("./assert/assertions_forward.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("open file: %s", err)
	}

	tpl := `// autogenrated file, do not edit
package assert

import "time"

type Assertions struct {
	t TestingT
}

func New(t TestingT) *Assertions {
	return &Assertions{ t: t}
}
{{range .}}
{{.Doc}}
func (a *Assertions) {{.Name}}({{.Params}}) {{.Results}} {
	return {{.Name}}(a.t, {{.Values}})
}
{{end}}`

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
		docs := make([]string, len(fn.Doc.List))
		for i, doc := range fn.Doc.List {
			docs[i] = strings.Replace(doc.Text, "(t, ", "(", -1)
		}
		node.Doc = strings.Join(docs, "\n")

		// parse the parameters
		var params, values []string

		for _, param := range fn.Type.Params.List {
			t := getType(param.Type)
			if t == "TestingT" {
				continue
			}

			names := make([]string, len(param.Names))
			for i, name := range param.Names {
				names[i] = name.Name
			}

			params = append(params, strings.Join(names, ", ")+" "+t)
			values = append(values, strings.Join(names, ", "))
		}

		node.Params = strings.Join(params, ", ")
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

	t, err := template.New("gen").Parse(tpl)
	if err != nil {
		log.Fatalf("template parsing error: %s", err)
	}

	t.Execute(out, nodes)

	if err := out.Close(); err != nil {
		log.Fatalf("error closing file: %s", err)
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
