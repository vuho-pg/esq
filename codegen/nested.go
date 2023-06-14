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
		Wrap("HasChild", Struct("HasChild",
			Field("Type").String().Required(),
			Field("Query").Type("Query").Required(),
			Field("IgnoreUnmapped").Bool(),
			Field("MaxChildren").Int(),
			Field("MinChildren").Int(),
			Field("ScoreMode").Enum(),
		)).Implement(queryInterface),
		Wrap("HasParent", Struct("HasParent",
			Field("ParentType").String().Required(),
			Field("Query").Type("Query").Required(),
			Field("Score").Bool(),
			Field("IgnoreUnmapped").Bool(),
		)).Implement(queryInterface),
		Wrap("ParentID", Struct("ParentID",
			Field("Type").String().Required(),
			Field("ID").String().Required(),
			Field("IgnoreUnmapped").Bool(),
		)).Implement(queryInterface),
	)
}
