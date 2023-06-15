package esq

type TermInnerBuilder struct {
	Value_           string   `json:"value"`
	Boost_           *float64 `json:"boost,omitempty"`
	CaseInsensitive_ *bool    `json:"case_insensitive,omitempty"`
}

type TermBuilder struct {
	fieldName string
	Term      map[string]*TermInnerBuilder `json:"term"`
}

func (*TermBuilder) IsQuery() {}

func Term(
	_fieldName string,
	_value string,
) *TermBuilder {
	return &TermBuilder{
		fieldName: _fieldName,
		Term: map[string]*TermInnerBuilder{
			_fieldName: {
				Value_: _value,
			},
		},
	}
}

func (_term *TermBuilder) Value(_value string) *TermBuilder {
	_term.Term[_term.fieldName].Value_ = _value
	return _term
}
func (_term *TermBuilder) Boost(_boost float64) *TermBuilder {
	_term.Term[_term.fieldName].Boost_ = &_boost
	return _term
}
func (_term *TermBuilder) CaseInsensitive(_caseInsensitive bool) *TermBuilder {
	_term.Term[_term.fieldName].CaseInsensitive_ = &_caseInsensitive
	return _term
}
