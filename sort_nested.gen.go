package esq

import "encoding/json"

type SortNestedBuilder struct {
	Path_        string             `json:"path"`
	Nested_      *SortNestedBuilder `json:"nested,omitempty"`
	Filter_      *Query             `json:"filter,omitempty"`
	MaxChildren_ *int               `json:"max_children,omitempty"`
}

func SortNested(
	_path string,
) *SortNestedBuilder {
	return &SortNestedBuilder{
		Path_: _path,
	}
}

func (_sortNestedBuilder *SortNestedBuilder) Path(_path string) *SortNestedBuilder {
	_sortNestedBuilder.Path_ = _path
	return _sortNestedBuilder
}
func (_sortNestedBuilder *SortNestedBuilder) Nested(_nested *SortNestedBuilder) *SortNestedBuilder {
	_sortNestedBuilder.Nested_ = _nested
	return _sortNestedBuilder
}
func (_sortNestedBuilder *SortNestedBuilder) Filter(_filter Query) *SortNestedBuilder {
	_sortNestedBuilder.Filter_ = &_filter
	return _sortNestedBuilder
}
func (_sortNestedBuilder *SortNestedBuilder) MaxChildren(_maxChildren int) *SortNestedBuilder {
	_sortNestedBuilder.MaxChildren_ = &_maxChildren
	return _sortNestedBuilder
}

func (_sortNestedBuilder *SortNestedBuilder) JSON() ([]byte, error) {
	return json.Marshal(_sortNestedBuilder)
}

func (_sortNestedBuilder *SortNestedBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_sortNestedBuilder, "", "\t")
}
