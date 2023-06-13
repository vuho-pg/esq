package esq

type QueryStringInnerBuilder struct {
	Query_                           string   `json:"query"`
	DefaultField_                    *string  `json:"default_field,omitempty"`
	AllowLeadingWildcard_            *bool    `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard_                 *bool    `json:"analyze_wildcard,omitempty"`
	Analyzer_                        *string  `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool    `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost_                           *float64 `json:"boost,omitempty"`
	DefaultOperator_                 *string  `json:"default_operator,omitempty"`
	EnablePositionIncrements_        *bool    `json:"enable_position_increments,omitempty"`
	Fields_                          []string `json:"fields,omitempty"`
	Fuzziness_                       *string  `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions_              *int     `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength_               *int     `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions_             *bool    `json:"fuzzy_transpositions,omitempty"`
	Lenient_                         *bool    `json:"lenient,omitempty"`
	MaxDeterminizedStates_           *int     `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch_              *string  `json:"minimum_should_match,omitempty"`
	QuoteAnalyzer_                   *string  `json:"quote_analyzer,omitempty"`
	PhraseSlop_                      *int     `json:"phrase_slop,omitempty"`
	QuoteFieldSuffix_                *string  `json:"quote_field_suffix,omitempty"`
	Rewrite_                         *string  `json:"rewrite,omitempty"`
	TimeZone_                        *string  `json:"time_zone,omitempty"`
}

type QueryStringBuilder struct {
	fieldName   string
	QueryString map[string]*QueryStringInnerBuilder `json:"query_string"`
}

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

func (queryString *QueryStringBuilder) Query(_query string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Query_ = _query
	return queryString
}
func (queryString *QueryStringBuilder) DefaultField(_defaultField string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].DefaultField_ = &_defaultField
	return queryString
}
func (queryString *QueryStringBuilder) AllowLeadingWildcard(_allowLeadingWildcard bool) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].AllowLeadingWildcard_ = &_allowLeadingWildcard
	return queryString
}
func (queryString *QueryStringBuilder) AnalyzeWildcard(_analyzeWildcard bool) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].AnalyzeWildcard_ = &_analyzeWildcard
	return queryString
}
func (queryString *QueryStringBuilder) Analyzer(_analyzer string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Analyzer_ = &_analyzer
	return queryString
}
func (queryString *QueryStringBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return queryString
}
func (queryString *QueryStringBuilder) Boost(_boost float64) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Boost_ = &_boost
	return queryString
}
func (queryString *QueryStringBuilder) DefaultOperator(_defaultOperator string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].DefaultOperator_ = &_defaultOperator
	return queryString
}
func (queryString *QueryStringBuilder) EnablePositionIncrements(_enablePositionIncrements bool) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].EnablePositionIncrements_ = &_enablePositionIncrements
	return queryString
}
func (queryString *QueryStringBuilder) Fields(_fields ...string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Fields_ = _fields
	return queryString
}
func (queryString *QueryStringBuilder) Fuzziness(_fuzziness string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Fuzziness_ = &_fuzziness
	return queryString
}
func (queryString *QueryStringBuilder) FuzzyMaxExpansions(_fuzzyMaxExpansions int) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].FuzzyMaxExpansions_ = &_fuzzyMaxExpansions
	return queryString
}
func (queryString *QueryStringBuilder) FuzzyPrefixLength(_fuzzyPrefixLength int) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].FuzzyPrefixLength_ = &_fuzzyPrefixLength
	return queryString
}
func (queryString *QueryStringBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return queryString
}
func (queryString *QueryStringBuilder) Lenient(_lenient bool) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Lenient_ = &_lenient
	return queryString
}
func (queryString *QueryStringBuilder) MaxDeterminizedStates(_maxDeterminizedStates int) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].MaxDeterminizedStates_ = &_maxDeterminizedStates
	return queryString
}
func (queryString *QueryStringBuilder) MinimumShouldMatch(_minimumShouldMatch string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return queryString
}
func (queryString *QueryStringBuilder) QuoteAnalyzer(_quoteAnalyzer string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].QuoteAnalyzer_ = &_quoteAnalyzer
	return queryString
}
func (queryString *QueryStringBuilder) PhraseSlop(_phraseSlop int) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].PhraseSlop_ = &_phraseSlop
	return queryString
}
func (queryString *QueryStringBuilder) QuoteFieldSuffix(_quoteFieldSuffix string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].QuoteFieldSuffix_ = &_quoteFieldSuffix
	return queryString
}
func (queryString *QueryStringBuilder) Rewrite(_rewrite string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].Rewrite_ = &_rewrite
	return queryString
}
func (queryString *QueryStringBuilder) TimeZone(_timeZone string) *QueryStringBuilder {
	queryString.QueryString[queryString.fieldName].TimeZone_ = &_timeZone
	return queryString
}
