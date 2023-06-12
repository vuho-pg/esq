package esq

import "github.com/vuho-pg/esq/common"

type MatchPhrasePrefixBuilder common.GenericMapType[string, *MatchPhrasePrefixFieldBuilder]

func MatchPhrasePrefix() MatchPhrasePrefixBuilder {
	return make(MatchPhrasePrefixBuilder)
}

func (m MatchPhrasePrefixBuilder) Set(key string, value *MatchPhrasePrefixFieldBuilder) MatchPhrasePrefixBuilder {
	m[key] = value
	return m
}
func (m MatchPhrasePrefixBuilder) IsQuery() {}
