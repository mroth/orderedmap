# orderedmap :ant:

[![Go Reference](https://pkg.go.dev/badge/github.com/mroth/orderedmap.svg)](https://pkg.go.dev/github.com/mroth/orderedmap)

Optimal, constant time implementation of ordered maps for Go with a simple API.

Designed to work with the new [range expressions proposal][1] for Go iterators.
That proposal is likely to evolve over time, this module will likewise evolve to
track it (and thus will not be stabilized until the proposal is accepted).

[1]: https://github.com/golang/go/issues/61405

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
val, ok := om.Get("foo", 1)
```

### Iteration

On gotip you can simply range across the `All()` function, which will yield
keyvalue pairs based on their insertion order:
```go
for k, v := range om.All {
    fmt.Printf("k = %v, v = %v\n", k, v)
}
```

See also `.Reverse` to iterate from newest to oldest instead.

Alternatively, on current stable versions of Go, you can already still utilize
the _yield function_ manually, if you wish to use this module in
production code today:

```go
om.All(func(k string, v int) bool {
    fmt.Printf("k = %v, v = %v", k, v)
    return true
})
```

## Support

To use this module with new range syntax, you'll want to use gotip set to the
proposal CL:

    go install golang.org/dl/gotip@latest && gotip download 510541

That said, you can already use this module today with any version of Go that
supports generics (>=go1.18), albeit with a different syntax for iteration (see
["Usage"](#usage)).

To learn more about the [range expressions proposal][1] in general, I recommend
reading Eli Bendersky's blog post _["Preview: ranging over functions in Go"][2]_.

[2]: https://eli.thegreenplace.net/2023/preview-ranging-over-functions-in-go/

## Comparison with other Go modules

Upon my review, [wk8/go-ordered-map](https://github.com/wk8/go-ordered-map)
appeared to be the best existing library, offering constant time operations and
reasonable memory footprint. This module took some design cues from it. That
said, there are some intentional design differences -- comparing this module
with it, we optimize for:

* :bug: Simpler API (less exposed surface area, similar to standard library maps).
* :seedling: Reduced feature set (no built-in YAML serialization, for example).
* :sparkles: Use new range expressions proposal for iteration.
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
