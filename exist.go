package esq

type ExistInnerBuilder struct {
	Field_ string `json:"field"`
}

type ExistBuilder struct {
	Exist *ExistInnerBuilder `json:"exist"`
}

func (*ExistBuilder) IsQuery() {}

func Exist(
	_field string,
) *ExistBuilder {
	return &ExistBuilder{
		Exist: &ExistInnerBuilder{
			Field_: _field,
		},
	}
}

func (_exist *ExistBuilder) Field(_field string) *ExistBuilder {
	_exist.Exist.Field_ = _field
	return _exist
}
