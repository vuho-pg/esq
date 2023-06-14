package main

import . "github.com/vuho-pg/esq/codegen/builder"

func joiningBuild() error {
	return Run(
		Wrap("Nested", Struct("Nested",
			Field("Path").String().Required(),
			Field("Query").Type("Query").Required(),
			Field("ScoreMode").Enum(),
			Field("IgnoreUnmapped").Bool(),
		),
		).Implement(queryInterface),
	)
}
