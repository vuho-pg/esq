package esq


type MatchPhrasePrefixBuilder struct {
    MatchPhrasePrefixField *MatchPhrasePrefixFieldBuilder `json:"match_phrase_prefix"`
}
func (m *MatchPhrasePrefixBuilder) IsQuery() {}


func MatchPhrasePrefix(
_query string,
) *MatchPhrasePrefixBuilder {
    return &MatchPhrasePrefixBuilder {
        MatchPhrasePrefixField:  &MatchPhrasePrefixFieldBuilder{
                                    Query: _query,
                                 },
        }
}
func (m *MatchPhrasePrefixBuilder) SetQuery(_query string) *MatchPhrasePrefixBuilder {
    m.MatchPhrasePrefixField.Query = _query
    return m
}
func (m *MatchPhrasePrefixBuilder) SetAnalyzer(_analyzer string) *MatchPhrasePrefixBuilder {
    m.MatchPhrasePrefixField.Analyzer = &_analyzer
    return m
}
func (m *MatchPhrasePrefixBuilder) SetMaxExpansions(_maxExpansions int) *MatchPhrasePrefixBuilder {
    m.MatchPhrasePrefixField.MaxExpansions = &_maxExpansions
    return m
}
func (m *MatchPhrasePrefixBuilder) SetSlop(_slop int) *MatchPhrasePrefixBuilder {
    m.MatchPhrasePrefixField.Slop = &_slop
    return m
}
func (m *MatchPhrasePrefixBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchPhrasePrefixBuilder {
    m.MatchPhrasePrefixField.ZeroTermsQuery = &_zeroTermsQuery
    return m
}