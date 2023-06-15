package esq

type MatchPhrasePrefixInnerBuilder struct {
	Query_          string          `json:"query"`
	Analyzer_       *string         `json:"analyzer,omitempty"`
	MaxExpansions_  *int            `json:"max_expansions,omitempty"`
	Slop_           *int            `json:"slop,omitempty"`
	ZeroTermsQuery_ *ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

type MatchPhrasePrefixBuilder struct {
	fieldName         string
	MatchPhrasePrefix map[string]*MatchPhrasePrefixInnerBuilder `json:"match_phrase_prefix"`
}

func (*MatchPhrasePrefixBuilder) IsQuery() {}

func MatchPhrasePrefix(
	_fieldName string,
	_query string,
) *MatchPhrasePrefixBuilder {
	return &MatchPhrasePrefixBuilder{
		fieldName: _fieldName,
		MatchPhrasePrefix: map[string]*MatchPhrasePrefixInnerBuilder{
			_fieldName: {
				Query_: _query,
			},
		},
	}
}

func (_matchPhrasePrefix *MatchPhrasePrefixBuilder) Query(_query string) *MatchPhrasePrefixBuilder {
	_matchPhrasePrefix.MatchPhrasePrefix[_matchPhrasePrefix.fieldName].Query_ = _query
	return _matchPhrasePrefix
}
func (_matchPhrasePrefix *MatchPhrasePrefixBuilder) Analyzer(_analyzer string) *MatchPhrasePrefixBuilder {
	_matchPhrasePrefix.MatchPhrasePrefix[_matchPhrasePrefix.fieldName].Analyzer_ = &_analyzer
	return _matchPhrasePrefix
}
func (_matchPhrasePrefix *MatchPhrasePrefixBuilder) MaxExpansions(_maxExpansions int) *MatchPhrasePrefixBuilder {
	_matchPhrasePrefix.MatchPhrasePrefix[_matchPhrasePrefix.fieldName].MaxExpansions_ = &_maxExpansions
	return _matchPhrasePrefix
}
func (_matchPhrasePrefix *MatchPhrasePrefixBuilder) Slop(_slop int) *MatchPhrasePrefixBuilder {
	_matchPhrasePrefix.MatchPhrasePrefix[_matchPhrasePrefix.fieldName].Slop_ = &_slop
	return _matchPhrasePrefix
}
func (_matchPhrasePrefix *MatchPhrasePrefixBuilder) ZeroTermsQuery(_zeroTermsQuery ZeroTermsQuery) *MatchPhrasePrefixBuilder {
	_matchPhrasePrefix.MatchPhrasePrefix[_matchPhrasePrefix.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return _matchPhrasePrefix
}
