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

type Test struct {
	Key  []byte
	Data []byte
	Want []byte
}

func Test2DESEncrypt(t *testing.T) {
	en := New2DesEnDecrypter()
	d := Test{
		Key:  []byte{0xF6, 0xCB, 0x98, 0x6B, 0x22, 0xC8, 0xDC, 0x43, 0x87, 0xF3, 0x6B, 0xF0, 0xA7, 0xA0, 0x28, 0x03},
		Data: []byte{0xB9, 0x80, 0x29, 0x39, 0x8B, 0x4F, 0x41, 0x08, 0xB9, 0x80, 0x29, 0x39, 0x8B, 0x4F, 0x41, 0x08},
		Want: []byte{0x42, 0xB3, 0x13, 0x4D, 0x6B, 0xB5, 0xC9, 0x07, 0x42, 0xB3, 0x13, 0x4D, 0x6B, 0xB5, 0xC9, 0x07},
	}
	res, err := en.Encrypt(d.Key, d.Data)
	if err != nil {
		t.Error("error: ", err)
	}
	out := fmt.Sprintf("%X", res)
	expect := fmt.Sprintf("%X", d.Want)
	if out != expect {
		t.Error("failed res(%s) want(%s)\n", out, expect)
	}
}

func Test2DESDecrypt(t *testing.T) {
	de := New2DesEnDecrypter()
	d := Test{
		Key:  []byte{0xF6, 0xCB, 0x98, 0x6B, 0x22, 0xC8, 0xDC, 0x43, 0x87, 0xF3, 0x6B, 0xF0, 0xA7, 0xA0, 0x28, 0x03},
		Data: []byte{0xB9, 0x80, 0x29, 0x39, 0x8B, 0x4F, 0x41, 0x08, 0xB9, 0x80, 0x29, 0x39, 0x8B, 0x4F, 0x41, 0x08},
		Want: []byte{0x39, 0xF6, 0xE8, 0x8E, 0x64, 0x46, 0x61, 0xD9, 0x39, 0xF6, 0xE8, 0x8E, 0x64, 0x46, 0x61, 0xD9},
	}
	res, err := de.Decrypt(d.Key, d.Data)
	if err != nil {
		t.Error("error: ", err)
	}
	out := fmt.Sprintf("%X", res)
	expect := fmt.Sprintf("%X", d.Want)
	if out != expect {
		t.Error("failed res(%s) want(%s)\n", out, expect)
	}
}
