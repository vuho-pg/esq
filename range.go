package esq

type RangeInnerBuilder struct {
	GT_       *any     `json:"gt,omitempty"`
	GTE_      *any     `json:"gte,omitempty"`
	LT_       *any     `json:"lt,omitempty"`
	LTE_      *any     `json:"lte,omitempty"`
	Format_   *string  `json:"format,omitempty"`
	Relation_ *string  `json:"relation,omitempty"`
	Timezone_ *string  `json:"timezone,omitempty"`
	Boost_    *float64 `json:"boost,omitempty"`
}

type RangeBuilder struct {
	fieldName string
	Range     map[string]*RangeInnerBuilder `json:"range"`
}

func (*RangeBuilder) IsQuery() {}

func Range(
	_fieldName string,
) *RangeBuilder {
	return &RangeBuilder{
		fieldName: _fieldName,
		Range: map[string]*RangeInnerBuilder{
			_fieldName: {},
		},
	}
}

func (_range *RangeBuilder) GT(_gT any) *RangeBuilder {
	_range.Range[_range.fieldName].GT_ = &_gT
	return _range
}
func (_range *RangeBuilder) GTE(_gTE any) *RangeBuilder {
	_range.Range[_range.fieldName].GTE_ = &_gTE
	return _range
}
func (_range *RangeBuilder) LT(_lT any) *RangeBuilder {
	_range.Range[_range.fieldName].LT_ = &_lT
	return _range
}
func (_range *RangeBuilder) LTE(_lTE any) *RangeBuilder {
	_range.Range[_range.fieldName].LTE_ = &_lTE
	return _range
}
func (_range *RangeBuilder) Format(_format string) *RangeBuilder {
	_range.Range[_range.fieldName].Format_ = &_format
	return _range
}
func (_range *RangeBuilder) Relation(_relation string) *RangeBuilder {
	_range.Range[_range.fieldName].Relation_ = &_relation
	return _range
}
func (_range *RangeBuilder) Timezone(_timezone string) *RangeBuilder {
	_range.Range[_range.fieldName].Timezone_ = &_timezone
	return _range
}
func (_range *RangeBuilder) Boost(_boost float64) *RangeBuilder {
	_range.Range[_range.fieldName].Boost_ = &_boost
	return _range
}
