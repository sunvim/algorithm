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

type pboc3DesMac struct{}

func NewPboc3DesCalculateMAC() ComputeMACer {
	return &pboc3DesMac{}
}

/**
 * 计算MAC(hex)
 * PBOC_3DES_MAC(16的整数补8000000000000000)
 * 前n-1组使用单长密钥DES 使用密钥是密钥的左8字节）
 * 最后1组使用双长密钥3DES （使用全部16字节密钥）
 */

func (this *pboc3DesMac) CalculateMAC(key, data, iv []byte) ([]byte, error) {
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
	resIV, _ = Xor(getPreEightBtyes, iv)

	des1 := NewDesEnDecrypter()
	front8Key := key[:8]
	//前n-组采用DES加密
	for i := 1; i < multiLen; i++ {
		resEncryptIV, err = des1.Encrypt(front8Key, resIV)
		if err != nil {
			return nil, err
		}
		d := data[8*i : 8*(i+1)]
		resIV, _ = Xor(resEncryptIV, d)
	}
	//最后一组数据采用3DES 双倍长加密
	des2 := New2DesEnDecrypter()
	resEncryptIV, err = des2.Encrypt(key, resEncryptIV)
	if err != nil {
		return nil, err
	}
	return resEncryptIV, nil
}
