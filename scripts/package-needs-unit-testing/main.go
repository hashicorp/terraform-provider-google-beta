package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(os.Args)
		fmt.Fprintf(os.Stderr, "Usage: %s <package_paths>\n", os.Args[0])
		os.Exit(1)
	}

	packagePaths := os.Args[1:] // All arguments are now package paths

	for _, packagePath := range packagePaths {
		processPackage(packagePath)
	}

}

func processPackage(packagePath string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, packagePath, func(info os.FileInfo) bool {
		return !info.IsDir() && filepath.Ext(info.Name()) == ".go" && filepath.Base(info.Name()) != "skip_test.go"
	}, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing package: %v\n", err)
		return
	}

	nonVcrTests := []string{} // Keep track of test function names
	hasNonVcrTest := false

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {
					if isTestFunc(fn) {
						if !callsVcrTest(fn, fset) {
							hasNonVcrTest = true
							nonVcrTests = append(nonVcrTests, fn.Name.Name) // Store the function name
						}
					}
				}
			}
		}
	}
	if hasNonVcrTest {
		fmt.Printf("TESTABLE:%s:%s\n", packagePath, strings.Join(nonVcrTests, ","))
	} else {
		fmt.Printf("SKIPPED:%s\n", packagePath)
	}
}

func isTestFunc(fn *ast.FuncDecl) bool {
	return fn.Name != nil && fn.Name.Name != "" && len(fn.Name.Name) > 4 && fn.Name.Name[:4] == "Test"
}

func callsVcrTest(fn *ast.FuncDecl, fset *token.FileSet) bool {
	if fn.Body == nil {
		return false
	}
	found := false
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		xIdent, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		if sel.Sel == nil {
			return true
		}
		if xIdent.Name == "acctest" && sel.Sel.Name == "VcrTest" {
			found = true
			return false
		}
		return true
	})

	return found
}
