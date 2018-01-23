// Copyright 2017   Mobius@Shanghai
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package algorithm

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	type Test struct {
		Input  []byte
		Input2 []byte
		Want   []byte
	}
	data := []Test{
		{
			Input:  []byte{0xF6, 0xCB, 0x98, 0x6B, 0x22, 0xC8, 0xDC, 0x43},
			Input2: []byte{0x43, 0x87, 0xF3, 0x6B, 0xF0, 0xA7, 0xA0, 0x28},
			Want:   []byte{0xB5, 0x4C, 0x6B, 0x00, 0xD2, 0x6F, 0x7C, 0x6B},
		},
		{
			Input:  []byte{0xB9, 0x80, 0x29, 0x39, 0x8B, 0x4F, 0x41, 0x08},
			Input2: []byte{0x39, 0xF6, 0xE8, 0x8E, 0x64, 0x46, 0x61, 0xD9},
			Want:   []byte{0x80, 0x76, 0xC1, 0xB7, 0xEF, 0x09, 0x20, 0xD1},
		},
	}

	for _, v := range data {
		res, _ := Xor(v.Input, v.Input2)
		src := fmt.Sprintf("%X", res)
		dst := fmt.Sprintf("%X", v.Want)
		if src != dst {
			t.Errorf("failed res(%s),want(%s)n", src, dst)
		}
	}
}
