<p align="center">
  <img src="https://rawgit.com/drish/cloak/master/cloak-logo.png" height="180" />
  <h3 align="center">Cloak (beta)</h3>
  <p align="center">A simple command line passphrase based file encryption tool.</p>
  <p align="center">
    <a href="https://travis-ci.org/drish/cloak"><img src="https://travis-ci.org/libeclipse/memguard.svg?branch=master"></a>
    <a href="https://github.com/drish/cloak/blob/master/LICENSE)"><img src="http://img.shields.io/badge/license-Apache%20V2-blue.svg"></a>
    <!-- <a href="https://ci.appveyor.com/project/libeclipse/memguard/branch/master"><img src="https://ci.appveyor.com/api/projects/status/g6cg347cam7lli5m/branch/master?svg=true"></a> -->
    <a href="https://goreportcard.com/report/github.com/drish/cloak"><img src="https://goreportcard.com/badge/github.com/drish/cloak"></a>
  </p>
</p>

---

Cloak is a simple command line passphrase based file encryption tool, its similar to `openssl enc`, but it uses [scrypt](http://www.tarsnap.com/scrypt.html) for passphrase key derivation and [nacl](https://nacl.cr.yp.to/) box for encryption.

This tools is still a WIP.


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
// 2017/04/30 15:13:21 generating random passphrase ...
// 2017/04/30 15:13:21 file passphrase:  14abe93eb3347f91ad6c90f4ed3d9c8f
// 2017/04/30 15:13:21 output file:  file
// 2017/04/30 15:13:21 finished !  

> cloak encrypt -f details.pdf -p coolpassphrase
// 2017/04/30 15:15:06 using user defined passphrase
// 2017/04/30 15:15:06 output file:  details
// 2017/04/30 15:15:06 finished ! 

> cloak decrypt -f details.pdf -p coolpassphrase
// 2017/04/30 15:16:26 finished ! 

```

### TODO 
	
- flag "-overwrite" "-o" overwrites original file
- flag "-r" encrypts all files in dir
- efficitenly read large files using line by line chans
- encrypt using msgpack format ?
- key splitting using shamir
- human readable passphrase generator ?
