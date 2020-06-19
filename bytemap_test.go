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

import (
	"testing"
)

func TestByteMap(t *testing.T) {
	var m ByteSliceMap
	ಠ := []byte("fಠo")

	m.Put(ಠ, 123)
	m.PutString("foo", 456)

	v1, ok1 := m.GetString("fಠo")
	if !ok1 {
		t.Error("GetString('fಠo') ok=false, want: true")
	}
	if got, want := v1.(int), 123; got != want {
		t.Errorf("GetString('fಠo')=%v, want: %v", got, want)
	}

	v2, ok2 := m.Get([]byte{'f', 'o', 'o'})
	if !ok2 {
		t.Error("Get('foo') ok=false, want: true")
	}
	if got, want := v2.(int), 456; got != want {
		t.Errorf("Get('foo')=%v, want: %v", got, want)
	}
}

func TestByteMapMissingValue(t *testing.T) {
	var m ByteSliceMap
	m.PutString("foo", nil)

	v1, ok1 := m.GetString("fಠo")
	if ok1 {
		t.Error("GetString('fಠo') ok=true, want: false")
	}
	if v1 != nil {
		t.Errorf("GetString('fಠo')=%v, want want nil", v1)
	}

	v2, ok2 := m.Get([]byte{'f', 'o', 'o'})
	if !ok2 {
		t.Error("Get('foo') ok=false, want: true")
	}
	if v2 != nil {
		t.Errorf("Get('foo')=%v, want: nil", v2)
	}
}
