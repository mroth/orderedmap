// Package orderedmap implements an ordered map, i.e. a map that also keeps track of
// the order in which keys were inserted.
//
// All operations are constant-time.
package orderedmap

import "container/list"

// OrderedMap is an ordered map that holds key value pairs and is able to
// iterate over values based on insertion order.
type OrderedMap[K comparable, V any] struct {
	pairs map[K]*pair[K, V]
	list  *list.List
}

type pair[K comparable, V any] struct {
	Key   K
	Value V

	element *list.Element
}

// New creates a new ordered map.
func New[K comparable, V any]() *OrderedMap[K, V] {
	return WithCapacity[K, V](0)
}

// WithCapacity creates a new ordered map with a capacity hint of n, similar to
// make(map[K]V, n).
func WithCapacity[K comparable, V any](n int) *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		pairs: make(map[K]*pair[K, V], n),
		list:  list.New(),
	}
}

// Get returns the value stored in the map for a key, or nil if no value is
// present.
// The ok result indicates whether value was found in the map.
func (m *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	if pair, present := m.pairs[key]; present {
		return pair.Value, true
	}

	return
}

// Set sets the value for a key.
func (m *OrderedMap[K, V]) Set(key K, value V) {
	if pair, present := m.pairs[key]; present {
		pair.Value = value
		return
	}

	pair := &pair[K, V]{
		Key:   key,
		Value: value,
	}
	pair.element = m.list.PushBack(pair)
	m.pairs[key] = pair
}

// Delete deletes the value for a key.
func (m *OrderedMap[K, V]) Delete(key K) {
	if pair, present := m.pairs[key]; present {
		m.list.Remove(pair.element)
		delete(m.pairs, key)
	}
}

// Len returns the length of the ordered map.
func (m *OrderedMap[K, V]) Len() int {
	if m == nil || m.pairs == nil {
		return 0
	}
	return len(m.pairs)
}
