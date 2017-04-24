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

var passphrase = flag.String("p", "", "[optional] user provided passphrase")

var usage = `Usage: cloak [options...] file.pdf [flags...]

Options:
  encrypt	encrypts file
  decrypt	decrypts file

Flags:
  -p 		[optional] user provided passphrase, /dev/urandom if not provided
`

func main() {

	encryptCommand := flag.NewFlagSet("encrypt", flag.ExitOnError)
	decryptCommand := flag.NewFlagSet("decrypt", flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()

	if flag.NArg() < 1 {
		usageAndExit("")
	}

	switch os.Args[1] {
	case "encrypt":
		encryptCommand.Parse(os.Args[2:])
		if len(encryptCommand.Args()) == 0 {
			usageAndExit("File to encrypt not provided")
		}
		path := encryptCommand.Arg(0)
		output, err := crypt.Encrypt(path, *passphrase)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("output: ", output)
		log.Println("finished ! ")
	case "decrypt":
		decryptCommand.Parse(os.Args[2:])
	default:
		usageAndExit("")
	}
}

func usageAndExit(msg string) {
	l := log.New(os.Stderr, "", 0)
	if msg != "" {
		l.Println(msg)
	}
	flag.Usage()
	os.Exit(1)
}
