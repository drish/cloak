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
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"log"
	"path/filepath"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

// on Linux, Reader uses getrandom(2) if available, /dev/urandom otherwise.
// on OpenBSD, Reader uses getentropy(2).
// on other Unix-like systems, Reader reads from /dev/urandom.
// on Windows systems, Reader uses the CryptGenRandom API.
func random(size int) []byte {
	r := make([]byte, size)
	_, err := rand.Read(r)
	if err != nil {
		log.Fatal("error: ", err)
		return nil
	}

	return r
}

// reads the target file
func readFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// creates the output encrypted file without extension
func createFile(file string, content []byte) (string, error) {
	var extension = filepath.Ext(file)
	var name = file[0 : len(file)-len(extension)]
	err := ioutil.WriteFile(name, content, 0644)
	if err != nil {
		return "", err
	}
	return name, nil
}

func handleError(e error) (string, error) {
	log.Fatal(e)
	return "", e
}

// scrypt derives a 64 bytes key based from the passphrase if its provided
// or randomly generates a passphrase if its not provided.
// uses nacl box to encrypt the data using derived scrypt key
func Encrypt(path string, passphrase []byte) (string, error) {

	if len(passphrase) == 0 {
		log.Println("generating random passphrase ...")
		passphrase = random(16)
		log.Println("file passphrase: ", hex.EncodeToString(passphrase))
	} else {
		log.Println("using user defined passphrase")
	}

	// generates a 32 bytes salt
	salt := random(32)

	// recommended parameters as of 2009 are N=16384, r=8, p=1.
	// should be increased as memory latency and CPU parallelism increases.
	var key [32]byte
	keyBytes, err := scrypt.Key(passphrase, salt, 16384, 8, 1, 32)
	if err != nil {
		return handleError(err)
	}

	// trick to set a fixed slice size for nacl
	copy(key[:], keyBytes)

	// must use a different nonce for each message you encrypt with the
	// same key. Since the nonce here is 192 bits long, a random value
	// provides a sufficiently small probability of repeats.
	var nonce [24]byte
	nonceBytes := random(24)
	copy(nonce[:], nonceBytes)

	data, err := readFile(path)
	if err != nil {
		return handleError(err)
	}

	encrypted := secretbox.Seal(nonce[:], data, &nonce, &key)

	output, err := createFile(path, encrypted)
	if err != nil {
		return handleError(err)
	}

	return output, nil
}
