package esq

import "encoding/json"

type InnerHitsBuilder struct {
	Name_     *string          `json:"name,omitempty"`
	From_     *any             `json:"from,omitempty"`
	Size_     *int             `json:"size,omitempty"`
	Sort_     []SortBuilder    `json:"sort,omitempty"`
	Collapse_ *CollapseBuilder `json:"collapse,omitempty"`
}

func InnerHits() *InnerHitsBuilder {
	return &InnerHitsBuilder{}
}

func (_innerHitsBuilder *InnerHitsBuilder) Name(_name string) *InnerHitsBuilder {
	_innerHitsBuilder.Name_ = &_name
	return _innerHitsBuilder
}
func (_innerHitsBuilder *InnerHitsBuilder) From(_from any) *InnerHitsBuilder {
	_innerHitsBuilder.From_ = &_from
	return _innerHitsBuilder
}
func (_innerHitsBuilder *InnerHitsBuilder) Size(_size int) *InnerHitsBuilder {
	_innerHitsBuilder.Size_ = &_size
	return _innerHitsBuilder
}
func (_innerHitsBuilder *InnerHitsBuilder) Sort(_sort ...SortBuilder) *InnerHitsBuilder {
	_innerHitsBuilder.Sort_ = _sort
	return _innerHitsBuilder
}
func (_innerHitsBuilder *InnerHitsBuilder) Collapse(_collapse *CollapseBuilder) *InnerHitsBuilder {
	_innerHitsBuilder.Collapse_ = _collapse
	return _innerHitsBuilder
}

func (_innerHitsBuilder *InnerHitsBuilder) JSON() ([]byte, error) {
	return json.Marshal(_innerHitsBuilder)
}

func (_innerHitsBuilder *InnerHitsBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_innerHitsBuilder, "", "\t")
}
