package trees

import "cmp"

type SortedMap[T cmp.Ordered, V any] struct {
	data *redBlackTree[T, V]
}

func NewSortedMap[T cmp.Ordered, V any]() *SortedMap[T, V] {
	return &SortedMap[T, V]{
		data: newRedBlackTree[T, V](),
	}
}

func (m *SortedMap[T, V]) Insert(key T, value V) {
	m.data.insert(key, value)
}

func (m *SortedMap[T, V]) Delete(key T) bool {
	return m.data.delete(key)
}

func (m *SortedMap[T, V]) Search(key T) (V, bool) {
	return m.data.search(key)
}

func (m *SortedMap[T, V]) IsEmpty() bool {
	return m.data.isEmpty()
}
