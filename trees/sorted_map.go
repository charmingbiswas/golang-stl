package trees

import (
	"cmp"
	"iter"
)

type SortedMap[T cmp.Ordered, V any] struct {
	data *redBlackTree[T, V]
}

// Initializes and returns a pointer to a new Sorted Map instance.
func NewSortedMap[T cmp.Ordered, V any]() *SortedMap[T, V] {
	return &SortedMap[T, V]{
		data: newRedBlackTree[T, V](),
	}
}

// Insert a key into the sorted map.
func (m *SortedMap[T, V]) Insert(key T, value V) {
	m.data.insert(key, value)
}

// Delete a key from the sorted map, returns true if key is found, otherwise false.
func (m *SortedMap[T, V]) Delete(key T) bool {
	return m.data.delete(key)
}

// Searches for a key in the sorted map, returns the value and boolean.
func (m *SortedMap[T, V]) Search(key T) (V, bool) {
	return m.data.search(key)
}

// Checks if the sorted map is empty, returns boolean
func (m *SortedMap[T, V]) IsEmpty() bool {
	return m.data.isEmpty()
}

// Checks if key exists in the sorted map, returns boolean.
func (m *SortedMap[T, V]) Has(key T) (V, bool) {
	return m.data.search(key)
}

// Retuns an iterator over the sorted map, works with 'for range' expression.
func (m *SortedMap[T, V]) Iterator() iter.Seq2[T, V] {
	return m.data.iterator()
}
