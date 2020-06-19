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

package triemap_test

import (
	"fmt"

	"github.com/google/triemap"
)

func ExampleRuneSliceMap() {
	var m triemap.RuneSliceMap

	m.PutString("foo", 123)
	m.Put([]rune{'b', 'a', 'r'}, 456)

	v, ok := m.Get([]rune{'f', 'o', 'o'})
	fmt.Println("Get([]rune{'f', 'o', 'o'}):", v, ok)

	v, ok = m.GetString("bar")
	fmt.Println("GetString('bar'):", v, ok)

	v, ok = m.GetString("baz")
	fmt.Println("GetString('baz'):", v, ok)
	// Output: Get([]rune{'f', 'o', 'o'}): 123 true
	// GetString('bar'): 456 true
	// GetString('baz'): <nil> false
}
