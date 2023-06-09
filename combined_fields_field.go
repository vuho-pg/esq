package esq

type CombinedFieldsFieldBuilder struct {
	Fields                          []string `json:"fields"`
	Query                           string   `json:"query"`
	AutoGenerateSynonymsPhraseQuery *bool    `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Operator                        *string  `json:"operator,omitempty"`
	MinimumShouldMatch              *string  `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  *string  `json:"zero_terms_query,omitempty"`
}

func CombinedFieldsField(
	_fields []string,
	_query string,
) *CombinedFieldsFieldBuilder {
	return &CombinedFieldsFieldBuilder{
		Fields: _fields,
		Query:  _query,
	}
}
func (c *CombinedFieldsFieldBuilder) SetFields(_fields ...string) *CombinedFieldsFieldBuilder {
	c.Fields = _fields
	return c
}
func (c *CombinedFieldsFieldBuilder) SetQuery(_query string) *CombinedFieldsFieldBuilder {
	c.Query = _query
	return c
}
func (c *CombinedFieldsFieldBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *CombinedFieldsFieldBuilder {
	c.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
	return c
}
func (c *CombinedFieldsFieldBuilder) SetOperator(_operator string) *CombinedFieldsFieldBuilder {
	c.Operator = &_operator
	return c
}
func (c *CombinedFieldsFieldBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *CombinedFieldsFieldBuilder {
	c.MinimumShouldMatch = &_minimumShouldMatch
	return c
}
func (c *CombinedFieldsFieldBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *CombinedFieldsFieldBuilder {
	c.ZeroTermsQuery = &_zeroTermsQuery
	return c
}
