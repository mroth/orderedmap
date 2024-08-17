# orderedmap :ant:

[![Go Reference](https://pkg.go.dev/badge/github.com/mroth/orderedmap.svg)](https://pkg.go.dev/github.com/mroth/orderedmap)

Optimal, constant time implementation of ordered maps for Go with a simple API.

Supports handy new style iteration via the range function in go1.23 and greater.


## Usage

### Basic operations

Creating an ordered map uses any generic type `K comparable, V any`:
```go
// equivalent of: make(map[string]int)
om := orderedmap.New[string, int]()
```

You can specify an initial capacity hint:
```go
// equivalent of: make(map[string]int, 1000)
om := orderedmap.WithCapacity[string, int](1000)
```

Setting a value:
```go
// equivalent of om["foo"] = 1
om.Set("foo", 1)
```

Retrieving a value is equally simple, and uses the same `bool` ok secondary
return pattern to indicate whether a value was found in the map:
```go
// equivalent of val, ok := om["foo"]
val, ok := om.Get("foo")
```

## Iteration :sparkles:

On go1.23, you can simply range across the `All()` function, which will yield
key value pairs based on their insertion order:
```go
for k, v := range om.All() {
    fmt.Printf("k = %v, v = %v\n", k, v)
}
```

See also `Backward()` to iterate from newest to oldest instead.


### Support

To use this module with new iterator range syntax, you'll need to use go1.23 or
greater.  The functionality is hidden behind build constraints so you can use
this module today with any version of Go that supports generics (>=go1.18),
albeit without the handy range iterator support!


## Comparison with other Go modules

Upon my review, [wk8/go-ordered-map](https://github.com/wk8/go-ordered-map)
appeared to be the best existing library, offering constant time operations and
reasonable memory footprint. This module took some design cues from it. That
said, there are some intentional design differences -- comparing this module
with it, we optimize for:

* :bug: Simpler API (less exposed surface area, similar to standard library maps).
* :seedling: Reduced feature set (no built-in YAML serialization, for example).
* :sparkles: Use new range expressions for easy iteration on go1.23 and greater.
* :zap: Equally performant.
* :zero: Zero dependencies.

### Other alternatives
As per other options, the README from `wk8/go-ordered-map` offers a summary:

* [iancoleman/orderedmap](https://github.com/iancoleman/orderedmap) only accepts
  `string` keys, its `Delete` operations are linear.
* [cevaris/ordered_map](https://github.com/cevaris/ordered_map) uses a channel
  for iterations, and leaks goroutines if the iteration is interrupted before
  fully traversing the map.
* [mantyr/iterator](https://github.com/mantyr/iterator) also uses a channel for
  iterations, and its `Delete` operations are linear.
* [samdolan/go-ordered-map](https://github.com/samdolan/go-ordered-map) adds
  unnecessary locking (users should add their own locking instead if they need
  it), its `Delete` and `Get` operations are linear, iterations trigger a linear
  memory allocation.
