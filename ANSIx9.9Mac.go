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

type ansix99mac struct{}

func NewANSIx99CalculateMAC() ComputeMACer {
	return &ansix99mac{}
}

//CalculateMAC 采用ANSI X9.9标准算法
func (this *ansix99mac) CalculateMAC(key, data, iv []byte) ([]byte, error) {
	dataLen := len(data)
	bn := dataLen % 8
	if bn != 0 {
		data = append(data, bytes.Repeat([]byte{0x00}, bn)...)
	}
	var (
		resIV []byte = make([]byte, 8)
		err   error
	)
	des1 := NewDesEnDecrypter()
	copy(resIV, iv)
	lastDataLen := len(data) / 8
	for index := 0; index < lastDataLen; index++ {
		resIV, err = Xor(data[index*8:(index+1)*8], resIV)
		if err != nil {
			return nil, err
		}
		resIV, err = des1.Encrypt(key, resIV)
		if err != nil {
			return nil, err
		}
	}
	return resIV, nil
}
