package orderedmap_test

import (
	"fmt"

	"github.com/mroth/orderedmap"
)

func Example() {
	m := orderedmap.New[string, int]()
	m.Set("foo", 1)
	m.Set("bar", 2)

	value, ok := m.Get("foo")
	if ok {
		fmt.Println(value)
	}
	// Output: 1
}
