package esq

type QueryStringInnerBuilder struct {
	Query_                           string    `json:"query"`
	DefaultField_                    *string   `json:"default_field,omitempty"`
	AllowLeadingWildcard_            *bool     `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard_                 *bool     `json:"analyze_wildcard,omitempty"`
	Analyzer_                        *string   `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool     `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost_                           *float64  `json:"boost,omitempty"`
	DefaultOperator_                 *Operator `json:"default_operator,omitempty"`
	EnablePositionIncrements_        *bool     `json:"enable_position_increments,omitempty"`
	Fields_                          []string  `json:"fields,omitempty"`
	Fuzziness_                       *string   `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions_              *int      `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength_               *int      `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions_             *bool     `json:"fuzzy_transpositions,omitempty"`
	Lenient_                         *bool     `json:"lenient,omitempty"`
	MaxDeterminizedStates_           *int      `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch_              *string   `json:"minimum_should_match,omitempty"`
	QuoteAnalyzer_                   *string   `json:"quote_analyzer,omitempty"`
	PhraseSlop_                      *int      `json:"phrase_slop,omitempty"`
	QuoteFieldSuffix_                *string   `json:"quote_field_suffix,omitempty"`
	Rewrite_                         *string   `json:"rewrite,omitempty"`
	TimeZone_                        *string   `json:"time_zone,omitempty"`
}

type QueryStringBuilder struct {
	fieldName   string
	QueryString map[string]*QueryStringInnerBuilder `json:"query_string"`
}

func (*QueryStringBuilder) IsQuery() {}

func QueryString(
	_fieldName string,
	_query string,
) *QueryStringBuilder {
	return &QueryStringBuilder{
		fieldName: _fieldName,
		QueryString: map[string]*QueryStringInnerBuilder{
			_fieldName: {
				Query_: _query,
			},
		},
	}
}

func (_queryString *QueryStringBuilder) Query(_query string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Query_ = _query
	return _queryString
}
func (_queryString *QueryStringBuilder) DefaultField(_defaultField string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].DefaultField_ = &_defaultField
	return _queryString
}
func (_queryString *QueryStringBuilder) AllowLeadingWildcard(_allowLeadingWildcard bool) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].AllowLeadingWildcard_ = &_allowLeadingWildcard
	return _queryString
}
func (_queryString *QueryStringBuilder) AnalyzeWildcard(_analyzeWildcard bool) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].AnalyzeWildcard_ = &_analyzeWildcard
	return _queryString
}
func (_queryString *QueryStringBuilder) Analyzer(_analyzer string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Analyzer_ = &_analyzer
	return _queryString
}
func (_queryString *QueryStringBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return _queryString
}
func (_queryString *QueryStringBuilder) Boost(_boost float64) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Boost_ = &_boost
	return _queryString
}
func (_queryString *QueryStringBuilder) DefaultOperator(_defaultOperator Operator) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].DefaultOperator_ = &_defaultOperator
	return _queryString
}
func (_queryString *QueryStringBuilder) EnablePositionIncrements(_enablePositionIncrements bool) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].EnablePositionIncrements_ = &_enablePositionIncrements
	return _queryString
}
func (_queryString *QueryStringBuilder) Fields(_fields ...string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Fields_ = _fields
	return _queryString
}
func (_queryString *QueryStringBuilder) Fuzziness(_fuzziness string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Fuzziness_ = &_fuzziness
	return _queryString
}
func (_queryString *QueryStringBuilder) FuzzyMaxExpansions(_fuzzyMaxExpansions int) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].FuzzyMaxExpansions_ = &_fuzzyMaxExpansions
	return _queryString
}
func (_queryString *QueryStringBuilder) FuzzyPrefixLength(_fuzzyPrefixLength int) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].FuzzyPrefixLength_ = &_fuzzyPrefixLength
	return _queryString
}
func (_queryString *QueryStringBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return _queryString
}
func (_queryString *QueryStringBuilder) Lenient(_lenient bool) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Lenient_ = &_lenient
	return _queryString
}
func (_queryString *QueryStringBuilder) MaxDeterminizedStates(_maxDeterminizedStates int) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].MaxDeterminizedStates_ = &_maxDeterminizedStates
	return _queryString
}
func (_queryString *QueryStringBuilder) MinimumShouldMatch(_minimumShouldMatch string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return _queryString
}
func (_queryString *QueryStringBuilder) QuoteAnalyzer(_quoteAnalyzer string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].QuoteAnalyzer_ = &_quoteAnalyzer
	return _queryString
}
func (_queryString *QueryStringBuilder) PhraseSlop(_phraseSlop int) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].PhraseSlop_ = &_phraseSlop
	return _queryString
}
func (_queryString *QueryStringBuilder) QuoteFieldSuffix(_quoteFieldSuffix string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].QuoteFieldSuffix_ = &_quoteFieldSuffix
	return _queryString
}
func (_queryString *QueryStringBuilder) Rewrite(_rewrite string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].Rewrite_ = &_rewrite
	return _queryString
}
func (_queryString *QueryStringBuilder) TimeZone(_timeZone string) *QueryStringBuilder {
	_queryString.QueryString[_queryString.fieldName].TimeZone_ = &_timeZone
	return _queryString
}
