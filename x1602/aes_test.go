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

func TestAES(t *testing.T) {
	en := NewAESEnDecrypter()
	d := Test{
		Key:  []byte(TopSecretKey),
		Data: []byte("0123456789ABCDEFFEDCBA98765432100123456789ABCDEF"),
		Want: []byte("0123456789ABCDEFFEDCBA98765432100123456789ABCDEF"),
	}
	res, err := en.Encrypt(d.Key, d.Data)
	if err != nil {
		t.Error("failed:" + err.Error())
	}
	// fmt.Printf("%X\n", res)
	// fmt.Printf("TopKey:%X\n", TopSecretKey)
	out, err := en.Decrypt(d.Key, res)
	if err != nil {
		t.Error("failed:" + err.Error())
	}
	src := fmt.Sprintf("%X", out)
	dst := fmt.Sprintf("%X", d.Want)
	if src != dst {
		t.Errorf("falied \nsrc(%s)\ndst(%s) \n", src, dst)
	}
}
