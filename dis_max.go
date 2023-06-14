package esq

type DisMaxInnerBuilder struct {
	Queries_    []Query  `json:"queries,omitempty"`
	TieBreaker_ *float64 `json:"tie_breaker,omitempty"`
}

type DisMaxBuilder struct {
	DisMax *DisMaxInnerBuilder `json:"dis_max"`
}

func (*DisMaxBuilder) IsQuery() {}

func DisMax() *DisMaxBuilder {
	return &DisMaxBuilder{
		DisMax: &DisMaxInnerBuilder{},
	}
}

func (_disMax *DisMaxBuilder) Queries(_queries ...Query) *DisMaxBuilder {
	_disMax.DisMax.Queries_ = _queries
	return _disMax
}
func (_disMax *DisMaxBuilder) TieBreaker(_tieBreaker float64) *DisMaxBuilder {
	_disMax.DisMax.TieBreaker_ = &_tieBreaker
	return _disMax
}
