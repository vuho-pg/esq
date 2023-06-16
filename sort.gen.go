package esq

type SortBuilder map[string]*SortValueBuilder

func Sort(key_ string, value_ *SortValueBuilder) SortBuilder {
	return SortBuilder{key_: value_}
}
