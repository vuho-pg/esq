package esq

type HasChildInnerBuilder struct {
	Type_           string     `json:"type"`
	Query_          Query      `json:"query"`
	IgnoreUnmapped_ *bool      `json:"ignore_unmapped,omitempty"`
	MaxChildren_    *int       `json:"max_children,omitempty"`
	MinChildren_    *int       `json:"min_children,omitempty"`
	ScoreMode_      *ScoreMode `json:"score_mode,omitempty"`
}

type HasChildBuilder struct {
	HasChild *HasChildInnerBuilder `json:"has_child"`
}

func (*HasChildBuilder) IsQuery() {}

func HasChild(
	_type string,
	_query Query,
) *HasChildBuilder {
	return &HasChildBuilder{
		HasChild: &HasChildInnerBuilder{
			Type_:  _type,
			Query_: _query,
		},
	}
}

func (_hasChild *HasChildBuilder) Type(_type string) *HasChildBuilder {
	_hasChild.HasChild.Type_ = _type
	return _hasChild
}
func (_hasChild *HasChildBuilder) Query(_query Query) *HasChildBuilder {
	_hasChild.HasChild.Query_ = _query
	return _hasChild
}
func (_hasChild *HasChildBuilder) IgnoreUnmapped(_ignoreUnmapped bool) *HasChildBuilder {
	_hasChild.HasChild.IgnoreUnmapped_ = &_ignoreUnmapped
	return _hasChild
}
func (_hasChild *HasChildBuilder) MaxChildren(_maxChildren int) *HasChildBuilder {
	_hasChild.HasChild.MaxChildren_ = &_maxChildren
	return _hasChild
}
func (_hasChild *HasChildBuilder) MinChildren(_minChildren int) *HasChildBuilder {
	_hasChild.HasChild.MinChildren_ = &_minChildren
	return _hasChild
}
func (_hasChild *HasChildBuilder) ScoreMode(_scoreMode ScoreMode) *HasChildBuilder {
	_hasChild.HasChild.ScoreMode_ = &_scoreMode
	return _hasChild
}
