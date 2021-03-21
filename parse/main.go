package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileName := "./main.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	for _, imp := range f.Imports {
		if strings.Contains(strings.ReplaceAll(imp.Path.Value, "\"", ""), "go") {
			imp.EndPos = imp.End()
			imp.Path.Value = strconv.Quote("git")
		}
	}

	//var output []byte
	//buffer := bytes.NewBuffer(output)

	//printer.Fprint(buffer, fset, f)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}

	fmt.Println(string(buf.Bytes()))

	if err := ioutil.WriteFile("./1.go", buf.Bytes(), 0666); err != nil {
		panic(err)
	}

	//fmt.Println(string(buffer.Bytes()))
}
