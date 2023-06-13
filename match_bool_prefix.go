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

func (matchBoolPrefix *MatchBoolPrefixBuilder) Query(_query string) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].Query_ = _query
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) Analyzer(_analyzer string) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].Analyzer_ = &_analyzer
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) Fuzziness(_fuzziness string) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].Fuzziness_ = &_fuzziness
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) MinimumShouldMatch(_minimumShouldMatch string) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) PrefixLength(_prefixLength int) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].PrefixLength_ = &_prefixLength
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) MaxExpansions(_maxExpansions int) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].MaxExpansions_ = &_maxExpansions
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return matchBoolPrefix
}
func (matchBoolPrefix *MatchBoolPrefixBuilder) FuzzyRewrite(_fuzzyRewrite string) *MatchBoolPrefixBuilder {
	matchBoolPrefix.MatchBoolPrefix[matchBoolPrefix.fieldName].FuzzyRewrite_ = &_fuzzyRewrite
	return matchBoolPrefix
}
