package esq

import "encoding/json"

type SortValueBuilder struct {
	Order_       *SortOrder         `json:"order,omitempty"`
	Format_      *string            `json:"format,omitempty"`
	Mode_        *SortMode          `json:"mode,omitempty"`
	NumericType_ *NumericType       `json:"numeric_type,omitempty"`
	Missing_     *SortMissing       `json:"missing,omitempty"`
	Nested_      *SortNestedBuilder `json:"nested,omitempty"`
}

func SortValue() *SortValueBuilder {
	return &SortValueBuilder{}
}

func (_sortValueBuilder *SortValueBuilder) Order(_order SortOrder) *SortValueBuilder {
	_sortValueBuilder.Order_ = &_order
	return _sortValueBuilder
}
func (_sortValueBuilder *SortValueBuilder) Format(_format string) *SortValueBuilder {
	_sortValueBuilder.Format_ = &_format
	return _sortValueBuilder
}
func (_sortValueBuilder *SortValueBuilder) Mode(_mode SortMode) *SortValueBuilder {
	_sortValueBuilder.Mode_ = &_mode
	return _sortValueBuilder
}
func (_sortValueBuilder *SortValueBuilder) NumericType(_numericType NumericType) *SortValueBuilder {
	_sortValueBuilder.NumericType_ = &_numericType
	return _sortValueBuilder
}
func (_sortValueBuilder *SortValueBuilder) Missing(_missing SortMissing) *SortValueBuilder {
	_sortValueBuilder.Missing_ = &_missing
	return _sortValueBuilder
}
func (_sortValueBuilder *SortValueBuilder) Nested(_nested *SortNestedBuilder) *SortValueBuilder {
	_sortValueBuilder.Nested_ = _nested
	return _sortValueBuilder
}

func (_sortValueBuilder *SortValueBuilder) JSON() ([]byte, error) {
	return json.Marshal(_sortValueBuilder)
}

func (_sortValueBuilder *SortValueBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_sortValueBuilder, "", "\t")
}
