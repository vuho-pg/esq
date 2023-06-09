package esq

type QueryStringBuilder struct {
	Query                           string    `json:"query"`
	DefaultField                    *string   `json:"default_field,omitempty"`
	AllowLeadingWildcard            *bool     `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard                 *bool     `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string   `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool     `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float64  `json:"boost,omitempty"`
	DefaultOperator                 *string   `json:"default_operator,omitempty"`
	EnablePositionIncrements        *bool     `json:"enable_position_increments,omitempty"`
	Fields                          *[]string `json:"fields,omitempty"`
	Fuzziness                       *string   `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions              *int      `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int      `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions             *bool     `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool     `json:"lenient,omitempty"`
	MaxDeterminizedStates           *int      `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch              *string   `json:"minimum_should_match,omitempty"`
	QuoteAnalyzer                   *string   `json:"quote_analyzer,omitempty"`
	PhraseSlop                      *int      `json:"phrase_slop,omitempty"`
	QuoteFieldSuffix                *string   `json:"quote_field_suffix,omitempty"`
	Rewrite                         *string   `json:"rewrite,omitempty"`
	TimeZone                        *string   `json:"time_zone,omitempty"`
}

func QueryString(
	_query string,
) *QueryStringBuilder {
	return &QueryStringBuilder{
		Query: _query,
	}
}
func (q *QueryStringBuilder) SetQuery(_query string) *QueryStringBuilder {
	q.Query = _query
	return q
}
func (q *QueryStringBuilder) SetDefaultField(_defaultField string) *QueryStringBuilder {
	q.DefaultField = &_defaultField
	return q
}
func (q *QueryStringBuilder) SetAllowLeadingWildcard(_allowLeadingWildcard bool) *QueryStringBuilder {
	q.AllowLeadingWildcard = &_allowLeadingWildcard
	return q
}
func (q *QueryStringBuilder) SetAnalyzeWildcard(_analyzeWildcard bool) *QueryStringBuilder {
	q.AnalyzeWildcard = &_analyzeWildcard
	return q
}
func (q *QueryStringBuilder) SetAnalyzer(_analyzer string) *QueryStringBuilder {
	q.Analyzer = &_analyzer
	return q
}
func (q *QueryStringBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *QueryStringBuilder {
	q.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
	return q
}
func (q *QueryStringBuilder) SetBoost(_boost float64) *QueryStringBuilder {
	q.Boost = &_boost
	return q
}
func (q *QueryStringBuilder) SetDefaultOperator(_defaultOperator string) *QueryStringBuilder {
	q.DefaultOperator = &_defaultOperator
	return q
}
func (q *QueryStringBuilder) SetEnablePositionIncrements(_enablePositionIncrements bool) *QueryStringBuilder {
	q.EnablePositionIncrements = &_enablePositionIncrements
	return q
}
func (q *QueryStringBuilder) SetFields(_fields ...string) *QueryStringBuilder {
	q.Fields = &_fields
	return q
}
func (q *QueryStringBuilder) SetFuzziness(_fuzziness string) *QueryStringBuilder {
	q.Fuzziness = &_fuzziness
	return q
}
func (q *QueryStringBuilder) SetFuzzyMaxExpansions(_fuzzyMaxExpansions int) *QueryStringBuilder {
	q.FuzzyMaxExpansions = &_fuzzyMaxExpansions
	return q
}
func (q *QueryStringBuilder) SetFuzzyPrefixLength(_fuzzyPrefixLength int) *QueryStringBuilder {
	q.FuzzyPrefixLength = &_fuzzyPrefixLength
	return q
}
func (q *QueryStringBuilder) SetFuzzyTranspositions(_fuzzyTranspositions bool) *QueryStringBuilder {
	q.FuzzyTranspositions = &_fuzzyTranspositions
	return q
}
func (q *QueryStringBuilder) SetLenient(_lenient bool) *QueryStringBuilder {
	q.Lenient = &_lenient
	return q
}
func (q *QueryStringBuilder) SetMaxDeterminizedStates(_maxDeterminizedStates int) *QueryStringBuilder {
	q.MaxDeterminizedStates = &_maxDeterminizedStates
	return q
}
func (q *QueryStringBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *QueryStringBuilder {
	q.MinimumShouldMatch = &_minimumShouldMatch
	return q
}
func (q *QueryStringBuilder) SetQuoteAnalyzer(_quoteAnalyzer string) *QueryStringBuilder {
	q.QuoteAnalyzer = &_quoteAnalyzer
	return q
}
func (q *QueryStringBuilder) SetPhraseSlop(_phraseSlop int) *QueryStringBuilder {
	q.PhraseSlop = &_phraseSlop
	return q
}
func (q *QueryStringBuilder) SetQuoteFieldSuffix(_quoteFieldSuffix string) *QueryStringBuilder {
	q.QuoteFieldSuffix = &_quoteFieldSuffix
	return q
}
func (q *QueryStringBuilder) SetRewrite(_rewrite string) *QueryStringBuilder {
	q.Rewrite = &_rewrite
	return q
}
func (q *QueryStringBuilder) SetTimeZone(_timeZone string) *QueryStringBuilder {
	q.TimeZone = &_timeZone
	return q
}
