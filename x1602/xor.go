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

func Xor(done, dtwo []byte) ([]byte, error) {
	if len(done) != len(dtwo) {
		return nil, errors.New("data length should be equal.")
	}
	dlen := len(done) - 1
	for dlen >= 0 {
		done[dlen] ^= dtwo[dlen]
		dlen = dlen - 1
	}
	return done, nil
}
