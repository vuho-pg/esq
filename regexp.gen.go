package esq

type RegexpInnerBuilder struct {
	Value_                   string  `json:"value"`
	Flags_                   *string `json:"flags,omitempty"`
	CaseInsensitive_         *bool   `json:"case_insensitive,omitempty"`
	MaxDetermiinizzedStates_ *int    `json:"max_determiinizzed_states,omitempty"`
	Rewrite_                 *string `json:"rewrite,omitempty"`
}

type RegexpBuilder struct {
	fieldName string
	Regexp    map[string]*RegexpInnerBuilder `json:"regexp"`
}

func (*RegexpBuilder) IsQuery() {}

func Regexp(
	_fieldName string,
	_value string,
) *RegexpBuilder {
	return &RegexpBuilder{
		fieldName: _fieldName,
		Regexp: map[string]*RegexpInnerBuilder{
			_fieldName: {
				Value_: _value,
			},
		},
	}
}

func (_regexp *RegexpBuilder) Value(_value string) *RegexpBuilder {
	_regexp.Regexp[_regexp.fieldName].Value_ = _value
	return _regexp
}
func (_regexp *RegexpBuilder) Flags(_flags string) *RegexpBuilder {
	_regexp.Regexp[_regexp.fieldName].Flags_ = &_flags
	return _regexp
}
func (_regexp *RegexpBuilder) CaseInsensitive(_caseInsensitive bool) *RegexpBuilder {
	_regexp.Regexp[_regexp.fieldName].CaseInsensitive_ = &_caseInsensitive
	return _regexp
}
func (_regexp *RegexpBuilder) MaxDetermiinizzedStates(_maxDetermiinizzedStates int) *RegexpBuilder {
	_regexp.Regexp[_regexp.fieldName].MaxDetermiinizzedStates_ = &_maxDetermiinizzedStates
	return _regexp
}
func (_regexp *RegexpBuilder) Rewrite(_rewrite string) *RegexpBuilder {
	_regexp.Regexp[_regexp.fieldName].Rewrite_ = &_rewrite
	return _regexp
}
