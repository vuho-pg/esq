package main

type Builder interface {
	Build() error
}

func main() {
	CreateType(
		Interface("Query").Method("IsQuery"))
	// TODO: Compound query
	CreateType(
		Type("Bool",
			Field("Must", "[]Query"),
			Field("Filter", "[]Query"),
			Field("MustNot", "[]Query"),
			Field("Should", "[]Query"),
			Field("MinimumShouldMatch", "int"),
			Field("Boost", "float64"),
		).Method("IsQuery"),
	)
	// Full text queries
	CreateType(
		// TODO: Interval
		Type("Match",
			Field("Query", "any").Require(),
			Field("Analyzer", "string"),
			Field("AutoGenerateSynonymsPhraseQuery", "bool"),
			Field("Fuzziness", "int"),
			Field("MaxExpansions", "int"),
			Field("PrefixLength", "int"),
			Field("FuzzyTranspositions", "bool"),
			Field("FuzzyRewrite", "string"),
			Field("Lenient", "bool"),
			Field("Operator", "string"),
			Field("MinimumShouldMatch", "string"),
			Field("ZeroTermsQuery", "string"),
		).Method("IsQuery"),
		Type("Term",
			Field("Value", "string").Require(),
			Field("Boost", "float64"),
			Field("CaseInsensitive", "bool"),
		).Method("IsQuery"),
		Type("Range",
			Field("GT", "any"),
			Field("GTE", "any"),
			Field("LT", "any"),
			Field("LTE", "any"),
			Field("Format", "string"),
		).Method("IsQuery"),
		Type("MatchPhraseField",
			Field("Query", "string").Require(),
			Field("Analyzer", "string"),
			Field("ZeroTermsQuery", "string"),
		).Method("IsQuery"),
		MapType("MatchPhrase", "string", "*MatchPhraseFieldBuilder").
			Method("IsQuery"),
		Type("MatchPhrasePrefixField",
			Field("Query", "string").Require(),
			Field("Analyzer", "string"),
			Field("MaxExpansions", "int"),
			Field("Slop", "int"),
			Field("ZeroTermsQuery", "string"),
		).Method("IsQuery"),
		MapType("MatchPhrasePrefix", "string", "*MatchPhrasePrefixFieldBuilder").
			Method("IsQuery"),
		Type("CombinedFieldsField",
			Field("Fields", "[]string").Require(),
			Field("Query", "string").Require(),
			Field("AutoGenerateSynonymsPhraseQuery", "bool"),
			Field("Operator", "string"),
			Field("MinimumShouldMatch", "string"),
			Field("ZeroTermsQuery", "string"),
		).Method("IsQuery"),
		MapType("CombinedFields", "string", "*CombinedFieldsFieldBuilder").
			Method("IsQuery"),
		Type("MultiMatchField",
			Field("Fields", "[]string").Require(),
			Field("Query", "string").Require(),
			Field("Type", "string").Default(`"best_fields"`),
			Field("TierBreaker", "float64"),
			Field("Analyzer", "string"),
			Field("Boost", "float64"),
			Field("Operator", "string"),
			Field("MinimumShouldMatch", "string"),
			Field("Fuzziness", "int"),
			Field("Lenient", "bool"),
			Field("PrefixLength", "int"),
			Field("MaxExpansions", "int"),
			Field("FuzzyRewrite", "string"),
			Field("ZeroTermsQuery", "string"),
			Field("AutoGenerateSynonymsPhraseQuery", "bool"),
			Field("FuzzyTranspositions", "bool"),
		).Method("IsQuery"),
		Type("QueryString",
			Field("Query", "string").Require(),
			Field("DefaultField", "string"),
			Field("AllowLeadingWildcard", "bool"),
			Field("AnalyzeWildcard", "bool"),
			Field("Analyzer", "string"),
			Field("AutoGenerateSynonymsPhraseQuery", "bool"),
			Field("Boost", "float64"),
			Field("DefaultOperator", "string"),
			Field("EnablePositionIncrements", "bool"),
			Field("Fields", "[]string"),
			Field("Fuzziness", "string"),
			Field("FuzzyMaxExpansions", "int"),
			Field("FuzzyPrefixLength", "int"),
			Field("FuzzyTranspositions", "bool"),
			Field("Lenient", "bool"),
			Field("MaxDeterminizedStates", "int"),
			Field("MinimumShouldMatch", "string"),
			Field("QuoteAnalyzer", "string"),
			Field("PhraseSlop", "int"),
			Field("QuoteFieldSuffix", "string"),
			Field("Rewrite", "string"),
			Field("TimeZone", "string"),
		).Method("IsQuery"),
		Type("SimpleQueryString",
			Field("Query", "string").Require(),
			Field("Fields", "[]string"),
			Field("DefaultOperator", "string"),
			Field("AnalyzeWildCard", "bool"),
			Field("Analyzer", "string"),
			Field("AutoGenerateSynonymsPhraseQuery", "bool"),
			Field("Flags", "string"),
			Field("FuzzyMaxExpansions", "int"),
			Field("FuzzyPrefixLength", "int"),
			Field("FuzzyTranspositions", "bool"),
			Field("Lenient", "bool"),
			Field("MinimumShouldMatch", "string"),
			Field("QuoteFieldSuffix", "string"),
		).Method("IsQuery"),
	)
	// TODO: Geo query

	for _, builder := range typeGenerators {
		if err := builder.Build(); err != nil {
			panic(err)
		}
	}

}
