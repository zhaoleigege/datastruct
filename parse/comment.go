package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {
	CommentAst()
}

func CommentAst() {
	source := `package a
               // B comment
               type B struct {
            	// C comment
	           C string
               }`
	buffer := make([]byte, 1024, 1024)
	for idx := range buffer {
		buffer[idx] = 0x20
	}
	copy(buffer[:], source)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", buffer, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	v := &visitor{
		file: file,
		fset: fset,
	}
	ast.Walk(v, file)

	//ast.Print(fset, file)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		panic(err)
	}

	println(len(buf.Bytes()))

	fmt.Println(string(buf.Bytes()))
}

type visitor struct {
	file *ast.File
	fset *token.FileSet
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return v
	}

	switch n := node.(type) {
	case *ast.GenDecl:
		if n.Tok != token.TYPE {
			break
		}
		ts := n.Specs[0].(*ast.TypeSpec)
		if ts.Name.Name == "B" {
			fields := ts.Type.(*ast.StructType).Fields
			addStructField(v.fset, fields, v.file, "int", "D", "D comment")
			addStructField(v.fset, fields, v.file, "float64", "E", "E comment")
		}
	}

	return v
}

func addStructField(fset *token.FileSet, fields *ast.FieldList, file *ast.File, typ string, name string, comment string) {
	prevField := fields.List[fields.NumFields()-1]

	c := &ast.Comment{Text: fmt.Sprint("// ", comment), Slash: prevField.End() + 1}
	cg := &ast.CommentGroup{List: []*ast.Comment{c}}
	o := ast.NewObj(ast.Var, name)
	f := &ast.Field{
		Doc:   cg,
		Names: []*ast.Ident{&ast.Ident{Name: name, Obj: o, NamePos: cg.End() + 1}},
	}
	o.Decl = f
	f.Type = &ast.Ident{Name: typ, NamePos: f.Names[0].End() + 1}

	fset.File(c.End()).AddLine(int(c.End()))
	fset.File(f.End()).AddLine(int(f.End()))

	fields.List = append(fields.List, f)
	file.Comments = append(file.Comments, cg)
}
