package esq

type MatchInnerBuilder struct {
	Query_                           string          `json:"query"`
	Analyzer_                        *string         `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool           `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness_                       *string         `json:"fuzziness,omitempty"`
	PrefixLength_                    *int            `json:"prefix_length,omitempty"`
	MaxExpansions_                   *int            `json:"max_expansions,omitempty"`
	FuzzyTranspositions_             *bool           `json:"fuzzy_transpositions,omitempty"`
	FuzzyRewrite_                    *string         `json:"fuzzy_rewrite,omitempty"`
	Lenient_                         *bool           `json:"lenient,omitempty"`
	Operator_                        *Operator       `json:"operator,omitempty"`
	MinimumShouldMatch_              *string         `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery_                  *ZeroTermsQuery `json:"zero_terms_query,omitempty"`
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

func (_match *MatchBuilder) Query(_query string) *MatchBuilder {
	_match.Match[_match.fieldName].Query_ = _query
	return _match
}
func (_match *MatchBuilder) Analyzer(_analyzer string) *MatchBuilder {
	_match.Match[_match.fieldName].Analyzer_ = &_analyzer
	return _match
}
func (_match *MatchBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MatchBuilder {
	_match.Match[_match.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return _match
}
func (_match *MatchBuilder) Fuzziness(_fuzziness string) *MatchBuilder {
	_match.Match[_match.fieldName].Fuzziness_ = &_fuzziness
	return _match
}
func (_match *MatchBuilder) PrefixLength(_prefixLength int) *MatchBuilder {
	_match.Match[_match.fieldName].PrefixLength_ = &_prefixLength
	return _match
}
func (_match *MatchBuilder) MaxExpansions(_maxExpansions int) *MatchBuilder {
	_match.Match[_match.fieldName].MaxExpansions_ = &_maxExpansions
	return _match
}
func (_match *MatchBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *MatchBuilder {
	_match.Match[_match.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return _match
}
func (_match *MatchBuilder) FuzzyRewrite(_fuzzyRewrite string) *MatchBuilder {
	_match.Match[_match.fieldName].FuzzyRewrite_ = &_fuzzyRewrite
	return _match
}
func (_match *MatchBuilder) Lenient(_lenient bool) *MatchBuilder {
	_match.Match[_match.fieldName].Lenient_ = &_lenient
	return _match
}
func (_match *MatchBuilder) Operator(_operator Operator) *MatchBuilder {
	_match.Match[_match.fieldName].Operator_ = &_operator
	return _match
}
func (_match *MatchBuilder) MinimumShouldMatch(_minimumShouldMatch string) *MatchBuilder {
	_match.Match[_match.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return _match
}
func (_match *MatchBuilder) ZeroTermsQuery(_zeroTermsQuery ZeroTermsQuery) *MatchBuilder {
	_match.Match[_match.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return _match
}
