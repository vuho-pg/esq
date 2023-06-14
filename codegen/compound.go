package main

import . "github.com/vuho-pg/esq/codegen/builder"

func compoundBuild() error {
	return Run(Wrap("Bool", Struct("Bool",
		Field("Must").Type("Query").Array(),
		Field("MustNot").Type("Query").Array(),
		Field("Should").Type("Query").Array(),
		Field("Filter").Type("Query").Array(),
		Field("MinimumShouldMatch").String(),
		Field("Boost").Float64(),
	)).Implement(queryInterface),
	)
}
