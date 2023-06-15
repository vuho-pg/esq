package esq

type SortBuilder map[string]string

func Sort(key_ string, value_ string) SortBuilder {
	return SortBuilder{key_: value_}
}
