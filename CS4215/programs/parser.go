package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := new(token.FileSet)
	f, _ := parser.ParseFile(fset, "./hello/hello.go", nil, 0)
	ast.Print(fset, f)
}
