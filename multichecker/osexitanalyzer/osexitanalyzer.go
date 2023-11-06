// Package osexitanalyzer содержит правило OsExitAnalyzer
package osexitanalyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const (
	PackageMain  = "main"
	FunctionMain = "main"
)

var OsExitAnalyzer = &analysis.Analyzer{
	Name: "osexitanalyzer",
	Doc:  "check for os.Exit is in main function",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		if file.Name.Name != PackageMain {
			continue
		}

		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.FuncDecl:
				if x.Name.Name != FunctionMain {
					return false
				}
			case *ast.CallExpr:
				if isOsExitFunction(x) {
					pass.Reportf(x.Pos(), "os.Exit is in main function")
				}
			}

			return true
		})
	}
	return nil, nil
}

func isOsExitFunction(c *ast.CallExpr) bool {
	expr, ok := c.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := expr.X.(*ast.Ident)
	if !ok {
		return false
	}

	if ident.Name != "os" || expr.Sel.Name != "Exit" {
		return false
	}

	return true
}
