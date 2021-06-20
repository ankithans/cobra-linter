package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	v := visitor{fset: token.NewFileSet()}
	for _, filePath := range os.Args[1:] {
		if filePath == "--" { // to be able to run like this "go run main.go --input.go"
			continue
		}

		f, err := parser.ParseFile(v.fset, filePath, nil, 0)
		if err != nil {
			log.Fatalf("Failed to parse file %s: %s", filePath, err)
		}

		ast.Walk(&v, f)
	}
}

type visitor struct {
	fset *token.FileSet
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	// _, ok := node.(*ast.AssignStmt)
	// if !ok {
	// 	return v
	// }

	// selectorExpr, ok := node.(*ast.SelectorExpr)
	// if !ok {
	// 	return v
	// }

	// if strings.Contains(selectorExpr.Sel.String(), "Command") {
	// ident, ok := node.(*ast.Ident)
	// if !ok {
	// 	return v
	// }
	// fmt.Println(ident)
	// }

	if node == nil {
		return nil
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, v.fset, node)
	fmt.Printf("%s: \n- %#v \n\n", buf.String(), node)

	return v
}
