package esq


type MatchBuilder struct {
    MatchField *MatchFieldBuilder `json:"match"`
}
func (m *MatchBuilder) IsQuery() {}


func Match(
_query any,
) *MatchBuilder {
    return &MatchBuilder {
        MatchField:  &MatchFieldBuilder{
                                    Query: _query,
                                 },
        }
}
func (m *MatchBuilder) SetQuery(_query any) *MatchBuilder {
    m.MatchField.Query = _query
    return m
}
func (m *MatchBuilder) SetAnalyzer(_analyzer string) *MatchBuilder {
    m.MatchField.Analyzer = &_analyzer
    return m
}
func (m *MatchBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MatchBuilder {
    m.MatchField.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
    return m
}
func (m *MatchBuilder) SetFuzziness(_fuzziness int) *MatchBuilder {
    m.MatchField.Fuzziness = &_fuzziness
    return m
}
func (m *MatchBuilder) SetMaxExpansions(_maxExpansions int) *MatchBuilder {
    m.MatchField.MaxExpansions = &_maxExpansions
    return m
}
func (m *MatchBuilder) SetPrefixLength(_prefixLength int) *MatchBuilder {
    m.MatchField.PrefixLength = &_prefixLength
    return m
}
func (m *MatchBuilder) SetFuzzyTranspositions(_fuzzyTranspositions bool) *MatchBuilder {
    m.MatchField.FuzzyTranspositions = &_fuzzyTranspositions
    return m
}
func (m *MatchBuilder) SetFuzzyRewrite(_fuzzyRewrite string) *MatchBuilder {
    m.MatchField.FuzzyRewrite = &_fuzzyRewrite
    return m
}
func (m *MatchBuilder) SetLenient(_lenient bool) *MatchBuilder {
    m.MatchField.Lenient = &_lenient
    return m
}
func (m *MatchBuilder) SetOperator(_operator string) *MatchBuilder {
    m.MatchField.Operator = &_operator
    return m
}
func (m *MatchBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *MatchBuilder {
    m.MatchField.MinimumShouldMatch = &_minimumShouldMatch
    return m
}
func (m *MatchBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchBuilder {
    m.MatchField.ZeroTermsQuery = &_zeroTermsQuery
    return m
}