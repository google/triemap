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

// ByteSliceMap emulates `map[[]byte]interface{}`, implemented as a Trie.
//
// It seems to perform worse than `map[string]interface{}` even when casting `string([]byte)`
type ByteSliceMap struct {
	value    interface{}
	children map[byte]*ByteSliceMap
}

// PutString is a convenience method to insert a value using a string key.
func (n *ByteSliceMap) PutString(s string, v interface{}) *ByteSliceMap {
	return n.Put([]byte(s), v)
}

// Put inserts a value into the `ByteMap` using `[]byte` as a key
func (n *ByteSliceMap) Put(s []byte, v interface{}) *ByteSliceMap {
	for _, r := range s {
		n = n.put(r)
	}
	n.value = v
	return n
}

func (n *ByteSliceMap) put(r byte) *ByteSliceMap {
	if child, ok := n.children[r]; ok {
		return child
	}
	var child ByteSliceMap
	if n.children == nil {
		n.children = map[byte]*ByteSliceMap{r: &child}
	} else {
		n.children[r] = &child
	}
	return &child
}

// Get returns a value as mapped by the `[]byte` key and a boolean of whether the value exists in the map.
func (n *ByteSliceMap) Get(s []byte) (interface{}, bool) {
	for _, r := range s {
		var ok bool
		if n, ok = n.children[r]; !ok {
			return nil, false
		}
	}
	return n.value, true
}

// GetString is a convenience method to get a value using a string key.
//
// See: `Get`
func (n *ByteSliceMap) GetString(s string) (interface{}, bool) {
	return n.Get([]byte(s))
}
