package esq

type IDSInnerBuilder struct {
	Values_ string `json:"values"`
}

type IDSBuilder struct {
	IDS *IDSInnerBuilder `json:"ids"`
}

func (*IDSBuilder) IsQuery() {}

func IDS(
	_values string,
) *IDSBuilder {
	return &IDSBuilder{
		IDS: &IDSInnerBuilder{
			Values_: _values,
		},
	}
}

func (_iDS *IDSBuilder) Values(_values string) *IDSBuilder {
	_iDS.IDS.Values_ = _values
	return _iDS
}
