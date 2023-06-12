package esq

type MatchBuilder struct {
    Query any `json:"query"`
    Analyzer *string `json:"analyzer,omitempty"`
    AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`
    Fuzziness *int `json:"fuzziness,omitempty"`
    MaxExpansions *int `json:"max_expansions,omitempty"`
    PrefixLength *int `json:"prefix_length,omitempty"`
    FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`
    FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`
    Lenient *bool `json:"lenient,omitempty"`
    Operator *string `json:"operator,omitempty"`
    MinimumShouldMatch *string `json:"minimum_should_match,omitempty"`
    ZeroTermsQuery *string `json:"zero_terms_query,omitempty"`
}
func (m *MatchBuilder) IsQuery() {}


func Match(
_query any,
) *MatchBuilder {
    return &MatchBuilder{
       Query: _query,
    }
}
func (m *MatchBuilder) SetQuery(_query any) *MatchBuilder {
    m.Query = _query
    return m
}
func (m *MatchBuilder) SetAnalyzer(_analyzer string) *MatchBuilder {
    m.Analyzer = &_analyzer
    return m
}
func (m *MatchBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MatchBuilder {
    m.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
    return m
}
func (m *MatchBuilder) SetFuzziness(_fuzziness int) *MatchBuilder {
    m.Fuzziness = &_fuzziness
    return m
}
func (m *MatchBuilder) SetMaxExpansions(_maxExpansions int) *MatchBuilder {
    m.MaxExpansions = &_maxExpansions
    return m
}
func (m *MatchBuilder) SetPrefixLength(_prefixLength int) *MatchBuilder {
    m.PrefixLength = &_prefixLength
    return m
}
func (m *MatchBuilder) SetFuzzyTranspositions(_fuzzyTranspositions bool) *MatchBuilder {
    m.FuzzyTranspositions = &_fuzzyTranspositions
    return m
}
func (m *MatchBuilder) SetFuzzyRewrite(_fuzzyRewrite string) *MatchBuilder {
    m.FuzzyRewrite = &_fuzzyRewrite
    return m
}
func (m *MatchBuilder) SetLenient(_lenient bool) *MatchBuilder {
    m.Lenient = &_lenient
    return m
}
func (m *MatchBuilder) SetOperator(_operator string) *MatchBuilder {
    m.Operator = &_operator
    return m
}
func (m *MatchBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *MatchBuilder {
    m.MinimumShouldMatch = &_minimumShouldMatch
    return m
}
func (m *MatchBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchBuilder {
    m.ZeroTermsQuery = &_zeroTermsQuery
    return m
}