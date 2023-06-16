package esq

import "encoding/json"

type FieldBuilder struct {
	Field_  string  `json:"field"`
	Format_ *string `json:"format,omitempty"`
}

func Field(
	_field string,
) *FieldBuilder {
	return &FieldBuilder{
		Field_: _field,
	}
}

func (_fieldBuilder *FieldBuilder) Field(_field string) *FieldBuilder {
	_fieldBuilder.Field_ = _field
	return _fieldBuilder
}
func (_fieldBuilder *FieldBuilder) Format(_format string) *FieldBuilder {
	_fieldBuilder.Format_ = &_format
	return _fieldBuilder
}

func (_fieldBuilder *FieldBuilder) JSON() ([]byte, error) {
	return json.Marshal(_fieldBuilder)
}

func (_fieldBuilder *FieldBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_fieldBuilder, "", "\t")
}
