package esq

type CombinedFieldsInnerBuilder struct {
	Fields_                          []string `json:"fields"`
	Query_                           string   `json:"query"`
	AutoGenerateSynonymsPhraseQuery_ *bool    `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Operator_                        *string  `json:"operator,omitempty"`
	MinimumShouldMatch_              *string  `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery_                  *string  `json:"zero_terms_query,omitempty"`
}

type CombinedFieldsBuilder struct {
	fieldName      string
	CombinedFields map[string]*CombinedFieldsInnerBuilder `json:"combined_fields"`
}

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

func (combinedFields *CombinedFieldsBuilder) Fields(_fields ...string) *CombinedFieldsBuilder {
	combinedFields.CombinedFields[combinedFields.fieldName].Fields_ = _fields
	return combinedFields
}
func (combinedFields *CombinedFieldsBuilder) Query(_query string) *CombinedFieldsBuilder {
	combinedFields.CombinedFields[combinedFields.fieldName].Query_ = _query
	return combinedFields
}
func (combinedFields *CombinedFieldsBuilder) AutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *CombinedFieldsBuilder {
	combinedFields.CombinedFields[combinedFields.fieldName].AutoGenerateSynonymsPhraseQuery_ = &_autoGenerateSynonymsPhraseQuery
	return combinedFields
}
func (combinedFields *CombinedFieldsBuilder) Operator(_operator string) *CombinedFieldsBuilder {
	combinedFields.CombinedFields[combinedFields.fieldName].Operator_ = &_operator
	return combinedFields
}
func (combinedFields *CombinedFieldsBuilder) MinimumShouldMatch(_minimumShouldMatch string) *CombinedFieldsBuilder {
	combinedFields.CombinedFields[combinedFields.fieldName].MinimumShouldMatch_ = &_minimumShouldMatch
	return combinedFields
}
func (combinedFields *CombinedFieldsBuilder) ZeroTermsQuery(_zeroTermsQuery string) *CombinedFieldsBuilder {
	combinedFields.CombinedFields[combinedFields.fieldName].ZeroTermsQuery_ = &_zeroTermsQuery
	return combinedFields
}
