# cloak

A simple command line passphrase based file encryption tool, its similar to `openssl enc`, but it uses [scrypt](http://www.tarsnap.com/scrypt.html) for passphrase key derivation and [nacl](https://nacl.cr.yp.to/) box for encryption.

WIP

[![Build Status](https://travis-ci.org/drish/cloak.svg?branch=master)](https://travis-ci.org/drish/cloak)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/drish/cloak/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/drish/cloak)](https://goreportcard.com/report/github.com/drish/cloak)

## Usage

```sh
Usage: cloak [options...] [flags...]

Example:

cloak encrypt -p rlycoolpass -f file.pdf

Options:
  encrypt	encrypts file
  decrypt	decrypts file

Flags:
  -f 	[required] file to encrypt
  -p 	[optional] user provided passphrase, if not provided /dev/urandom is used
```

## Examples 

```sh
> cloak encrypt -f file.pdf
// passphrase generated: 

> cloak encrypt -f file.pdf -p coolpassphrase
// file

> cloak decrypt -f file.pdf -p coolpassphrase
```

### TODO 
	
- flag "-overwrite" "-o" overwrites original file
- flag "-r" encrypts all files in dir
- efficitenly read large files using line by line chans
- encrypt using msgpack format ?
- key splitting using shamir
- human readable passphrase generator ?
