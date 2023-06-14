package esq

type NestedInnerBuilder struct {
	Path_           string  `json:"path"`
	Query_          Query   `json:"query"`
	ScoreMode_      *string `json:"score_mode,omitempty"`
	IgnoreUnmapped_ *bool   `json:"ignore_unmapped,omitempty"`
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

func (nested *NestedBuilder) Path(_path string) *NestedBuilder {
	nested.Nested.Path_ = _path
	return nested
}
func (nested *NestedBuilder) Query(_query Query) *NestedBuilder {
	nested.Nested.Query_ = _query
	return nested
}
func (nested *NestedBuilder) ScoreMode(_scoreMode string) *NestedBuilder {
	nested.Nested.ScoreMode_ = &_scoreMode
	return nested
}
func (nested *NestedBuilder) IgnoreUnmapped(_ignoreUnmapped bool) *NestedBuilder {
	nested.Nested.IgnoreUnmapped_ = &_ignoreUnmapped
	return nested
}
