//go:build go1.23

package orderedmap

import "iter"

// All returns an iterator over key-value pairs from m.
// The ordering will be oldest to newest, based on when a key was first set.
func (m *OrderedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for el := m.list.Front(); el != nil; el = el.Next() {
			p := el.Value.(*pair[K, V])
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}

// Backward returns an iterator over key-value pairs from m in reverse.
// The ordering will be newest to oldest, based on when a key was first set.
func (m *OrderedMap[K, V]) Backward() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for el := m.list.Back(); el != nil; el = el.Prev() {
			p := el.Value.(*pair[K, V])
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}

// Keys returns an iterator over keys in m.
// The ordering will be oldest to newest, based on when a key was first set.
func (m *OrderedMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for el := m.list.Front(); el != nil; el = el.Next() {
			p := el.Value.(*pair[K, V])
			if !yield(p.Key) {
				return
			}
		}
	}
}

// Values returns an iterator over values in m.
// The ordering will be oldest to newest, based on when a key was first set.
func (m *OrderedMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for el := m.list.Front(); el != nil; el = el.Next() {
			p := el.Value.(*pair[K, V])
			if !yield(p.Value) {
				return
			}
		}
	}
}
