package esq

type BoostingInnerBuilder struct {
	Positive_      Query   `json:"positive"`
	Negative_      Query   `json:"negative"`
	NegativeBoost_ float64 `json:"negative_boost"`
}

type BoostingBuilder struct {
	Boosting *BoostingInnerBuilder `json:"boosting"`
}

func (*BoostingBuilder) IsQuery() {}

func Boosting(
	_positive Query,
	_negative Query,
	_negativeBoost float64,
) *BoostingBuilder {
	return &BoostingBuilder{
		Boosting: &BoostingInnerBuilder{
			Positive_:      _positive,
			Negative_:      _negative,
			NegativeBoost_: _negativeBoost,
		},
	}
}

func (_boosting *BoostingBuilder) Positive(_positive Query) *BoostingBuilder {
	_boosting.Boosting.Positive_ = _positive
	return _boosting
}
func (_boosting *BoostingBuilder) Negative(_negative Query) *BoostingBuilder {
	_boosting.Boosting.Negative_ = _negative
	return _boosting
}
func (_boosting *BoostingBuilder) NegativeBoost(_negativeBoost float64) *BoostingBuilder {
	_boosting.Boosting.NegativeBoost_ = _negativeBoost
	return _boosting
}
