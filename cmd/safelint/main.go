package main

import (
	"github.com/Skylock-ai/safe/safelint"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(safelint.UnwrapFmtAnalyzer, safelint.SafeHandlerAnalyzer)
}
