package esq


type CombinedFieldBuilder struct {
    CombinedFieldsField *CombinedFieldsFieldBuilder `json:"combined_field"`
}
func (c *CombinedFieldBuilder) IsQuery() {}


func CombinedField(
_fields []string,
_query string,
) *CombinedFieldBuilder {
    return &CombinedFieldBuilder {
        CombinedFieldsField:  &CombinedFieldsFieldBuilder{
                                    Fields: _fields,
                                    Query: _query,
                                 },
        }
}
func (c *CombinedFieldBuilder) SetFields(_fields ...string) *CombinedFieldBuilder {
    c.CombinedFieldsField.Fields = _fields
    return c
}
func (c *CombinedFieldBuilder) SetQuery(_query string) *CombinedFieldBuilder {
    c.CombinedFieldsField.Query = _query
    return c
}
func (c *CombinedFieldBuilder) SetAutoGenerateSynonymsPhraseQuery(_autoGenerateSynonymsPhraseQuery bool) *CombinedFieldBuilder {
    c.CombinedFieldsField.AutoGenerateSynonymsPhraseQuery = &_autoGenerateSynonymsPhraseQuery
    return c
}
func (c *CombinedFieldBuilder) SetOperator(_operator string) *CombinedFieldBuilder {
    c.CombinedFieldsField.Operator = &_operator
    return c
}
func (c *CombinedFieldBuilder) SetMinimumShouldMatch(_minimumShouldMatch string) *CombinedFieldBuilder {
    c.CombinedFieldsField.MinimumShouldMatch = &_minimumShouldMatch
    return c
}
func (c *CombinedFieldBuilder) SetZeroTermsQuery(_zeroTermsQuery string) *CombinedFieldBuilder {
    c.CombinedFieldsField.ZeroTermsQuery = &_zeroTermsQuery
    return c
}