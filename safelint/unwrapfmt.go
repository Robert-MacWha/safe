package safelint

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var UnwrapFmtAnalyzer = &analysis.Analyzer{
	Name: "unwrapfmt",
	Doc:  "Linter to ensure that UnwrapFmt is called with a single '%%w' verb",
	Run:  runUnwrapFmt,
}

func runUnwrapFmt(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// Look for function calls
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isResultReceiver(callExpr, pass) {
				return true
			}

			if !isUnwrapFmtCall(callExpr, pass) {
				return true
			}

			// Check that params[0] has a single %w verb
			lit, ok := callExpr.Args[0].(*ast.BasicLit)
			if !ok {
				return true
			}

			if strings.Count(strings.Trim(lit.Value, `"`), "%w") != 1 {
				pass.Reportf(lit.Pos(), "UnwrapFmt should be called with a single %%w verb")
				return false
			}

			if strings.Count(strings.Trim(lit.Value, `"`), "%") != 1 {
				pass.Reportf(lit.Pos(), "UnwrapFmt should be called with a single  %%w verb")
				return false
			}

			return true
		})
	}
	return nil, nil
}

func isUnwrapFmtCall(callExpr *ast.CallExpr, pass *analysis.Pass) bool {
	funIdent, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	selection := pass.TypesInfo.Selections[funIdent]
	if selection == nil {
		return false
	}

	// Get function signature
	funcObj, ok := selection.Obj().(*types.Func)
	if !ok || funcObj.Name() != "UnwrapFmt" {
		return false
	}

	sig, ok := funcObj.Type().(*types.Signature)
	if !ok {
		return false
	}

	if sig.Recv() == nil {
		return false
	}

	// Check parmas
	params := sig.Params()
	if params.Len() != 1 {
		return false
	}

	if !types.Identical(params.At(0).Type(), types.Typ[types.String]) {
		return false
	}

	// Check return
	if sig.Results().Len() != 1 {
		return false
	}

	if !types.Identical(sig.Results().At(0).Type(), types.Universe.Lookup("error").Type()) {
		return false
	}

	return true
}
