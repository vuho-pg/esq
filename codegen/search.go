package main

import . "github.com/vuho-pg/esq/codegen/builder"

func searchBuild() error {
	return Run(
		Enum("SortMode", "string").Value("min", "max", "sum", "avg", "median"),
		Enum("SortOrder", "string").Value("asc", "desc"),
		Enum("NumericType", "string").Value("double", "long", "date", "date_nanos"),
		Struct("SortValue",
			Field("Order").Type("SortOrder"),
			Field("Format").String(),
			Field("Mode").Type("SortMode"),
			Field("NumericType").Enum(),
			Field("Missing").Type("SortMissing"),
		).Extend("Nested"),
		Struct("Field",
			Field("Field").String().Required(),
			Field("Format").String(),
		),
		Struct("DocvalueField",
			Field("Field").String().Required(),
			Field("Format").String(),
		),
		Struct("PIT",
			Field("ID").String().Required(),
			Field("KeepAlive").String(),
		),
		KeyValue("Sort", "string", "*SortValueBuilder"),
		Struct("SortNested",
			Field("Path").String().Required(),
			Field("Nested").Type("*SortNestedBuilder"),
			Field("Filter").Type("Query"),
			Field("MaxChildren").Int(),
		),
		Enum("SortMissing", "string").Value("_first", "_last"),
		Struct("Search",
			Field("DocvalueFields").Type("DocvalueFieldBuilder").Array().ForcePtr(),
			Field("Fields").Type("FieldBuilder").Array().ForcePtr(),
			Field("Explain").Bool(),
			Field("From").Int(),
			//Field("IndicesBoost").
			Field("MinScore").Float64(),
			Field("PIT").Type("PITBuilder"),
			Field("Query").Type("Query"),
			//Field("RunetimeMappings").Type("RuntimeMappings"),
			Field("SeqNoPrimaryTerm").Bool(),
			Field("Size").Int(),
			//Field("_Source")
			Field("Stats").String().Array(),
			Field("TerminateAfter").Int(),
			Field("Timeout").String(),
			Field("Version").Bool(),
			Field("Sort").Type("SortNestedBuilder").Array().ForcePtr(),
		),
		// Async search

	)
}
