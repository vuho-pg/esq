package esq

type CombinedFieldsInnerBuilder struct {
	Fields_                          []string        `json:"fields"`
	Query_                           string          `json:"query"`
	AutoGenerateSynonymsPhraseQuery_ *bool           `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Operator_                        *Operator       `json:"operator,omitempty"`
	MinimumShouldMatch_              *string         `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery_                  *ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

type CombinedFieldsBuilder struct {
	fieldName      string
	CombinedFields map[string]*CombinedFieldsInnerBuilder `json:"combined_fields"`
}

func (*CombinedFieldsBuilder) IsQuery() {}

func CombinedFields(
	_fieldName string,
	_fields []string,
	_query string,
) *CombinedFieldsBuilder {
	return &CombinedFieldsBuilder{
		fieldName: _fieldName,
		CombinedFields: map[string]*CombinedFieldsInnerBuilder{
			_fieldName: {
				Fields_: _fields,
				Query_:  _query,
			},
		},
	}
}

func (_combinedFields *CombinedFieldsBuilder) Fields(_fields ...string) *CombinedFieldsBuilder {
	_combinedFields.CombinedFields[_combinedFields.fieldName].Fields_ = _fields
	return _combinedFields
}
func (_combinedFields *CombinedFieldsBuilder) Query(_query string) *CombinedFieldsBuilder {
	_combinedFields.CombinedFields[_combinedFields.fieldName].Query_ = _query
	return _combinedFields
}
func (_combinedFields *CombinedFieldsBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *CombinedFieldsBuilder {
	_combinedFields.CombinedFields[_combinedFields.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return _combinedFields
}
func (_combinedFields *CombinedFieldsBuilder) Operator(_operator Operator) *CombinedFieldsBuilder {
	_combinedFields.CombinedFields[_combinedFields.fieldName].Operator_ = &_operator
	return _combinedFields
}
func (_combinedFields *CombinedFieldsBuilder) MinimumShouldMatch(_minimumShouldMatch string) *CombinedFieldsBuilder {
	_combinedFields.CombinedFields[_combinedFields.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return _combinedFields
}
func (_combinedFields *CombinedFieldsBuilder) ZeroTermsQuery(_zeroTermsQuery ZeroTermsQuery) *CombinedFieldsBuilder {
	_combinedFields.CombinedFields[_combinedFields.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return _combinedFields
}
