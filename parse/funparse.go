package main

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

// FuncParse test
func FuncParse(ctx context.Context) string {

	return "test"
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "funparse.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		fd.Doc.List = append(fd.Doc.List, &ast.Comment{
			Slash: fd.End(),
			Text:  "// testss",
		})
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}

	fmt.Println(string(buf.Bytes()))
}
