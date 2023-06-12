package esq

type MatchPhraseFieldBuilder struct {
    Query string `json:"query"`
    Analyzer *string `json:"analyzer,omitempty"`
    ZeroTermsQuery *string `json:"zero_terms_query,omitempty"`
}
func (m *MatchPhraseFieldBuilder) IsQuery() {}


func MatchPhraseField(
_query string,
) *MatchPhraseFieldBuilder {
    return &MatchPhraseFieldBuilder{
       Query: _query,
    }
}
func (m *MatchPhraseFieldBuilder) SetQuery(_query string) *MatchPhraseFieldBuilder {
    m.Query = _query
    return m
}
func (m *MatchPhraseFieldBuilder) SetAnalyzer(_analyzer string) *MatchPhraseFieldBuilder {
    m.Analyzer = &_analyzer
    return m
}
func (m *MatchPhraseFieldBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchPhraseFieldBuilder {
    m.ZeroTermsQuery = &_zeroTermsQuery
    return m
}