package esq

type TypeInnerBuilder struct {
	Value_ string `json:"value"`
}

type TypeBuilder struct {
	Type *TypeInnerBuilder `json:"type"`
}

func (*TypeBuilder) IsQuery() {}

func Type(
	_value string,
) *TypeBuilder {
	return &TypeBuilder{
		Type: &TypeInnerBuilder{
			Value_: _value,
		},
	}
}

func (_type *TypeBuilder) Value(_value string) *TypeBuilder {
	_type.Type.Value_ = _value
	return _type
}
