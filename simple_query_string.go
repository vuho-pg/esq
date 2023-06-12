package esq

type SimpleQueryStringBuilder struct {
    Query string `json:"query"`
    Fields *[]string `json:"fields,omitempty"`
    DefaultOperator *string `json:"default_operator,omitempty"`
    AnalyzeWildCard *bool `json:"analyze_wild_card,omitempty"`
    Analyzer *string `json:"analyzer,omitempty"`
    AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`
    Flags *string `json:"flags,omitempty"`
    FuzzyMaxExpansions *int `json:"fuzzy_max_expansions,omitempty"`
    FuzzyPrefixLength *int `json:"fuzzy_prefix_length,omitempty"`
    FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`
    Lenient *bool `json:"lenient,omitempty"`
    MinimumShouldMatch *string `json:"minimum_should_match,omitempty"`
    QuoteFieldSuffix *string `json:"quote_field_suffix,omitempty"`
}
func (s *SimpleQueryStringBuilder) IsQuery() {}


func SimpleQueryString(
_query string,
) *SimpleQueryStringBuilder {
    return &SimpleQueryStringBuilder{
       Query: _query,
    }
}
func (s *SimpleQueryStringBuilder) SetQuery(_query string) *SimpleQueryStringBuilder {
    s.Query = _query
    return s
}
func (s *SimpleQueryStringBuilder) SetFields(_fields ...string) *SimpleQueryStringBuilder {
    s.Fields = &_fields
    return s
}
func (s *SimpleQueryStringBuilder) SetDefaultOperator(_defaultOperator string) *SimpleQueryStringBuilder {
    s.DefaultOperator = &_defaultOperator
    return s
}
func (s *SimpleQueryStringBuilder) SetAnalyzeWildCard(_analyzeWildCard bool) *SimpleQueryStringBuilder {
    s.AnalyzeWildCard = &_analyzeWildCard
    return s
}
func (s *SimpleQueryStringBuilder) SetAnalyzer(_analyzer string) *SimpleQueryStringBuilder {
    s.Analyzer = &_analyzer
    return s
}
func (s *SimpleQueryStringBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *SimpleQueryStringBuilder {
    s.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
    return s
}
func (s *SimpleQueryStringBuilder) SetFlags(_flags string) *SimpleQueryStringBuilder {
    s.Flags = &_flags
    return s
}
func (s *SimpleQueryStringBuilder) SetFuzzyMaxExpansions(_fuzzyMaxExpansions int) *SimpleQueryStringBuilder {
    s.FuzzyMaxExpansions = &_fuzzyMaxExpansions
    return s
}
func (s *SimpleQueryStringBuilder) SetFuzzyPrefixLength(_fuzzyPrefixLength int) *SimpleQueryStringBuilder {
    s.FuzzyPrefixLength = &_fuzzyPrefixLength
    return s
}
func (s *SimpleQueryStringBuilder) SetFuzzyTranspositions(_fuzzyTranspositions bool) *SimpleQueryStringBuilder {
    s.FuzzyTranspositions = &_fuzzyTranspositions
    return s
}
func (s *SimpleQueryStringBuilder) SetLenient(_lenient bool) *SimpleQueryStringBuilder {
    s.Lenient = &_lenient
    return s
}
func (s *SimpleQueryStringBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *SimpleQueryStringBuilder {
    s.MinimumShouldMatch = &_minimumShouldMatch
    return s
}
func (s *SimpleQueryStringBuilder) SetQuoteFieldSuffix(_quoteFieldSuffix string) *SimpleQueryStringBuilder {
    s.QuoteFieldSuffix = &_quoteFieldSuffix
    return s
}