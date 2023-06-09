package common

type GenericMapType[T comparable, V any] map[T]V

func (m GenericMapType[T, V]) Set(key T, value V) {
	m[key] = value
}
