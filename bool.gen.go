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

func (_bool *BoolBuilder) Must(_must ...Query) *BoolBuilder {
	_bool.Bool.Must_ = _must
	return _bool
}
func (_bool *BoolBuilder) MustNot(_mustNot ...Query) *BoolBuilder {
	_bool.Bool.MustNot_ = _mustNot
	return _bool
}
func (_bool *BoolBuilder) Should(_should ...Query) *BoolBuilder {
	_bool.Bool.Should_ = _should
	return _bool
}
func (_bool *BoolBuilder) Filter(_filter ...Query) *BoolBuilder {
	_bool.Bool.Filter_ = _filter
	return _bool
}
func (_bool *BoolBuilder) MinimumShouldMatch(_minimumShouldMatch string) *BoolBuilder {
	_bool.Bool.MinimumShouldMatch_ = &_minimumShouldMatch
	return _bool
}
func (_bool *BoolBuilder) Boost(_boost float64) *BoolBuilder {
	_bool.Bool.Boost_ = &_boost
	return _bool
}
