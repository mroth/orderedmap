package orderedmap_test

import (
	"fmt"

	"github.com/mroth/orderedmap"
)

func Example() {
	om := orderedmap.New[string, int]()
	om.Set("foo", 1)
	om.Set("bar", 2)

	value, ok := om.Get("foo")
	if ok {
		fmt.Println(value)
	}
	// Output: 1
}
