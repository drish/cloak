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
	"os"
	"path/filepath"
	"testing"
)

var passphrase = []byte("wozn")

func TestDecrypt(t *testing.T) {

	file, _ := ioutil.TempFile("", "encrypt-test.txt")

	filename := file.Name()
	ext := filepath.Ext(filename)

	defer os.Remove(filename)

	ioutil.WriteFile(filename, []byte(data), 0644)

	Encrypt(filename, passphrase)

	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]

	_, err := Decrypt(name, passphrase)
	if err != nil {
		t.Fatalf("Decrypt %s: %v", filename, err)
	}

	defer os.Remove("out" + string(ext))

	if _, err := os.Stat("out" + string(ext)); os.IsNotExist(err) {
		t.Fatalf("Decrypt couldnt generate output file")
	}

}
