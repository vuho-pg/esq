package esq

type FuzzyInnerBuilder struct {
	Value_         string  `json:"value"`
	Fuzziness_     *string `json:"fuzziness,omitempty"`
	MaxExpansions_ *int    `json:"max_expansions,omitempty"`
	PreflixLength_ *int    `json:"preflix_length,omitempty"`
	Transposition_ *bool   `json:"transposition,omitempty"`
	Rewrite_       *string `json:"rewrite,omitempty"`
}

type FuzzyBuilder struct {
	fieldName string
	Fuzzy     map[string]*FuzzyInnerBuilder `json:"fuzzy"`
}

func (*FuzzyBuilder) IsQuery() {}

func Fuzzy(
	_fieldName string,
	_value string,
) *FuzzyBuilder {
	return &FuzzyBuilder{
		fieldName: _fieldName,
		Fuzzy: map[string]*FuzzyInnerBuilder{
			_fieldName: {
				Value_: _value,
			},
		},
	}
}

func (_fuzzy *FuzzyBuilder) Value(_value string) *FuzzyBuilder {
	_fuzzy.Fuzzy[_fuzzy.fieldName].Value_ = _value
	return _fuzzy
}
func (_fuzzy *FuzzyBuilder) Fuzziness(_fuzziness string) *FuzzyBuilder {
	_fuzzy.Fuzzy[_fuzzy.fieldName].Fuzziness_ = &_fuzziness
	return _fuzzy
}
func (_fuzzy *FuzzyBuilder) MaxExpansions(_maxExpansions int) *FuzzyBuilder {
	_fuzzy.Fuzzy[_fuzzy.fieldName].MaxExpansions_ = &_maxExpansions
	return _fuzzy
}
func (_fuzzy *FuzzyBuilder) PreflixLength(_preflixLength int) *FuzzyBuilder {
	_fuzzy.Fuzzy[_fuzzy.fieldName].PreflixLength_ = &_preflixLength
	return _fuzzy
}
func (_fuzzy *FuzzyBuilder) Transposition(_transposition bool) *FuzzyBuilder {
	_fuzzy.Fuzzy[_fuzzy.fieldName].Transposition_ = &_transposition
	return _fuzzy
}
func (_fuzzy *FuzzyBuilder) Rewrite(_rewrite string) *FuzzyBuilder {
	_fuzzy.Fuzzy[_fuzzy.fieldName].Rewrite_ = &_rewrite
	return _fuzzy
}
