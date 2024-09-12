package gcl

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/robert-macwha/safe/internal/safelint"
	"golang.org/x/tools/go/analysis"
)

type Plugin struct {
}

func init() {
	register.Plugin("safelint", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

func (f *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		safelint.UnwrapFmtAnalyzer,
		safelint.SafeHandlerAnalyzer,
	}, nil
}

func (f *Plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
