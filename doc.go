// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package triemap allows creating a map out of `[]rune` and `[]byte` without casting with
// `string(v)`.
//
// Due to lack of generics the values are stored and retrieved as `interface{}` and require type
// asserting like `m.Get(k).(Foo)`.
//
// Compared to casting a rune or byte slice to a string, putting data in the map is actually better
// using the std map like `m[string(k)] = v`, however getting data out of the map is much faster
// with the `RuneSliceMap`.
//
//     name                     time/op
//     GetByteMap20/ByteMap-16  1.82µs ± 0%
//     GetByteMap20/StdMap-16    883ns ±14%
//     GetRuneMap20/RuneMap-16   639ns ± 0%
//     GetRuneMap20/StdMap-16   1.45µs ± 5%
//     PutByteMap20/ByteMap-16  2.13µs ± 2%
//     PutByteMap20/StdMap-16    579ns ± 6%
//     PutRuneMap20/RuneMap-16  1.09µs ± 5%
//     PutRuneMap20/StdMap-16   1.26µs ± 4%
//
//     name                     alloc/op
//     GetByteMap20/ByteMap-16   0.00B
//     GetByteMap20/StdMap-16     320B ± 0%
//     GetRuneMap20/RuneMap-16   0.00B
//     GetRuneMap20/StdMap-16     320B ± 0%
//     PutByteMap20/ByteMap-16    320B ± 0%
//     PutByteMap20/StdMap-16     128B ± 0%
//     PutRuneMap20/RuneMap-16    320B ± 0%
//     PutRuneMap20/StdMap-16     208B ± 0%
//
//     name                     allocs/op
//     GetByteMap20/ByteMap-16    0.00
//     GetByteMap20/StdMap-16     20.0 ± 0%
//     GetRuneMap20/RuneMap-16    0.00
//     GetRuneMap20/StdMap-16     20.0 ± 0%
//     PutByteMap20/ByteMap-16    20.0 ± 0%
//     PutByteMap20/StdMap-16     19.0 ± 0%
//     PutRuneMap20/RuneMap-16    20.0 ± 0%
//     PutRuneMap20/StdMap-16     20.0 ± 0%
//
// `ByteSliceMap` is always slower than casting to string and using a standard map, but it
// reduces allocations, so it's a decent alternative on memory-constrained resources.
//
// `RuneSliceMap` performs much faster and produces less allocations than casting to string and
// using standard maps.
package triemap
