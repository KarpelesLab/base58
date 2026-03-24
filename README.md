[![GoDoc](https://godoc.org/github.com/KarpelesLab/base58?status.svg)](https://godoc.org/github.com/KarpelesLab/base58)
[![Coverage Status](https://coveralls.io/repos/github/KarpelesLab/base58/badge.svg?branch=master)](https://coveralls.io/github/KarpelesLab/base58?branch=master)

Modified implementation based on `github.com/mr-tron/base58` itself based on https://github.com/trezor/trezor-crypto/blob/master/base58.c

It's nice to have a fast base58 implementation but do we really need to have the slow version in the same lib too?

## Usage

```go
// to decode some base58 string
dec, err := base58.Bitcoin.Decode(in)
if err != nil {
    // handle err
}
// or, to encode:
enc := base58.Bitcoin.Encode(dec)
```


