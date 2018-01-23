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
	"errors"
)

type DisperseOper struct{}

func NewDisperser() Disperser {
	return &DisperseOper{}
}

func (this *DisperseOper) OnceDisperse(key, data []byte) ([]byte, error) {
	//check in args
	if len(key) != 16 {
		return nil, errors.New("key is wrong,should be 16 bytes.")
	}
	if len(data) != 8 {
		return nil, errors.New("data is wrong,should be 8 bytes.")
	}
	d := make([]byte, len(data))
	copy(d, data)
	_, err := Xor(data, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	if err != nil {
		return nil, err
	}
	d = append(d, data...)
	en := New2DesEnDecrypter()
	return en.Encrypt(key, d)
}

func (this *DisperseOper) TwoTimesDisperse(key, data []byte) ([]byte, error) {
	//check in args
	if len(key) != 16 {
		return nil, errors.New("key is wrong,should be 16 bytes.")
	}
	if len(data) != 16 {
		return nil, errors.New("data is wrong,should be 16 bytes.")
	}
	tmp, err := this.OnceDisperse(key, data[:8])
	if err != nil {
		return nil, err
	}
	return this.OnceDisperse(tmp, data[8:])
}

func (this *DisperseOper) ThreeTimesDisperse(key, data []byte) ([]byte, error) {
	//check in args
	if len(key) != 16 {
		return nil, errors.New("key is wrong,should be 16 bytes.")
	}
	if len(data) != 24 {
		return nil, errors.New("data is wrong,should be 24 bytes.")
	}
	tmp, err := this.TwoTimesDisperse(key, data[:16])
	if err != nil {
		return nil, err
	}
	return this.OnceDisperse(tmp, data[16:])
}
