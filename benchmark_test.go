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

package triemap

import "testing"

var (
	keys = []string{
		"Lorem",
		"ipsum",
		"dolor",
		"sit",
		"amet",
		"consectetur",
		"adipiscing",
		"elit",
		"Donec",
		"euismod",
		"leo",
		"id",
		"laoreet",
		"feugiat",
		"augue",
		"magna",
		"tincidunt",
		"dolor",
		"a",
		"rhoncus",
	}

	values = []string{
		"rhoncus",
		"a",
		"dolor",
		"tincidunt",
		"magna",
		"augue",
		"feugiat",
		"laoreet",
		"id",
		"leo",
		"euismod",
		"Donec",
		"elit",
		"adipiscing",
		"consectetur",
		"amet",
		"sit",
		"dolor",
		"ipsum",
		"Lorem",
	}
)

func BenchmarkPutRuneMap20(b *testing.B) {
	var runeKeys [][]rune
	for _, v := range keys {
		runeKeys = append(runeKeys, []rune(v))
	}

	runeMaps := make([]RuneSliceMap, 20)
	stdMaps := make([]map[string]string, 20)
	for i := range stdMaps {
		stdMaps[i] = make(map[string]string)
	}

	testCases := []struct {
		desc string
		put  func(int, []rune, string)
	}{
		{
			"RuneMap",
			func(i int, key []rune, v string) { runeMaps[i].Put(key, v) },
		},
		{
			"StdMap",
			func(i int, key []rune, v string) { stdMaps[i][string(key)] = v },
		},
	}

	b.ResetTimer()
	for _, tc := range testCases {
		b.Run(tc.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 20; j++ {
					tc.put(j, runeKeys[j], values[j])
				}
			}
		})
	}
}

func BenchmarkGetRuneMap20(b *testing.B) {
	var runeKeys [][]rune
	for _, v := range keys {
		runeKeys = append(runeKeys, []rune(v))
	}

	var r RuneSliceMap
	m := make(map[string]string)
	for j := 0; j < 20; j++ {
		r.Put(runeKeys[j], values[j])
		m[keys[j]] = values[j]
	}

	testCases := []struct {
		desc string
		get  func([]rune) (interface{}, bool)
	}{
		{"RuneMap", func(key []rune) (interface{}, bool) { return r.Get(key) }},
		{"StdMap", func(key []rune) (interface{}, bool) {
			v, ok := m[string(key)]
			return v, ok
		}},
	}

	b.ResetTimer()
	for _, tc := range testCases {
		b.Run(tc.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 20; j++ {
					_, _ = tc.get(runeKeys[j])
				}
			}
		})
	}
}

func BenchmarkPutByteMap20(b *testing.B) {
	var byteKeys [][]byte
	for _, v := range keys {
		byteKeys = append(byteKeys, []byte(v))
	}

	byteMaps := make([]ByteSliceMap, 20)
	stdMaps := make([]map[string]string, 20)
	for i := range stdMaps {
		stdMaps[i] = make(map[string]string)
	}

	testCases := []struct {
		desc string
		put  func(int, []byte, string)
	}{
		{
			"ByteMap",
			func(i int, key []byte, v string) { byteMaps[i].Put(key, v) },
		},
		{
			"StdMap",
			func(i int, key []byte, v string) { stdMaps[i][string(key)] = v },
		},
	}

	b.ResetTimer()
	for _, tc := range testCases {
		b.Run(tc.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 20; j++ {
					tc.put(j, byteKeys[j], values[j])
				}
			}
		})
	}
}

func BenchmarkGetByteMap20(b *testing.B) {
	var byteKeys [][]byte
	for _, v := range keys {
		byteKeys = append(byteKeys, []byte(v))
	}

	var r ByteSliceMap
	m := make(map[string]string)
	for j := 0; j < 20; j++ {
		r.Put(byteKeys[j], values[j])
		m[keys[j]] = values[j]
	}

	testCases := []struct {
		desc string
		get  func([]byte) (interface{}, bool)
	}{
		{"ByteMap", func(key []byte) (interface{}, bool) { return r.Get(key) }},
		{"StdMap", func(key []byte) (interface{}, bool) {
			v, ok := m[string(key)]
			return v, ok
		}},
	}

	b.ResetTimer()
	for _, tc := range testCases {
		b.Run(tc.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 20; j++ {
					_, _ = tc.get(byteKeys[j])
				}
			}
		})
	}
}
