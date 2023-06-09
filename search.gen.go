package esq

import "encoding/json"

type SearchBuilder struct {
	TrackTotalHits_   *bool                   `json:"track_total_hits,omitempty"`
	DocvalueFields_   []*DocvalueFieldBuilder `json:"docvalue_fields,omitempty"`
	Fields_           []*FieldBuilder         `json:"fields,omitempty"`
	Explain_          *bool                   `json:"explain,omitempty"`
	MinScore_         *float64                `json:"min_score,omitempty"`
	PIT_              *PITBuilder             `json:"pit,omitempty"`
	Query_            *Query                  `json:"query,omitempty"`
	SeqNoPrimaryTerm_ *bool                   `json:"seq_no_primary_term,omitempty"`
	Stats_            []string                `json:"stats,omitempty"`
	TerminateAfter_   *int                    `json:"terminate_after,omitempty"`
	Timeout_          *string                 `json:"timeout,omitempty"`
	Version_          *bool                   `json:"version,omitempty"`
	Sort_             []SortBuilder           `json:"sort,omitempty"`
	Collapse_         *CollapseBuilder        `json:"collapse,omitempty"`
	From_             *int                    `json:"from,omitempty"`
	Size_             *int                    `json:"size,omitempty"`
	SearchAfter_      []any                   `json:"search_after,omitempty"`
}

func Search() *SearchBuilder {
	return &SearchBuilder{}
}

func (_searchBuilder *SearchBuilder) TrackTotalHits(_trackTotalHits bool) *SearchBuilder {
	_searchBuilder.TrackTotalHits_ = &_trackTotalHits
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) DocvalueFields(_docvalueFields ...*DocvalueFieldBuilder) *SearchBuilder {
	_searchBuilder.DocvalueFields_ = _docvalueFields
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Fields(_fields ...*FieldBuilder) *SearchBuilder {
	_searchBuilder.Fields_ = _fields
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Explain(_explain bool) *SearchBuilder {
	_searchBuilder.Explain_ = &_explain
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) MinScore(_minScore float64) *SearchBuilder {
	_searchBuilder.MinScore_ = &_minScore
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) PIT(_pIT PITBuilder) *SearchBuilder {
	_searchBuilder.PIT_ = &_pIT
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Query(_query Query) *SearchBuilder {
	_searchBuilder.Query_ = &_query
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) SeqNoPrimaryTerm(_seqNoPrimaryTerm bool) *SearchBuilder {
	_searchBuilder.SeqNoPrimaryTerm_ = &_seqNoPrimaryTerm
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Stats(_stats ...string) *SearchBuilder {
	_searchBuilder.Stats_ = _stats
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) TerminateAfter(_terminateAfter int) *SearchBuilder {
	_searchBuilder.TerminateAfter_ = &_terminateAfter
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Timeout(_timeout string) *SearchBuilder {
	_searchBuilder.Timeout_ = &_timeout
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Version(_version bool) *SearchBuilder {
	_searchBuilder.Version_ = &_version
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Sort(_sort ...SortBuilder) *SearchBuilder {
	_searchBuilder.Sort_ = _sort
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Collapse(_collapse *CollapseBuilder) *SearchBuilder {
	_searchBuilder.Collapse_ = _collapse
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) From(_from int) *SearchBuilder {
	_searchBuilder.From_ = &_from
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) Size(_size int) *SearchBuilder {
	_searchBuilder.Size_ = &_size
	return _searchBuilder
}
func (_searchBuilder *SearchBuilder) SearchAfter(_searchAfter ...any) *SearchBuilder {
	_searchBuilder.SearchAfter_ = _searchAfter
	return _searchBuilder
}

func (_searchBuilder *SearchBuilder) JSON() ([]byte, error) {
	return json.Marshal(_searchBuilder)
}

func (_searchBuilder *SearchBuilder) JSONIndent() ([]byte, error) {
	return json.MarshalIndent(_searchBuilder, "", "\t")
}
