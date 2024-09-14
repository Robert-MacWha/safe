package safelint

import (
	"go/ast"
	"go/types"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var SafeHandlerAnalyzer = &analysis.Analyzer{
	Name: "safehandler",
	Doc:  "Linter to ensure the safe handler is used when functions are unwrapped",
	Run:  runSafeHandler,
}

func runSafeHandler(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// Ignore if the file is a test file
			if isTestFile(pass.Fset.File(file.Pos()).Name()) {
				return true
			}

			// Look for function declarations
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			safeHandleCalled := false
			ast.Inspect(funcDecl.Body, func(bodyNode ast.Node) bool {
				callExpr, ok := bodyNode.(*ast.CallExpr)
				if !ok {
					return true
				}

				if isSafePkg(callExpr) && isHandleCall(callExpr) {
					safeHandleCalled = true
				}

				if isUnwrapCall(callExpr) && (isResultReceiver(callExpr, pass) || isOptionReciever(callExpr, pass)) {
					if !safeHandleCalled {
						pass.Reportf(callExpr.Pos(), "safe.Handle(&res) must be called before any unwrapping methods (Unwrap, UnwrapErr, UnwrapFmt)")
						return false
					}
				}

				return true
			})

			return true
		})
	}

	return nil, nil
}

func isSafePkg(callExpr *ast.CallExpr) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkgIdent, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	return pkgIdent.Name == "safe"
}

// Helper function to check if the receiver of the Unwrap call is a Result[T] type
func isResultReceiver(callExpr *ast.CallExpr, pass *analysis.Pass) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	// Get the type of the receiver
	recvType := pass.TypesInfo.TypeOf(sel.X)
	namedType, ok := recvType.(*types.Named)
	if !ok {
		return false
	}

	return namedType.Obj().Name() == "Result"
}

func isOptionReciever(callExpr *ast.CallExpr, pass *analysis.Pass) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	// Get the type of the receiver
	recvType := pass.TypesInfo.TypeOf(sel.X)

	namedType, ok := recvType.(*types.Named)
	if !ok {
		return false
	}

	return namedType.Obj().Name() == "Option"
}

func isHandleCall(callExpr *ast.CallExpr) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if sel.Sel.Name == "Handle" {
		return true
	}

	return false
}

// Helper function to check if the call is an Unwrap, UnwrapErr, or UnwrapFmt call
func isUnwrapCall(callExpr *ast.CallExpr) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	funcName := sel.Sel.Name
	return funcName == "Unwrap" || funcName == "UnwrapErr" || funcName == "Expect"
}

// Helper function to determine if a file is a test file
func isTestFile(filename string) bool {
	return strings.HasSuffix(filepath.Base(filename), "_test.go")
}
