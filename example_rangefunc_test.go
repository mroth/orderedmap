//go:build go1.22 && goexperiment.rangefunc
// +build go1.22,goexperiment.rangefunc

package orderedmap_test

import (
	"fmt"

	"github.com/mroth/orderedmap"
)

func ExampleAll() {
	om := orderedmap.New[string, int]()
	om.Set("foo", 1)
	om.Set("bar", 2)
	om.Set("baz", 3)

	for k, v := range om.All() {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	//Output:
	// k = foo, v = 1
	// k = bar, v = 2
	// k = baz, v = 3
}

func ExampleBackward() {
	om := orderedmap.New[string, int]()
	om.Set("foo", 1)
	om.Set("bar", 2)
	om.Set("baz", 3)

	for k, v := range om.Backward() {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	//Output:
	// k = baz, v = 3
	// k = bar, v = 2
	// k = foo, v = 1
}
