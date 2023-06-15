package esq

type ConstantScoreInnerBuilder struct {
	Filter_ Query    `json:"filter"`
	Boost_  *float64 `json:"boost,omitempty"`
}

type ConstantScoreBuilder struct {
	ConstantScore *ConstantScoreInnerBuilder `json:"constant_score"`
}

func (*ConstantScoreBuilder) IsQuery() {}

func ConstantScore(
	_filter Query,
) *ConstantScoreBuilder {
	return &ConstantScoreBuilder{
		ConstantScore: &ConstantScoreInnerBuilder{
			Filter_: _filter,
		},
	}
}

func (_constantScore *ConstantScoreBuilder) Filter(_filter Query) *ConstantScoreBuilder {
	_constantScore.ConstantScore.Filter_ = _filter
	return _constantScore
}
func (_constantScore *ConstantScoreBuilder) Boost(_boost float64) *ConstantScoreBuilder {
	_constantScore.ConstantScore.Boost_ = &_boost
	return _constantScore
}
