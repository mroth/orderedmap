//go:build go1.23

package orderedmap_test

import (
	"fmt"

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
