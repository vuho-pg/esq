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
		Wrap("Boosting", Struct("Boosting",
			Field("Positive").Type("Query").Required(),
			Field("Negative").Type("Query").Required(),
			Field("NegativeBoost").Float64().Required(),
		)).Implement(queryInterface),
		Wrap("ConstantScore", Struct("ConstantScore",
			Field("Filter").Type("Query").Required(),
			Field("Boost").Float64(),
		)).Implement(queryInterface),
		Wrap("DisMax", Struct("DisMax",
			Field("Queries").Type("Query").Array(),
			Field("TieBreaker").Float64(),
		)).Implement(queryInterface),
	)
}
