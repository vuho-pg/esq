package esq

type MatchInnerBuilder struct {
	Query_                           string  `json:"query"`
	Analyzer_                        *string `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool   `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness_                       *string `json:"fuzziness,omitempty"`
	PrefixLength_                    *int    `json:"prefix_length,omitempty"`
	MaxExpansions_                   *int    `json:"max_expansions,omitempty"`
	FuzzyTranspositions_             *bool   `json:"fuzzy_transpositions,omitempty"`
	FuzzyRewrite_                    *string `json:"fuzzy_rewrite,omitempty"`
	Lenient_                         *bool   `json:"lenient,omitempty"`
	Operator_                        *string `json:"operator,omitempty"`
	MinimumShouldMatch_              *string `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery_                  *string `json:"zero_terms_query,omitempty"`
}

type MatchBuilder struct {
	fieldName string
	Match     map[string]*MatchInnerBuilder `json:"match"`
}

func (*MatchBuilder) IsQuery() {}

func Match(
	_fieldName string,
	_query string,
) *MatchBuilder {
	return &MatchBuilder{
		fieldName: _fieldName,
		Match: map[string]*MatchInnerBuilder{
			_fieldName: {
				Query_: _query,
			},
		},
	}
}

func (match *MatchBuilder) Query(_query string) *MatchBuilder {
	match.Match[match.fieldName].Query_ = _query
	return match
}
func (match *MatchBuilder) Analyzer(_analyzer string) *MatchBuilder {
	match.Match[match.fieldName].Analyzer_ = &_analyzer
	return match
}
func (match *MatchBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MatchBuilder {
	match.Match[match.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return match
}
func (match *MatchBuilder) Fuzziness(_fuzziness string) *MatchBuilder {
	match.Match[match.fieldName].Fuzziness_ = &_fuzziness
	return match
}
func (match *MatchBuilder) PrefixLength(_prefixLength int) *MatchBuilder {
	match.Match[match.fieldName].PrefixLength_ = &_prefixLength
	return match
}
func (match *MatchBuilder) MaxExpansions(_maxExpansions int) *MatchBuilder {
	match.Match[match.fieldName].MaxExpansions_ = &_maxExpansions
	return match
}
func (match *MatchBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *MatchBuilder {
	match.Match[match.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return match
}
func (match *MatchBuilder) FuzzyRewrite(_fuzzyRewrite string) *MatchBuilder {
	match.Match[match.fieldName].FuzzyRewrite_ = &_fuzzyRewrite
	return match
}
func (match *MatchBuilder) Lenient(_lenient bool) *MatchBuilder {
	match.Match[match.fieldName].Lenient_ = &_lenient
	return match
}
func (match *MatchBuilder) Operator(_operator string) *MatchBuilder {
	match.Match[match.fieldName].Operator_ = &_operator
	return match
}
func (match *MatchBuilder) MinimumShouldMatch(_minimumShouldMatch string) *MatchBuilder {
	match.Match[match.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return match
}
func (match *MatchBuilder) ZeroTermsQuery(_zeroTermsQuery string) *MatchBuilder {
	match.Match[match.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return match
}
