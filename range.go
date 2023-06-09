package esq

type RangeBuilder struct {
	GT     *any    `json:"gt,omitempty"`
	GTE    *any    `json:"gte,omitempty"`
	LT     *any    `json:"lt,omitempty"`
	LTE    *any    `json:"lte,omitempty"`
	Format *string `json:"format,omitempty"`
}

func Range() *RangeBuilder {
	return &RangeBuilder{}
}
func (r *RangeBuilder) SetGT(_gT any) *RangeBuilder {
	r.GT = &_gT
	return r
}
func (r *RangeBuilder) SetGTE(_gTE any) *RangeBuilder {
	r.GTE = &_gTE
	return r
}
func (r *RangeBuilder) SetLT(_lT any) *RangeBuilder {
	r.LT = &_lT
	return r
}
func (r *RangeBuilder) SetLTE(_lTE any) *RangeBuilder {
	r.LTE = &_lTE
	return r
}
func (r *RangeBuilder) SetFormat(_format string) *RangeBuilder {
	r.Format = &_format
	return r
}
