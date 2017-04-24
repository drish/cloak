// Copyright © 2017 carlos derich <carlosderich@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crypt

import (
	"io/ioutil"
	"testing"
)

func TestEncrypt(t *testing.T) {

	file, _ := ioutil.TempFile("", "encrypt-test.txt")

	filename := file.Name()

	// https://en.wikiquote.org/wiki/Rick_Cook
	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better idiot-proof programs, and the Universe trying " +
		"to produce bigger and better idiots. So far, the Universe is winning."

	ioutil.WriteFile(filename, []byte(data), 0644)

	_, err := Encrypt(filename, "")
	if err != nil {
		t.Fatalf("Encrypt %s: %v", filename, err)
	}
}
