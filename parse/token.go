package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
)

func main() {
	binaryExpr()
}

func binaryExpr() {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}

func par() {
	expr, _ := parser.ParseExpr(`9527`)
	ast.Print(nil, expr)
}

func basic() {
	lit9527 := &ast.BasicLit{
		Kind:  token.INT,
		Value: "9527",
	}

	ast.Print(nil, lit9527)
}

func scan() {
	src := []byte(`println("test")`)

	fset := token.NewFileSet()
	file := fset.AddFile("hello.go", fset.Base(), len(src))

	s := scanner.Scanner{}
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
