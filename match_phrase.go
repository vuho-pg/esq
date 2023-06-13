package esq


type MatchPhraseBuilder struct {
    MatchPhraseField *MatchPhraseFieldBuilder `json:"match_phrase"`
}
func (m *MatchPhraseBuilder) IsQuery() {}


func MatchPhrase(
_query string,
) *MatchPhraseBuilder {
    return &MatchPhraseBuilder {
        MatchPhraseField:  &MatchPhraseFieldBuilder{
                                    Query: _query,
                                 },
        }
}
func (m *MatchPhraseBuilder) SetQuery(_query string) *MatchPhraseBuilder {
    m.MatchPhraseField.Query = _query
    return m
}
func (m *MatchPhraseBuilder) SetAnalyzer(_analyzer string) *MatchPhraseBuilder {
    m.MatchPhraseField.Analyzer = &_analyzer
    return m
}
func (m *MatchPhraseBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchPhraseBuilder {
    m.MatchPhraseField.ZeroTermsQuery = &_zeroTermsQuery
    return m
}