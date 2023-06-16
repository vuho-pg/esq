package esq

import "encoding/json"

type CollapseBuilder struct {
	Field_                      string              `json:"field"`
	InnerHits_                  []*InnerHitsBuilder `json:"inner_hits,omitempty"`
	MaxConcurrentGroupSearches_ *int                `json:"max_concurrent_group_searches,omitempty"`
}

func Collapse(
	_field string,
) *CollapseBuilder {
	return &CollapseBuilder{
		Field_: _field,
	}
}

func (_collapseBuilder *CollapseBuilder) Field(_field string) *CollapseBuilder {
	_collapseBuilder.Field_ = _field
	return _collapseBuilder
}
func (_collapseBuilder *CollapseBuilder) InnerHits(_innerHits ...*InnerHitsBuilder) *CollapseBuilder {
	_collapseBuilder.InnerHits_ = _innerHits
	return _collapseBuilder
}
func (_collapseBuilder *CollapseBuilder) MaxConcurrentGroupSearches(_maxConcurrentGroupSearches int) *CollapseBuilder {
	_collapseBuilder.MaxConcurrentGroupSearches_ = &_maxConcurrentGroupSearches
	return _collapseBuilder
}

func (_collapseBuilder *CollapseBuilder) JSON() ([]byte, error) {
	return json.Marshal(_collapseBuilder)
}

func (_collapseBuilder *CollapseBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_collapseBuilder, "", "\t")
}
