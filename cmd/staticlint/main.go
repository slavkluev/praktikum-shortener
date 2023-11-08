package main

import (
	"github.com/slavkluev/praktikum-shortener/multichecker/osexitanalyzer"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/quickfix"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"
)

func main() {
	mychecks := []*analysis.Analyzer{
		osexitanalyzer.OsExitAnalyzer,
		printf.Analyzer,
		structtag.Analyzer,
	}

	for _, v := range staticcheck.Analyzers {
		mychecks = append(mychecks, v.Analyzer)
	}

	for _, v := range simple.Analyzers {
		mychecks = append(mychecks, v.Analyzer)
	}

	for _, v := range stylecheck.Analyzers {
		mychecks = append(mychecks, v.Analyzer)
	}

	for _, v := range quickfix.Analyzers {
		mychecks = append(mychecks, v.Analyzer)
	}

	multichecker.Main(
		mychecks...,
	)
}
