package esq

type MatchFieldBuilder struct {
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
func (m *MatchFieldBuilder) IsQuery() {}


func MatchField(
_query any,
) *MatchFieldBuilder {
    return &MatchFieldBuilder{
       Query: _query,
    }
}
func (m *MatchFieldBuilder) SetQuery(_query any) *MatchFieldBuilder {
    m.Query = _query
    return m
}
func (m *MatchFieldBuilder) SetAnalyzer(_analyzer string) *MatchFieldBuilder {
    m.Analyzer = &_analyzer
    return m
}
func (m *MatchFieldBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MatchFieldBuilder {
    m.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
    return m
}
func (m *MatchFieldBuilder) SetFuzziness(_fuzziness int) *MatchFieldBuilder {
    m.Fuzziness = &_fuzziness
    return m
}
func (m *MatchFieldBuilder) SetMaxExpansions(_maxExpansions int) *MatchFieldBuilder {
    m.MaxExpansions = &_maxExpansions
    return m
}
func (m *MatchFieldBuilder) SetPrefixLength(_prefixLength int) *MatchFieldBuilder {
    m.PrefixLength = &_prefixLength
    return m
}
func (m *MatchFieldBuilder) SetFuzzyTranspositions(_fuzzyTranspositions bool) *MatchFieldBuilder {
    m.FuzzyTranspositions = &_fuzzyTranspositions
    return m
}
func (m *MatchFieldBuilder) SetFuzzyRewrite(_fuzzyRewrite string) *MatchFieldBuilder {
    m.FuzzyRewrite = &_fuzzyRewrite
    return m
}
func (m *MatchFieldBuilder) SetLenient(_lenient bool) *MatchFieldBuilder {
    m.Lenient = &_lenient
    return m
}
func (m *MatchFieldBuilder) SetOperator(_operator string) *MatchFieldBuilder {
    m.Operator = &_operator
    return m
}
func (m *MatchFieldBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *MatchFieldBuilder {
    m.MinimumShouldMatch = &_minimumShouldMatch
    return m
}
func (m *MatchFieldBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchFieldBuilder {
    m.ZeroTermsQuery = &_zeroTermsQuery
    return m
}