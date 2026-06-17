package orderedmap

import (
	"slices"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	tests := []struct {
		name     string
		build    func(m *OrderedMap[string, int])
		wantKeys []string
		wantVals []int
	}{
		{
			name:     "set keeps insertion order",
			build:    func(m *OrderedMap[string, int]) { m.Set("foo", 1); m.Set("bar", 2) },
			wantKeys: []string{"foo", "bar"},
			wantVals: []int{1, 2},
		},
		{
			name:     "overwrite updates value and keeps position",
			build:    func(m *OrderedMap[string, int]) { m.Set("foo", 1); m.Set("bar", 2); m.Set("foo", 9) },
			wantKeys: []string{"foo", "bar"},
			wantVals: []int{9, 2},
		},
		{
			name:     "delete removes a key and preserves order",
			build:    func(m *OrderedMap[string, int]) { m.Set("foo", 1); m.Set("bar", 2); m.Set("hey", 3); m.Delete("bar") },
			wantKeys: []string{"foo", "hey"},
			wantVals: []int{1, 3},
		},
		{
			name:     "delete of an absent key is a no-op",
			build:    func(m *OrderedMap[string, int]) { m.Set("foo", 1); m.Delete("missing") },
			wantKeys: []string{"foo"},
			wantVals: []int{1},
		},
		{
			name:     "clear empties the map but keeps it usable",
			build:    func(m *OrderedMap[string, int]) { m.Set("foo", 1); m.Set("bar", 2); m.Clear(); m.Set("baz", 3) },
			wantKeys: []string{"baz"},
			wantVals: []int{3},
		},
		{
			name:     "empty map",
			build:    func(m *OrderedMap[string, int]) {},
			wantKeys: nil,
			wantVals: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New[string, int]()
			tt.build(m)

			if got := slices.Collect(m.Keys()); !slices.Equal(got, tt.wantKeys) {
				t.Errorf("keys = %v, want %v", got, tt.wantKeys)
			}
			if got := slices.Collect(m.Values()); !slices.Equal(got, tt.wantVals) {
				t.Errorf("values = %v, want %v", got, tt.wantVals)
			}
			if got := m.Len(); got != len(tt.wantKeys) {
				t.Errorf("Len() = %d, want %d", got, len(tt.wantKeys))
			}
			// Every expected key round-trips through Get...
			for i, k := range tt.wantKeys {
				if v, ok := m.Get(k); !ok || v != tt.wantVals[i] {
					t.Errorf("Get(%q) = %v, %v; want %v, true", k, v, ok, tt.wantVals[i])
				}
			}
			// ...and an unknown key reports not-present.
			if _, ok := m.Get("absent"); ok {
				t.Error("Get(absent) returned ok = true, want false")
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	m := New[string, int]()
	for b.Loop() {
		m.Set("foo", 1)
	}
}

func BenchmarkGet(b *testing.B) {
	m := New[string, int]()
	m.Set("foo", 1)
	for b.Loop() {
		_, _ = m.Get("foo")
	}
}
