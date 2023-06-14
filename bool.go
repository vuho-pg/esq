package esq

type BoolInnerBuilder struct {
	Must_               []Query  `json:"must,omitempty"`
	MustNot_            []Query  `json:"must_not,omitempty"`
	Should_             []Query  `json:"should,omitempty"`
	Filter_             []Query  `json:"filter,omitempty"`
	MinimumShouldMatch_ *string  `json:"minimum_should_match,omitempty"`
	Boost_              *float64 `json:"boost,omitempty"`
}

type BoolBuilder struct {
	Bool *BoolInnerBuilder `json:"bool"`
}

func (*BoolBuilder) IsQuery() {}

func Bool() *BoolBuilder {
	return &BoolBuilder{
		Bool: &BoolInnerBuilder{},
	}
}

func (bool *BoolBuilder) Must(_must ...Query) *BoolBuilder {
	bool.Bool.Must_ = _must
	return bool
}
func (bool *BoolBuilder) MustNot(_mustNot ...Query) *BoolBuilder {
	bool.Bool.MustNot_ = _mustNot
	return bool
}
func (bool *BoolBuilder) Should(_should ...Query) *BoolBuilder {
	bool.Bool.Should_ = _should
	return bool
}
func (bool *BoolBuilder) Filter(_filter ...Query) *BoolBuilder {
	bool.Bool.Filter_ = _filter
	return bool
}
func (bool *BoolBuilder) MinimumShouldMatch(_minimumShouldMatch string) *BoolBuilder {
	bool.Bool.MinimumShouldMatch_ = &_minimumShouldMatch
	return bool
}
func (bool *BoolBuilder) Boost(_boost float64) *BoolBuilder {
	bool.Bool.Boost_ = &_boost
	return bool
}
