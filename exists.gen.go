package esq

type ExistsInnerBuilder struct {
	Field_ string `json:"field"`
}

type ExistsBuilder struct {
	Exists *ExistsInnerBuilder `json:"exists"`
}

func (*ExistsBuilder) IsQuery() {}

func Exists(
	_field string,
) *ExistsBuilder {
	return &ExistsBuilder{
		Exists: &ExistsInnerBuilder{
			Field_: _field,
		},
	}
}

func (_exists *ExistsBuilder) Field(_field string) *ExistsBuilder {
	_exists.Exists.Field_ = _field
	return _exists
}
