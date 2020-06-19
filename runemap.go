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

// RuneSliceMap emulates `map[[]rune]interface{}`, implemented as a Trie.
type RuneSliceMap struct {
	value    interface{}
	children map[rune]*RuneSliceMap
}

// PutString is a convenience method to insert a value using a string key.
func (n *RuneSliceMap) PutString(s string, v interface{}) *RuneSliceMap {
	return n.Put([]rune(s), v)
}

// Put inserts a value into the `RuneMap` using `[]rune` as a key
func (n *RuneSliceMap) Put(s []rune, v interface{}) *RuneSliceMap {
	for _, r := range s {
		n = n.put(r)
	}
	n.value = v
	return n
}

func (n *RuneSliceMap) put(r rune) *RuneSliceMap {
	if child, ok := n.children[r]; ok {
		return child
	}
	var child RuneSliceMap
	if n.children == nil {
		n.children = map[rune]*RuneSliceMap{r: &child}
	} else {
		n.children[r] = &child
	}
	return &child
}

// Get returns a value as mapped by the `[]rune` key and a boolean of whether the value exists in the map.
func (n *RuneSliceMap) Get(s []rune) (interface{}, bool) {
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
func (n *RuneSliceMap) GetString(s string) (interface{}, bool) {
	return n.Get([]rune(s))
}
