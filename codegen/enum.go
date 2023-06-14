package main

import . "github.com/vuho-pg/esq/codegen/builder"

func enumBuild() error {
	return Run(
		Enum("Operator", "string").Value("OR", "AND"),
		Enum("ZeroTermsQuery", "string").Value("none", "all"),
		Enum("ScoreMode", "string").Value("avg", "max", "min", "none", "sum"),
	)
}
