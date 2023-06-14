package esq

type MatchBoolPrefixInnerBuilder struct {
	Query_               string  `json:"query"`
	Analyzer_            *string `json:"analyzer,omitempty"`
	Fuzziness_           *string `json:"fuzziness,omitempty"`
	MinimumShouldMatch_  *string `json:"minimum_should_match,omitempty"`
	PrefixLength_        *int    `json:"prefix_length,omitempty"`
	MaxExpansions_       *int    `json:"max_expansions,omitempty"`
	FuzzyTranspositions_ *bool   `json:"fuzzy_transpositions,omitempty"`
	FuzzyRewrite_        *string `json:"fuzzy_rewrite,omitempty"`
}

type MatchBoolPrefixBuilder struct {
	fieldName       string
	MatchBoolPrefix map[string]*MatchBoolPrefixInnerBuilder `json:"match_bool_prefix"`
}

func (*MatchBoolPrefixBuilder) IsQuery() {}

func MatchBoolPrefix(
	_fieldName string,
	_query string,
) *MatchBoolPrefixBuilder {
	return &MatchBoolPrefixBuilder{
		fieldName: _fieldName,
		MatchBoolPrefix: map[string]*MatchBoolPrefixInnerBuilder{
			_fieldName: {
				Query_: _query,
			},
		},
	}
}

func (_matchBoolPrefix *MatchBoolPrefixBuilder) Query(_query string) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].Query_ = _query
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) Analyzer(_analyzer string) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].Analyzer_ = &_analyzer
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) Fuzziness(_fuzziness string) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].Fuzziness_ = &_fuzziness
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) MinimumShouldMatch(_minimumShouldMatch string) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) PrefixLength(_prefixLength int) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].PrefixLength_ = &_prefixLength
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) MaxExpansions(_maxExpansions int) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].MaxExpansions_ = &_maxExpansions
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return _matchBoolPrefix
}
func (_matchBoolPrefix *MatchBoolPrefixBuilder) FuzzyRewrite(_fuzzyRewrite string) *MatchBoolPrefixBuilder {
	_matchBoolPrefix.MatchBoolPrefix[_matchBoolPrefix.fieldName].FuzzyRewrite_ = &_fuzzyRewrite
	return _matchBoolPrefix
}
