package orderedmap_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mroth/orderedmap"
)

func ExampleOrderedMap_All() {
	m := orderedmap.New[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("baz", 3)

	for k, v := range m.All() {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	//Output:
	// k = foo, v = 1
	// k = bar, v = 2
	// k = baz, v = 3
}

func ExampleOrderedMap_Backward() {
	m := orderedmap.New[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("baz", 3)

	for k, v := range m.Backward() {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	//Output:
	// k = baz, v = 3
	// k = bar, v = 2
	// k = foo, v = 1
}

func ExampleOrderedMap_Keys() {
	m := orderedmap.New[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)
	m.Set("baz", 3)

	for k := range m.Keys() {
		fmt.Println(k)
	}
	//Output:
	// foo
	// bar
	// baz
}

func ExampleOrderedMap_Values() {
	m := orderedmap.New[string, string]()
	m.Set("foo", "uno")
	m.Set("bar", "dos")
	m.Set("baz", "tres")

	for v := range m.Values() {
		fmt.Println(v)
	}
	//Output:
	// uno
	// dos
	// tres
}

func TestCollect(t *testing.T) {
	type pair struct {
		k string
		v int
	}
	tests := []struct {
		name     string
		pairs    []pair
		wantKeys []string
		wantVals []int
	}{
		{
			name:     "preserves source order",
			pairs:    []pair{{"foo", 1}, {"bar", 2}, {"hey", 3}},
			wantKeys: []string{"foo", "bar", "hey"},
			wantVals: []int{1, 2, 3},
		},
		{
			name:     "duplicate key keeps position, last value wins",
			pairs:    []pair{{"foo", 1}, {"bar", 2}, {"foo", 9}},
			wantKeys: []string{"foo", "bar"},
			wantVals: []int{9, 2},
		},
		{
			name:     "empty",
			pairs:    nil,
			wantKeys: nil,
			wantVals: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := orderedmap.Collect(func(yield func(string, int) bool) {
				for _, p := range tt.pairs {
					if !yield(p.k, p.v) {
						return
					}
				}
			})

			if got := slices.Collect(m.Keys()); !slices.Equal(got, tt.wantKeys) {
				t.Errorf("keys = %v, want %v", got, tt.wantKeys)
			}
			if got := slices.Collect(m.Values()); !slices.Equal(got, tt.wantVals) {
				t.Errorf("values = %v, want %v", got, tt.wantVals)
			}
		})
	}
}
