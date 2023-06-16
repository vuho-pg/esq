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
			Field("Nested").Type("SortNestedBuilder").ForcePtr(),
		),
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
			Field("Nested").Type("SortNestedBuilder").ForcePtr(),
			Field("Filter").Type("Query"),
			Field("MaxChildren").Int(),
		),
		Enum("SortMissing", "string").Value("_first", "_last"),
		Struct("InnerHits",
			Field("Name").String(),
			Field("From").Any(),
			Field("Size").Int(),
			Field("Sort").Type("SortBuilder").Array(),
			Field("Collapse").Type("CollapseBuilder").ForcePtr(),
		),
		Struct("Collapse",
			Field("Field").String().Required(),
			Field("InnerHits").Type("InnerHitsBuilder").Array().ForcePtr(),
			Field("MaxConcurrentGroupSearches").Int(),
		),
		Struct("Search",
			Field("TrackTotalHits").Bool(),
			Field("DocvalueFields").Type("DocvalueFieldBuilder").Array().ForcePtr(),
			Field("Fields").Type("FieldBuilder").Array().ForcePtr(),
			Field("Explain").Bool(),
			//Field("IndicesBoost").
			Field("MinScore").Float64(),
			Field("PIT").Type("PITBuilder"),
			Field("Query").Type("Query"),
			//Field("RunetimeMappings").Type("RuntimeMappings"),
			Field("SeqNoPrimaryTerm").Bool(),
			//Field("_Source")
			Field("Stats").String().Array(),
			Field("TerminateAfter").Int(),
			Field("Timeout").String(),
			Field("Version").Bool(),
			Field("Sort").Type("SortBuilder").Array(),
			Field("Collapse").Type("CollapseBuilder").ForcePtr(),
			// pagination
			Field("From").Int(),
			Field("Size").Int(),
			Field("SearchAfter").Any().Array(),
			//TODO: aggregations
			//TODO: highlight

		),
		// Async search

	)
}
