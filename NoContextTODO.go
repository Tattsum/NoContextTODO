package NoContextTODO

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "NoContextTODO is bad used context.TODO"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "NoContextTODO",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var ctxpkg *types.Package

func run(pass *analysis.Pass) (interface{}, error) {
	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() != "context" {
			return nil, nil
		}
		ctxpkg = pkg
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:

			// check function is context.TODO
			caller, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}

			if pass.TypesInfo.ObjectOf(caller.Sel).Pkg() != ctxpkg {
				return
			}

			if caller.Sel.Name != "TODO" {
				return
			}

			pass.Reportf(caller.Pos(), "NG")
		}
	})

	return nil, nil
}
