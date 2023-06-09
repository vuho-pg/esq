package esq

type MatchPhrasePrefixFieldBuilder struct {
	Query          string  `json:"query"`
	Analyzer       *string `json:"analyzer,omitempty"`
	MaxExpansions  *int    `json:"max_expansions,omitempty"`
	Slop           *int    `json:"slop,omitempty"`
	ZeroTermsQuery *string `json:"zero_terms_query,omitempty"`
}

func MatchPhrasePrefixField(
	_query string,
) *MatchPhrasePrefixFieldBuilder {
	return &MatchPhrasePrefixFieldBuilder{
		Query: _query,
	}
}
func (m *MatchPhrasePrefixFieldBuilder) SetQuery(_query string) *MatchPhrasePrefixFieldBuilder {
	m.Query = _query
	return m
}
func (m *MatchPhrasePrefixFieldBuilder) SetAnalyzer(_analyzer string) *MatchPhrasePrefixFieldBuilder {
	m.Analyzer = &_analyzer
	return m
}
func (m *MatchPhrasePrefixFieldBuilder) SetMaxExpansions(_maxExpansions int) *MatchPhrasePrefixFieldBuilder {
	m.MaxExpansions = &_maxExpansions
	return m
}
func (m *MatchPhrasePrefixFieldBuilder) SetSlop(_slop int) *MatchPhrasePrefixFieldBuilder {
	m.Slop = &_slop
	return m
}
func (m *MatchPhrasePrefixFieldBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *MatchPhrasePrefixFieldBuilder {
	m.ZeroTermsQuery = &_zeroTermsQuery
	return m
}
