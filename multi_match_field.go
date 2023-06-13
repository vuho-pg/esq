package esq

type MultiMatchFieldBuilder struct {
    Fields []string `json:"fields"`
    Query string `json:"query"`
    Type *string `json:"type,omitempty"`
    TierBreaker *float64 `json:"tier_breaker,omitempty"`
    Analyzer *string `json:"analyzer,omitempty"`
    Boost *float64 `json:"boost,omitempty"`
    Operator *string `json:"operator,omitempty"`
    MinimumShouldMatch *string `json:"minimum_should_match,omitempty"`
    Fuzziness *int `json:"fuzziness,omitempty"`
    Lenient *bool `json:"lenient,omitempty"`
    PrefixLength *int `json:"prefix_length,omitempty"`
    MaxExpansions *int `json:"max_expansions,omitempty"`
    FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`
    ZeroTermsQuery *string `json:"zero_terms_query,omitempty"`
    AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`
    FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`
}
func (m *MultiMatchFieldBuilder) IsQuery() {}


func MultiMatchField(
_fields []string,
_query string,
) *MultiMatchFieldBuilder {
    return &MultiMatchFieldBuilder{
       Fields: _fields,
       Query: _query,
       Type: common.PtrOf("best_fields"),
    }
}
func (m *MultiMatchFieldBuilder) SetFields(_fields ...string) *MultiMatchFieldBuilder {
    m.Fields = _fields
    return m
}
func (m *MultiMatchFieldBuilder) SetQuery(_query string) *MultiMatchFieldBuilder {
    m.Query = _query
    return m
}
func (m *MultiMatchFieldBuilder) SetType(_type string) *MultiMatchFieldBuilder {
    m.Type = &_type
    return m
}
func (m *MultiMatchFieldBuilder) SetTierBreaker(_tierBreaker float64) *MultiMatchFieldBuilder {
    m.TierBreaker = &_tierBreaker
    return m
}
func (m *MultiMatchFieldBuilder) SetAnalyzer(_analyzer string) *MultiMatchFieldBuilder {
    m.Analyzer = &_analyzer
    return m
}
func (m *MultiMatchFieldBuilder) SetBoost(_boost float64) *MultiMatchFieldBuilder {
    m.Boost = &_boost
    return m
}
func (m *MultiMatchFieldBuilder) SetOperator(_operator string) *MultiMatchFieldBuilder {
    m.Operator = &_operator
    return m
}
func (m *MultiMatchFieldBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *MultiMatchFieldBuilder {
    m.MinimumShouldMatch = &_minimumShouldMatch
    return m
}
func (m *MultiMatchFieldBuilder) SetFuzziness(_fuzziness int) *MultiMatchFieldBuilder {
    m.Fuzziness = &_fuzziness
    return m
}
func (m *MultiMatchFieldBuilder) SetLenient(_lenient bool) *MultiMatchFieldBuilder {
    m.Lenient = &_lenient
    return m
}
func (m *MultiMatchFieldBuilder) SetPrefixLength(_prefixLength int) *MultiMatchFieldBuilder {
    m.PrefixLength = &_prefixLength
    return m
}
func (m *MultiMatchFieldBuilder) SetMaxExpansions(_maxExpansions int) *MultiMatchFieldBuilder {
    m.MaxExpansions = &_maxExpansions
    return m
}
func (m *MultiMatchFieldBuilder) SetFuzzyRewrite(_fuzzyRewrite string) *MultiMatchFieldBuilder {
    m.FuzzyRewrite = &_fuzzyRewrite
    return m
}
func (m *MultiMatchFieldBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MultiMatchFieldBuilder {
    m.ZeroTermsQuery = &_zeroTermsQuery
    return m
}
func (m *MultiMatchFieldBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MultiMatchFieldBuilder {
    m.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
    return m
}
func (m *MultiMatchFieldBuilder) SetFuzzyTranspositions(_fuzzyTranspositions bool) *MultiMatchFieldBuilder {
    m.FuzzyTranspositions = &_fuzzyTranspositions
    return m
}