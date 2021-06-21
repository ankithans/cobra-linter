package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
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
	compositelit, ok := node.(*ast.CompositeLit)
	if !ok {
		return v
	}

	for i := range compositelit.Elts {
		kv, ok := compositelit.Elts[i].(*ast.KeyValueExpr)
		if !ok {
			return v
		}

		key := kv.Key.(*ast.Ident).Name
		val := kv.Value.(*ast.BasicLit).Value
		l := len(strings.Trim(val, "\""))

		switch key {
		case "Use":
			if l < 3 {
				fmt.Printf("Length of %s should be greater than 3", key)
				return v
			}
		case "Short":
			if l < 5 {
				fmt.Printf("Length of %s should be greater than 5", key)
				return v
			}
		case "Long":
			if l < 15 {
				fmt.Printf("Length of %s should be greater than 15", key)
				return v
			}
		case "Example":
			if l < 15 {
				fmt.Printf("Length of %s should be greater than 15", key)
				return v
			}
		default:
			panic("Wrong Type")
		}
	}

	return v
}
