package esq

type WildcardInnerBuilder struct {
	Value_           string   `json:"value"`
	Boost_           *float64 `json:"boost,omitempty"`
	CaseInsensitive_ *bool    `json:"case_insensitive,omitempty"`
	Rewrite_         *string  `json:"rewrite,omitempty"`
}

type WildcardBuilder struct {
	fieldName string
	Wildcard  map[string]*WildcardInnerBuilder `json:"wildcard"`
}

func (*WildcardBuilder) IsQuery() {}

func Wildcard(
	_fieldName string,
	_value string,
) *WildcardBuilder {
	return &WildcardBuilder{
		fieldName: _fieldName,
		Wildcard: map[string]*WildcardInnerBuilder{
			_fieldName: {
				Value_: _value,
			},
		},
	}
}

func (_wildcard *WildcardBuilder) Value(_value string) *WildcardBuilder {
	_wildcard.Wildcard[_wildcard.fieldName].Value_ = _value
	return _wildcard
}
func (_wildcard *WildcardBuilder) Boost(_boost float64) *WildcardBuilder {
	_wildcard.Wildcard[_wildcard.fieldName].Boost_ = &_boost
	return _wildcard
}
func (_wildcard *WildcardBuilder) CaseInsensitive(_caseInsensitive bool) *WildcardBuilder {
	_wildcard.Wildcard[_wildcard.fieldName].CaseInsensitive_ = &_caseInsensitive
	return _wildcard
}
func (_wildcard *WildcardBuilder) Rewrite(_rewrite string) *WildcardBuilder {
	_wildcard.Wildcard[_wildcard.fieldName].Rewrite_ = &_rewrite
	return _wildcard
}
