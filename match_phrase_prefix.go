package esq

type MatchPhrasePrefixInnerBuilder struct {
	Query_          string  `json:"query"`
	Analyzer_       *string `json:"analyzer,omitempty"`
	MaxExpansions_  *int    `json:"max_expansions,omitempty"`
	Slop_           *int    `json:"slop,omitempty"`
	ZeroTermsQuery_ *string `json:"zero_terms_query,omitempty"`
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

func (matchPhrasePrefix *MatchPhrasePrefixBuilder) Query(_query string) *MatchPhrasePrefixBuilder {
	matchPhrasePrefix.MatchPhrasePrefix[matchPhrasePrefix.fieldName].Query_ = _query
	return matchPhrasePrefix
}
func (matchPhrasePrefix *MatchPhrasePrefixBuilder) Analyzer(_analyzer string) *MatchPhrasePrefixBuilder {
	matchPhrasePrefix.MatchPhrasePrefix[matchPhrasePrefix.fieldName].Analyzer_ = &_analyzer
	return matchPhrasePrefix
}
func (matchPhrasePrefix *MatchPhrasePrefixBuilder) MaxExpansions(_maxExpansions int) *MatchPhrasePrefixBuilder {
	matchPhrasePrefix.MatchPhrasePrefix[matchPhrasePrefix.fieldName].MaxExpansions_ = &_maxExpansions
	return matchPhrasePrefix
}
func (matchPhrasePrefix *MatchPhrasePrefixBuilder) Slop(_slop int) *MatchPhrasePrefixBuilder {
	matchPhrasePrefix.MatchPhrasePrefix[matchPhrasePrefix.fieldName].Slop_ = &_slop
	return matchPhrasePrefix
}
func (matchPhrasePrefix *MatchPhrasePrefixBuilder) ZeroTermsQuery(_zeroTermsQuery string) *MatchPhrasePrefixBuilder {
	matchPhrasePrefix.MatchPhrasePrefix[matchPhrasePrefix.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return matchPhrasePrefix
}
