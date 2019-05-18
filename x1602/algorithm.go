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

//EnDecrypter support only one operate form
//DES algorithm support ECB  encrypt mode
type EnDecrypter interface {
	Encrypt(key, data []byte) ([]byte, error)
	Decrypt(key, data []byte) ([]byte, error)
}

type Disperser interface {
	OnceDisperse(key, data []byte) ([]byte, error)
	TwoTimesDisperse(key, data []byte) ([]byte, error)
	ThreeTimesDisperse(key, data []byte) ([]byte, error)
}

type ComputeMACer interface {
	//CalculateMAC
	//iv : initialize vector,default 0x0000000000000000
	CalculateMAC(key, data, iv []byte) ([]byte, error)
}
