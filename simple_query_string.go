package esq

type SimpleQueryStringInnerBuilder struct {
	Query_                           string   `json:"query"`
	Fields_                          []string `json:"fields,omitempty"`
	DefaultOperator_                 *string  `json:"default_operator,omitempty"`
	AnalyzeWildcard_                 *bool    `json:"analyze_wildcard,omitempty"`
	Analyzer_                        *string  `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool    `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Flags_                           *string  `json:"flags,omitempty"`
	FuzzyMaxExpansions_              *int     `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength_               *int     `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions_             *bool    `json:"fuzzy_transpositions,omitempty"`
	Lenient_                         *bool    `json:"lenient,omitempty"`
	MinimumShouldMatch_              *string  `json:"minimum_should_match,omitempty"`
	QuoteFieldSuffix_                *string  `json:"quote_field_suffix,omitempty"`
}

type SimpleQueryStringBuilder struct {
	fieldName         string
	SimpleQueryString map[string]*SimpleQueryStringInnerBuilder `json:"simple_query_string"`
}

func (*SimpleQueryStringBuilder) IsQuery() {}

func SimpleQueryString(
	_fieldName string,
	_query string,
) *SimpleQueryStringBuilder {
	return &SimpleQueryStringBuilder{
		fieldName: _fieldName,
		SimpleQueryString: map[string]*SimpleQueryStringInnerBuilder{
			_fieldName: {
				Query_: _query,
			},
		},
	}
}

func (simpleQueryString *SimpleQueryStringBuilder) Query(_query string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].Query_ = _query
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) Fields(_fields ...string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].Fields_ = _fields
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) DefaultOperator(_defaultOperator string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].DefaultOperator_ = &_defaultOperator
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) AnalyzeWildcard(_analyzeWildcard bool) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].AnalyzeWildcard_ = &_analyzeWildcard
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) Analyzer(_analyzer string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].Analyzer_ = &_analyzer
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) Flags(_flags string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].Flags_ = &_flags
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) FuzzyMaxExpansions(_fuzzyMaxExpansions int) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].FuzzyMaxExpansions_ = &_fuzzyMaxExpansions
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) FuzzyPrefixLength(_fuzzyPrefixLength int) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].FuzzyPrefixLength_ = &_fuzzyPrefixLength
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) Lenient(_lenient bool) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].Lenient_ = &_lenient
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) MinimumShouldMatch(_minimumShouldMatch string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return simpleQueryString
}
func (simpleQueryString *SimpleQueryStringBuilder) QuoteFieldSuffix(_quoteFieldSuffix string) *SimpleQueryStringBuilder {
	simpleQueryString.SimpleQueryString[simpleQueryString.fieldName].QuoteFieldSuffix_ = &_quoteFieldSuffix
	return simpleQueryString
}
