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

type des2 struct{}

func New2DesEnDecrypter() EnDecrypter {
	return &des2{}
}

func (this *des2) Encrypt(key, data []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("error key")
	}
	//1:取密钥前8个字节数据采用DES加密
	front8Key := key[:8]
	en := NewDesEnDecrypter()
	firstRes, err := en.Encrypt(front8Key, data)
	if err != nil {
		return nil, err
	}
	//2:取密钥后8个字节采用DES解密
	back8Key := key[8:]
	secondRes, err := en.Decrypt(back8Key, firstRes)
	if err != nil {
		return nil, err
	}
	//3:再次用前8个字节的密钥采用DES加密
	last, err := en.Encrypt(front8Key, secondRes)
	if err != nil {
		return nil, err
	}
	return last, nil
}

func (this *des2) Decrypt(key, data []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("error key")
	}
	//1:取密钥前8个字节数据采用DES 解密
	front8Key := key[:8]
	en := NewDesEnDecrypter()
	firstRes, err := en.Decrypt(front8Key, data)
	if err != nil {
		return nil, err
	}
	//2:取密钥后8个字节采用DES加密
	back8Key := key[8:]
	secondRes, err := en.Encrypt(back8Key, firstRes)
	if err != nil {
		return nil, err
	}
	//3:再次用前8个字节的密钥采用DES解密
	last, err := en.Decrypt(front8Key, secondRes)
	if err != nil {
		return nil, err
	}
	return last, nil
}
