//go:build gotip
// +build gotip

package orderedmap_test

// to run these tests for now:
//
//	go install golang.org/dl/gotip@latest
//	gotip download 510541
//	go test -tags=gotip

import (
	"fmt"

	"github.com/mroth/orderedmap"
)

func ExampleAll() {
	om := orderedmap.New[string, int]()
	om.Set("foo", 1)
	om.Set("bar", 2)
	om.Set("baz", 3)

	for k, v := range om.All {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	//Output:
	// k = foo, v = 1
	// k = bar, v = 2
	// k = baz, v = 3
}

func ExampleReverse() {
	om := orderedmap.New[string, int]()
	om.Set("foo", 1)
	om.Set("bar", 2)
	om.Set("baz", 3)

	for k, v := range om.Reverse {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	//Output:
	// k = baz, v = 3
	// k = bar, v = 2
	// k = foo, v = 1
}
