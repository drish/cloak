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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drish/cloak/crypt"
)

var usage = `Usage: cloak [options...] [flags...]

Example:

cloak encrypt -p rlycoolpass -f file.pdf

Options:
  encrypt	encrypts file
  decrypt	decrypts file

Flags:
  -f 	[required] file to encrypt
  -p 	[optional] user provided passphrase, if not provided /dev/urandom is used
`

func main() {

	encryptCommand := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encPassphrase := encryptCommand.String("p", "", "[optional] user provided passphrase to encrypt file")
	encFilepath := encryptCommand.String("f", "", "[required] file to encrypt")

	decryptCommand := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decPassphrase := decryptCommand.String("p", "", "[optional] user provided passphrase to decrypt")
	decFilepath := decryptCommand.String("f", "", "[required] file to decrypt")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	if len(os.Args) < 2 {
		usageAndExit("")
	}

	switch os.Args[1] {
	case "encrypt":
		encryptCommand.Parse(os.Args[2:])
	case "decrypt":
		decryptCommand.Parse(os.Args[2:])
	default:
		usageAndExit("")
	}

	if encryptCommand.Parsed() {

		if *encFilepath == "" {
			usageAndExit("Path to file to encrypt is required. Flag -f ")
		}

		output, err := crypt.Encrypt(*encFilepath, []byte(*encPassphrase))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("output file: ", output)
		log.Println("finished ! ")
		return
	}

	if *decPassphrase == "" {
		usageAndExit("Passphrase to decrypt file is required.")
	}

	if *decFilepath == "" {
		usageAndExit("File to decrypt is required.")
	}

	_, err := crypt.Decrypt(*decFilepath, []byte(*decPassphrase))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("finished ! ")
	return
}

func usageAndExit(msg string) {
	l := log.New(os.Stderr, "", 0)
	if msg != "" {
		l.Println(msg)
	}
	flag.Usage()
	os.Exit(1)
}
