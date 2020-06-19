# triemap
Map using `[]rune` or `[]byte` instead of `string`

> Disclaimer: This is not an official Google product.

* `RuneSliceMap` is much faster than casting and using a map like: `map[string(slice)]`.
* `ByteSliceMap` is considerably slower than standard maps, but it reduces allocations.

## Usage

```go
var m triemap.RuneSliceMap

m.PutString("foo", 123)
m.Put([]rune{'b', 'a', 'r'}, 456)

v, ok := m.Get([]rune{'f', 'o', 'o'})
fmt.Println(v, ok) // 123, true

v, ok = m.GetString("bar")
fmt.Println(v, ok) // 456, true

v, ok = m.GetString("baz")
fmt.Println(v, ok) // nil, baz
```

## Benchmarks

```
name                     time/op
GetByteMap20/ByteMap-16  1.82µs ± 0%
GetByteMap20/StdMap-16    883ns ±14%
GetRuneMap20/RuneMap-16   639ns ± 0%
GetRuneMap20/StdMap-16   1.45µs ± 5%
PutByteMap20/ByteMap-16  2.13µs ± 2%
PutByteMap20/StdMap-16    579ns ± 6%
PutRuneMap20/RuneMap-16  1.09µs ± 5%
PutRuneMap20/StdMap-16   1.26µs ± 4%

name                     alloc/op
GetByteMap20/ByteMap-16   0.00B
GetByteMap20/StdMap-16     320B ± 0%
GetRuneMap20/RuneMap-16   0.00B
GetRuneMap20/StdMap-16     320B ± 0%
PutByteMap20/ByteMap-16    320B ± 0%
PutByteMap20/StdMap-16     128B ± 0%
PutRuneMap20/RuneMap-16    320B ± 0%
PutRuneMap20/StdMap-16     208B ± 0%

name                     allocs/op
GetByteMap20/ByteMap-16    0.00
GetByteMap20/StdMap-16     20.0 ± 0%
GetRuneMap20/RuneMap-16    0.00
GetRuneMap20/StdMap-16     20.0 ± 0%
PutByteMap20/ByteMap-16    20.0 ± 0%
PutByteMap20/StdMap-16     19.0 ± 0%
PutRuneMap20/RuneMap-16    20.0 ± 0%
PutRuneMap20/StdMap-16     20.0 ± 0%
```

## Code of Conduct

[Same as Go](https://golang.org/conduct)

## Contributing

Read our [contributions](CONTRIBUTING.md) doc.

tl;dr: Get Google's CLA and send a PR!

## Licence

[Apache 2.0](LICENSE)
