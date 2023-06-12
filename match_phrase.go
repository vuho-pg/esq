package esq

import "github.com/vuho-pg/esq/common"

type MatchPhraseBuilder common.GenericMapType[string, *MatchPhraseFieldBuilder]

func MatchPhrase() MatchPhraseBuilder {
	return make(MatchPhraseBuilder)
}

func (m MatchPhraseBuilder) Set(key string, value *MatchPhraseFieldBuilder) MatchPhraseBuilder {
	m[key] = value
	return m
}
func (m MatchPhraseBuilder) IsQuery() {}
