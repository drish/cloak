# cloak

A simple command line passphrase based file encryption tool, its similar to `openssl enc` but it uses [scrypt](http://www.tarsnap.com/scrypt.html) for passphrate key derivation and [nacl](https://nacl.cr.yp.to/) box for encryption  


[![Build Status](https://travis-ci.org/drish/cloak.svg?branch=master)](https://travis-ci.org/drish/cloak)

## Usage

```sh
Usage: cloak [options...] file.pdf [flags...]

Options:
  encrypt	encrypts file
  decrypt	decrypts file

Flags:
  -p 		[optional] user provided passphrase, /dev/urandom if not provided
```

## Examples 

```sh
> cloak encrypt file.pdf

> cloak encrypt file.pdf -p="coolpassword"
// file

> cloak decrypt file.pdf -p="coolpassword"
```

### TODO 
	
	- flag "-overwrite" "-o" overwrites original file
	- flag "-r" encrypts all files in dir
	- efficitenly read large files using line by line chans
	- encrypt on msgpack format ?