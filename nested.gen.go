package esq

type NestedInnerBuilder struct {
	Path_           string     `json:"path"`
	Query_          Query      `json:"query"`
	ScoreMode_      *ScoreMode `json:"score_mode,omitempty"`
	IgnoreUnmapped_ *bool      `json:"ignore_unmapped,omitempty"`
}

type NestedBuilder struct {
	Nested *NestedInnerBuilder `json:"nested"`
}

func (*NestedBuilder) IsQuery() {}

func Nested(
	_path string,
	_query Query,
) *NestedBuilder {
	return &NestedBuilder{
		Nested: &NestedInnerBuilder{
			Path_:  _path,
			Query_: _query,
		},
	}
}

func (_nested *NestedBuilder) Path(_path string) *NestedBuilder {
	_nested.Nested.Path_ = _path
	return _nested
}
func (_nested *NestedBuilder) Query(_query Query) *NestedBuilder {
	_nested.Nested.Query_ = _query
	return _nested
}
func (_nested *NestedBuilder) ScoreMode(_scoreMode ScoreMode) *NestedBuilder {
	_nested.Nested.ScoreMode_ = &_scoreMode
	return _nested
}
func (_nested *NestedBuilder) IgnoreUnmapped(_ignoreUnmapped bool) *NestedBuilder {
	_nested.Nested.IgnoreUnmapped_ = &_ignoreUnmapped
	return _nested
}
