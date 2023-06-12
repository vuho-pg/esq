package esq

import "github.com/vuho-pg/esq/common"

type CombinedFieldsBuilder common.GenericMapType[string, *CombinedFieldsFieldBuilder]

func CombinedFields() CombinedFieldsBuilder {
	return make(CombinedFieldsBuilder)
}

func (c CombinedFieldsBuilder) Set(key string, value *CombinedFieldsFieldBuilder) CombinedFieldsBuilder {
	c[key] = value
	return c
}
func (c CombinedFieldsBuilder) IsQuery() {}
