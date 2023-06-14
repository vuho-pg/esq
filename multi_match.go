package esq

type MultiMatchInnerBuilder struct {
	Query_                           string          `json:"query"`
	Fields_                          []string        `json:"fields"`
	Type_                            *string         `json:"type,omitempty"`
	Analyzer_                        *string         `json:"analyzer,omitempty"`
	Boost_                           *float64        `json:"boost,omitempty"`
	Operator_                        *Operator       `json:"operator,omitempty"`
	MinimumShouldMatch_              *string         `json:"minimum_should_match,omitempty"`
	Fuzziness_                       *string         `json:"fuzziness,omitempty"`
	Lenient_                         *bool           `json:"lenient,omitempty"`
	PrefixLength_                    *int            `json:"prefix_length,omitempty"`
	MaxExpansions_                   *int            `json:"max_expansions,omitempty"`
	FuzzyRewrite_                    *string         `json:"fuzzy_rewrite,omitempty"`
	ZeroTermsQuery_                  *ZeroTermsQuery `json:"zero_terms_query,omitempty"`
	CutOffFrequency_                 *float64        `json:"cut_off_frequency,omitempty"`
	AutoGenerateSynonymsPhraseQuery_ *bool           `json:"auto_generate_synonyms_phrase_query,omitempty"`
	FuzzyTranspositions_             *bool           `json:"fuzzy_transpositions,omitempty"`
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

func (_multiMatch *MultiMatchBuilder) Query(_query string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Query_ = _query
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Fields(_fields ...string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Fields_ = _fields
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Type(_type string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Type_ = &_type
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Analyzer(_analyzer string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Analyzer_ = &_analyzer
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Boost(_boost float64) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Boost_ = &_boost
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Operator(_operator Operator) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Operator_ = &_operator
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) MinimumShouldMatch(_minimumShouldMatch string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Fuzziness(_fuzziness string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Fuzziness_ = &_fuzziness
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) Lenient(_lenient bool) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].Lenient_ = &_lenient
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) PrefixLength(_prefixLength int) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].PrefixLength_ = &_prefixLength
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) MaxExpansions(_maxExpansions int) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].MaxExpansions_ = &_maxExpansions
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) FuzzyRewrite(_fuzzyRewrite string) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].FuzzyRewrite_ = &_fuzzyRewrite
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) ZeroTermsQuery(_zeroTermsQuery ZeroTermsQuery) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) CutOffFrequency(_cutOffFrequency float64) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].CutOffFrequency_ = &_cutOffFrequency
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return _multiMatch
}
func (_multiMatch *MultiMatchBuilder) FuzzyTranspositions(_fuzzyTranspositions bool) *MultiMatchBuilder {
	_multiMatch.MultiMatch[_multiMatch.fieldName].FuzzyTranspositions_ = &_fuzzyTranspositions
	return _multiMatch
}
