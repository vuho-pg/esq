package esq

type HasParentInnerBuilder struct {
	ParentType_     string `json:"parent_type"`
	Query_          Query  `json:"query"`
	Score_          *bool  `json:"score,omitempty"`
	IgnoreUnmapped_ *bool  `json:"ignore_unmapped,omitempty"`
}

type HasParentBuilder struct {
	HasParent *HasParentInnerBuilder `json:"has_parent"`
}

func (*HasParentBuilder) IsQuery() {}

func HasParent(
	_parentType string,
	_query Query,
) *HasParentBuilder {
	return &HasParentBuilder{
		HasParent: &HasParentInnerBuilder{
			ParentType_: _parentType,
			Query_:      _query,
		},
	}
}

func (_hasParent *HasParentBuilder) ParentType(_parentType string) *HasParentBuilder {
	_hasParent.HasParent.ParentType_ = _parentType
	return _hasParent
}
func (_hasParent *HasParentBuilder) Query(_query Query) *HasParentBuilder {
	_hasParent.HasParent.Query_ = _query
	return _hasParent
}
func (_hasParent *HasParentBuilder) Score(_score bool) *HasParentBuilder {
	_hasParent.HasParent.Score_ = &_score
	return _hasParent
}
func (_hasParent *HasParentBuilder) IgnoreUnmapped(_ignoreUnmapped bool) *HasParentBuilder {
	_hasParent.HasParent.IgnoreUnmapped_ = &_ignoreUnmapped
	return _hasParent
}
