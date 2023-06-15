package esq

import "encoding/json"

type DocvalueFieldBuilder struct {
	Field_  string  `json:"field"`
	Format_ *string `json:"format,omitempty"`
}

func DocvalueField(
	_field string,
) *DocvalueFieldBuilder {
	return &DocvalueFieldBuilder{
		Field_: _field,
	}
}

func (_docvalueFieldBuilder *DocvalueFieldBuilder) Field(_field string) *DocvalueFieldBuilder {
	_docvalueFieldBuilder.Field_ = _field
	return _docvalueFieldBuilder
}
func (_docvalueFieldBuilder *DocvalueFieldBuilder) Format(_format string) *DocvalueFieldBuilder {
	_docvalueFieldBuilder.Format_ = &_format
	return _docvalueFieldBuilder
}

func (_docvalueFieldBuilder *DocvalueFieldBuilder) JSON() ([]byte, error) {
	return json.Marshal(_docvalueFieldBuilder)
}
