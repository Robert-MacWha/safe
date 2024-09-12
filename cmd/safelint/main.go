package main

import (
	"github.com/Robert-MacWha/safe/safelint"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(safelint.UnwrapFmtAnalyzer, safelint.SafeHandlerAnalyzer)
}
