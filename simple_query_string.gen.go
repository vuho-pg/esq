package esq

type SimpleQueryStringInnerBuilder struct {
	Query_                           string    `json:"query"`
	Fields_                          []string  `json:"fields,omitempty"`
	DefaultOperator_                 *Operator `json:"default_operator,omitempty"`
	AnalyzeWildcard_                 *bool     `json:"analyze_wildcard,omitempty"`
	Analyzer_                        *string   `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool     `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Flags_                           *string   `json:"flags,omitempty"`
	FuzzyMaxExpansions_              *int      `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength_               *int      `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions_             *bool     `json:"fuzzy_transpositions,omitempty"`
	Lenient_                         *bool     `json:"lenient,omitempty"`
	MinimumShouldMatch_              *string   `json:"minimum_should_match,omitempty"`
	QuoteFieldSuffix_                *string   `json:"quote_field_suffix,omitempty"`
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

func (_simpleQueryString *SimpleQueryStringBuilder) Query(_query string) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].Query_ = _query
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) Fields(_fields ...string) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].Fields_ = _fields
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) DefaultOperator(_defaultOperator Operator) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].DefaultOperator_ = &_defaultOperator
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) AnalyzeWildcard(_analyzeWildcard bool) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].AnalyzeWildcard_ = &_analyzeWildcard
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) Analyzer(_analyzer string) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].Analyzer_ = &_analyzer
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) Flags(_flags string) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].Flags_ = &_flags
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) FuzzyMaxExpansions(_fuzzyMaxExpansions int) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].FuzzyMaxExpansions_ = &_fuzzyMaxExpansions
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) FuzzyPrefixLength(_fuzzyPrefixLength int) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].FuzzyPrefixLength_ = &_fuzzyPrefixLength
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) Lenient(_lenient bool) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].Lenient_ = &_lenient
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) MinimumShouldMatch(_minimumShouldMatch string) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return _simpleQueryString
}
func (_simpleQueryString *SimpleQueryStringBuilder) QuoteFieldSuffix(_quoteFieldSuffix string) *SimpleQueryStringBuilder {
	_simpleQueryString.SimpleQueryString[_simpleQueryString.fieldName].QuoteFieldSuffix_ = &_quoteFieldSuffix
	return _simpleQueryString
}
