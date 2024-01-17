//go:build go1.22 && goexperiment.rangefunc
// +build go1.22,goexperiment.rangefunc

package orderedmap

import "iter"

// All returns the yield function suitable for ranging over the ordered map.
// The ordering will be oldest to newest, based on when a given key was first set.
func (om *OrderedMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for el := om.list.Front(); el != nil; el = el.Next() {
			p := el.Value.(*pair[K, V])
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}

// Backward returns the yield function suitable for ranging over the ordered map in reverse.
// The ordering will be newest to oldest, based on when a given key was first set.
func (om *OrderedMap[K, V]) Backward() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for el := om.list.Back(); el != nil; el = el.Prev() {
			p := el.Value.(*pair[K, V])
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}
