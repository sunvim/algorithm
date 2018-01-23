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
	"bytes"
)

type pbocDesMac struct{}

var (
	DefaultInitialVector = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
)

func NewPbocDesCalculateMAC() ComputeMACer {
	return &pbocDesMac{}
}

func (this *pbocDesMac) CalculateMAC(key, data, iv []byte) ([]byte, error) {
	multiLen := len(data)/8 + 1
	multiByteLen := multiLen * 8
	data = append(data, 0x80)
	tailLen := multiByteLen - len(data)
	data = append(data, bytes.Repeat([]byte{0x00}, tailLen)...)
	//initialize vector
	getPreEightBtyes := data[:8]
	resIV := make([]byte, 8)
	resEncryptIV := make([]byte, 8)
	var err error
	for i := 0; i < 8; i++ {
		resIV[i] = getPreEightBtyes[i] ^ iv[i]
	}
	en := NewDesEnDecrypter()
	resEncryptIV, err = en.Encrypt(key, resIV)
	if err != nil {
		return nil, err
	}
	//encrypt the rest data
	for i := 1; i < multiLen; i++ {
		b2 := data[i*8 : (i+1)*8]
		for byteIndex := 0; byteIndex < 8; byteIndex++ {
			resIV[byteIndex] = b2[byteIndex] ^ resEncryptIV[byteIndex]
		}
		resEncryptIV, err = en.Encrypt(key, resIV)
		if err != nil {
			return nil, err
		}
	}
	return resEncryptIV[:len(key)], nil
}
