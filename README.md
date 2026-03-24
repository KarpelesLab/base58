[![GoDoc](https://godoc.org/github.com/KarpelesLab/base58?status.svg)](https://godoc.org/github.com/KarpelesLab/base58)
[![Coverage Status](https://coveralls.io/repos/github/KarpelesLab/base58/badge.svg?branch=master)](https://coveralls.io/github/KarpelesLab/base58?branch=master)

# base58

A fast, zero-dependency [Base58](https://en.wikipedia.org/wiki/Binary-to-text_encoding#Base58) encoding library for Go with support for both standard and chunked encoding.

## Features

- **Standard Base58** — variable-length encoding commonly used for Bitcoin addresses
- **Chunked Base58** — fixed-block encoding that splits input into 8-byte chunks, providing O(n) performance instead of O(n²) for large payloads
- **Buffer reuse** — `EncodeTo` allows encoding into a caller-provided buffer with zero allocations
- **Custom alphabets** — includes Bitcoin and Flickr alphabets, or create your own with `NewEncoding`

## Install

```
go get github.com/KarpelesLab/base58
```

## Usage

### Standard encoding

```go
// Encode
encoded := base58.Bitcoin.Encode(data)

// Decode
decoded, err := base58.Bitcoin.Decode(encoded)

// Encode with buffer reuse (zero allocations)
buf := make([]byte, base58.EncodedMaxLen(len(data)))
result := base58.Bitcoin.EncodeTo(buf[:0], data)
```

### Chunked encoding

Chunked encoding processes data in independent 8-byte blocks, making it significantly faster for larger payloads at the cost of a slightly longer output.

```go
// Encode
encoded := base58.Bitcoin.EncodeChunked(data)

// Decode
decoded, err := base58.Bitcoin.DecodeChunked(encoded)
```

### Custom alphabet

```go
enc := base58.NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
encoded := enc.Encode(data)
```

## Predefined alphabets

| Name | Alphabet |
|------|----------|
| `Bitcoin` | `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz` |
| `Flickr` | `123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ` |

## Performance

Standard base58 is O(n²) due to full-buffer carry propagation. Chunked encoding is O(n) since each 8-byte block is encoded independently. Benchmarks on an Intel i9-14900K:

| Size | Encode | EncodeChunked | Decode | DecodeChunked |
|------|--------|---------------|--------|---------------|
| 25 B | 403 ns | 181 ns | 199 ns | 159 ns |
| 64 B | 2.66 µs | 338 ns | 991 ns | 327 ns |
| 256 B | 45.5 µs | 1.25 µs | 11.9 µs | 1.27 µs |
| 1024 B | 679 µs | 4.93 µs | 191 µs | 5.40 µs |

Run benchmarks yourself:

```
go test -bench=. -benchmem
```
