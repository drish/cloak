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
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

// creates an output file
func createPlainTextFile(data, ext []byte) error {

	outputFile := "out" + string(ext)

	err := ioutil.WriteFile(outputFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// encrypted file is encoded in hex and has the following structure:
// first 24 bytes = nonce
// last line = salt
func Decrypt(path string, passphrase []byte) (string, error) {

	file, err := readFile(path)
	if err != nil {
		return handleError(err)
	}

	// split the file by new line
	full := strings.Split(string(file), "\n")

	encryptedData := full[0]
	salt := full[1]
	ext := full[2]

	// decodes salt last line
	decodedSalt, err := hex.DecodeString(salt)
	if err != nil {
		return handleError(err)
	}

	// salt should be 32 bytes
	if len(decodedSalt) != 32 {
		return handleError(err)
	}

	// decodes encrypted file
	decodedEncryptedData, err := hex.DecodeString(encryptedData)
	if err != nil {
		return handleError(err)
	}

	// decodes file extension
	decodedFileExt, err := hex.DecodeString(ext)
	if err != nil {
		return handleError(err)
	}

	// reconstruct the key from the passphrase provided by the user + salt saved on file
	var key [32]byte
	keyBytes, err := scrypt.Key(passphrase, []byte(decodedSalt), 16384, 8, 1, 32)
	if err != nil {
		return handleError(err)
	}
	copy(key[:], keyBytes)

	var decryptNonce [24]byte
	copy(decryptNonce[:], decodedEncryptedData[:24])

	decrypted, ok := secretbox.Open([]byte{}, []byte(decodedEncryptedData[24:]), &decryptNonce, &key)
	if !ok {
		log.Fatal("Unable to decrypt")
		return "", nil
	}

	err = createPlainTextFile(decrypted, decodedFileExt)
	if err != nil {
		return handleError(err)
	}

	return "", nil
}
