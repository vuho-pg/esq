package esq

type PrefixInnerBuilder struct {
	Value_           string  `json:"value"`
	Rewrite_         *string `json:"rewrite,omitempty"`
	CaseInsensitive_ *bool   `json:"case_insensitive,omitempty"`
}

type PrefixBuilder struct {
	fieldName string
	Prefix    map[string]*PrefixInnerBuilder `json:"prefix"`
}

func (*PrefixBuilder) IsQuery() {}

func Prefix(
	_fieldName string,
	_value string,
) *PrefixBuilder {
	return &PrefixBuilder{
		fieldName: _fieldName,
		Prefix: map[string]*PrefixInnerBuilder{
			_fieldName: {
				Value_: _value,
			},
		},
	}
}

func (_prefix *PrefixBuilder) Value(_value string) *PrefixBuilder {
	_prefix.Prefix[_prefix.fieldName].Value_ = _value
	return _prefix
}
func (_prefix *PrefixBuilder) Rewrite(_rewrite string) *PrefixBuilder {
	_prefix.Prefix[_prefix.fieldName].Rewrite_ = &_rewrite
	return _prefix
}
func (_prefix *PrefixBuilder) CaseInsensitive(_caseInsensitive bool) *PrefixBuilder {
	_prefix.Prefix[_prefix.fieldName].CaseInsensitive_ = &_caseInsensitive
	return _prefix
}
