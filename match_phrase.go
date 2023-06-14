package esq

type MatchPhraseInnerBuilder struct {
	Query_          string  `json:"query"`
	Analyzer_       *string `json:"analyzer,omitempty"`
	ZeroTermsQuery_ *string `json:"zero_terms_query,omitempty"`
}

type MatchPhraseBuilder struct {
	fieldName   string
	MatchPhrase map[string]*MatchPhraseInnerBuilder `json:"match_phrase"`
}

func (*MatchPhraseBuilder) IsQuery() {}

func MatchPhrase(
	_fieldName string,
	_query string,
) *MatchPhraseBuilder {
	return &MatchPhraseBuilder{
		fieldName: _fieldName,
		MatchPhrase: map[string]*MatchPhraseInnerBuilder{
			_fieldName: {
				Query_: _query,
			},
		},
	}
}

func (matchPhrase *MatchPhraseBuilder) Query(_query string) *MatchPhraseBuilder {
	matchPhrase.MatchPhrase[matchPhrase.fieldName].Query_ = _query
	return matchPhrase
}
func (matchPhrase *MatchPhraseBuilder) Analyzer(_analyzer string) *MatchPhraseBuilder {
	matchPhrase.MatchPhrase[matchPhrase.fieldName].Analyzer_ = &_analyzer
	return matchPhrase
}
func (matchPhrase *MatchPhraseBuilder) ZeroTermsQuery(_zeroTermsQuery string) *MatchPhraseBuilder {
	matchPhrase.MatchPhrase[matchPhrase.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return matchPhrase
}
