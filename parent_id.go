package esq

type ParentIDInnerBuilder struct {
	Type_           string `json:"type"`
	ID_             string `json:"id"`
	IgnoreUnmapped_ *bool  `json:"ignore_unmapped,omitempty"`
}

type ParentIDBuilder struct {
	ParentID *ParentIDInnerBuilder `json:"parent_id"`
}

func (*ParentIDBuilder) IsQuery() {}

func ParentID(
	_type string,
	_iD string,
) *ParentIDBuilder {
	return &ParentIDBuilder{
		ParentID: &ParentIDInnerBuilder{
			Type_: _type,
			ID_:   _iD,
		},
	}
}

func (_parentID *ParentIDBuilder) Type(_type string) *ParentIDBuilder {
	_parentID.ParentID.Type_ = _type
	return _parentID
}
func (_parentID *ParentIDBuilder) ID(_iD string) *ParentIDBuilder {
	_parentID.ParentID.ID_ = _iD
	return _parentID
}
func (_parentID *ParentIDBuilder) IgnoreUnmapped(_ignoreUnmapped bool) *ParentIDBuilder {
	_parentID.ParentID.IgnoreUnmapped_ = &_ignoreUnmapped
	return _parentID
}
