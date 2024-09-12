package main

import (
	"github.com/robert-macwha/safe/safelint"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(safelint.UnwrapFmtAnalyzer, safelint.SafeHandlerAnalyzer)
}
