package esq

type MatchPhraseInnerBuilder struct {
	Query_          any             `json:"query"`
	Analyzer_       *string         `json:"analyzer,omitempty"`
	ZeroTermsQuery_ *ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

type MatchPhraseBuilder struct {
	fieldName   string
	MatchPhrase map[string]*MatchPhraseInnerBuilder `json:"match_phrase"`
}

func (*MatchPhraseBuilder) IsQuery() {}

func MatchPhrase(
	_fieldName string,
	_query any,
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

func (_matchPhrase *MatchPhraseBuilder) Query(_query any) *MatchPhraseBuilder {
	_matchPhrase.MatchPhrase[_matchPhrase.fieldName].Query_ = _query
	return _matchPhrase
}
func (_matchPhrase *MatchPhraseBuilder) Analyzer(_analyzer string) *MatchPhraseBuilder {
	_matchPhrase.MatchPhrase[_matchPhrase.fieldName].Analyzer_ = &_analyzer
	return _matchPhrase
}
func (_matchPhrase *MatchPhraseBuilder) ZeroTermsQuery(_zeroTermsQuery ZeroTermsQuery) *MatchPhraseBuilder {
	_matchPhrase.MatchPhrase[_matchPhrase.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return _matchPhrase
}
