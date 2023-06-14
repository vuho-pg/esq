package esq

type MultiMatchInnerBuilder struct {
	Query_                           string   `json:"query"`
	Fields_                          []string `json:"fields"`
	Type_                            *string  `json:"type,omitempty"`
	Analyzer_                        *string  `json:"analyzer,omitempty"`
	Boost_                           *float64 `json:"boost,omitempty"`
	Operator_                        *string  `json:"operator,omitempty"`
	MinimumShouldMatch_              *string  `json:"minimum_should_match,omitempty"`
	Fuzziness_                       *string  `json:"fuzziness,omitempty"`
	Lenient_                         *bool    `json:"lenient,omitempty"`
	PrefixLength_                    *int     `json:"prefix_length,omitempty"`
	MaxExpansions_                   *int     `json:"max_expansions,omitempty"`
	FuzzyRewrite_                    *string  `json:"fuzzy_rewrite,omitempty"`
	ZeroTermsQuery_                  *string  `json:"zero_terms_query,omitempty"`
	CutOffFrequency_                 *float64 `json:"cut_off_frequency,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool    `json:"auto_generate_synonyms_phrase_query,omitempty"`
	FuzzyTranspositions_             *bool    `json:"fuzzy_transpositions,omitempty"`
}

type MultiMatchBuilder struct {
	fieldName  string
	MultiMatch map[string]*MultiMatchInnerBuilder `json:"multi_match"`
}

func (*MultiMatchBuilder) IsQuery() {}

func MultiMatch(
	_fieldName string,
	_query string,
	_fields []string,
) *MultiMatchBuilder {
	return &MultiMatchBuilder{
		fieldName: _fieldName,
		MultiMatch: map[string]*MultiMatchInnerBuilder{
			_fieldName: {
				Query_:  _query,
				Fields_: _fields,
			},
		},
	}
}

func (multiMatch *MultiMatchBuilder) Query(_query string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Query_ = _query
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Fields(_fields ...string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Fields_ = _fields
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Type(_type string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Type_ = &_type
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Analyzer(_analyzer string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Analyzer_ = &_analyzer
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Boost(_boost float64) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Boost_ = &_boost
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Operator(_operator string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Operator_ = &_operator
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) MinimumShouldMatch(_minimumShouldMatch string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Fuzziness(_fuzziness string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Fuzziness_ = &_fuzziness
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) Lenient(_lenient bool) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].Lenient_ = &_lenient
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) PrefixLength(_prefixLength int) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].PrefixLength_ = &_prefixLength
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) MaxExpansions(_maxExpansions int) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].MaxExpansions_ = &_maxExpansions
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) FuzzyRewrite(_fuzzyRewrite string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].FuzzyRewrite_ = &_fuzzyRewrite
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) ZeroTermsQuery(_zeroTermsQuery string) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) CutOffFrequency(_cutOffFrequency float64) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].CutOffFrequency_ = &_cutOffFrequency
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return multiMatch
}
func (multiMatch *MultiMatchBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *MultiMatchBuilder {
	multiMatch.MultiMatch[multiMatch.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return multiMatch
}
