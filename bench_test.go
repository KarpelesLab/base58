package base58_test

import (
	"crypto/rand"
	"testing"

	"github.com/KarpelesLab/base58"
)

// Standard base58 encode benchmarks

func BenchmarkEncode_25bytes(b *testing.B) {
	// Typical bitcoin address payload (1 version byte + 20 hash + 4 checksum)
	data := make([]byte, 25)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Encode(data)
	}
}

func BenchmarkEncode_64bytes(b *testing.B) {
	data := make([]byte, 64)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Encode(data)
	}
}

func BenchmarkEncode_256bytes(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Encode(data)
	}
}

func BenchmarkEncode_1024bytes(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Encode(data)
	}
}

// Standard base58 decode benchmarks

func BenchmarkDecode_25bytes(b *testing.B) {
	data := make([]byte, 25)
	rand.Read(data)
	encoded := base58.Bitcoin.Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Decode(encoded)
	}
}

func BenchmarkDecode_64bytes(b *testing.B) {
	data := make([]byte, 64)
	rand.Read(data)
	encoded := base58.Bitcoin.Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Decode(encoded)
	}
}

func BenchmarkDecode_256bytes(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	encoded := base58.Bitcoin.Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Decode(encoded)
	}
}

func BenchmarkDecode_1024bytes(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	encoded := base58.Bitcoin.Encode(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Decode(encoded)
	}
}

// EncodeTo (buffer reuse) benchmarks

func BenchmarkEncodeTo_25bytes(b *testing.B) {
	data := make([]byte, 25)
	rand.Read(data)
	buf := make([]byte, base58.EncodedMaxLen(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeTo(buf[:0], data)
	}
}

func BenchmarkEncodeTo_256bytes(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	buf := make([]byte, base58.EncodedMaxLen(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeTo(buf[:0], data)
	}
}

// Chunked encode benchmarks

func BenchmarkEncodeChunked_25bytes(b *testing.B) {
	data := make([]byte, 25)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeChunked(data)
	}
}

func BenchmarkEncodeChunked_64bytes(b *testing.B) {
	data := make([]byte, 64)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeChunked(data)
	}
}

func BenchmarkEncodeChunked_256bytes(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeChunked(data)
	}
}

func BenchmarkEncodeChunked_1024bytes(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeChunked(data)
	}
}

// Chunked decode benchmarks

func BenchmarkDecodeChunked_25bytes(b *testing.B) {
	data := make([]byte, 25)
	rand.Read(data)
	encoded := base58.Bitcoin.EncodeChunked(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.DecodeChunked(encoded)
	}
}

func BenchmarkDecodeChunked_64bytes(b *testing.B) {
	data := make([]byte, 64)
	rand.Read(data)
	encoded := base58.Bitcoin.EncodeChunked(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.DecodeChunked(encoded)
	}
}

func BenchmarkDecodeChunked_256bytes(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	encoded := base58.Bitcoin.EncodeChunked(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.DecodeChunked(encoded)
	}
}

func BenchmarkDecodeChunked_1024bytes(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	encoded := base58.Bitcoin.EncodeChunked(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.DecodeChunked(encoded)
	}
}

// Encode vs EncodeTo comparison at same size

func BenchmarkEncode_vs_EncodeTo_64bytes_Encode(b *testing.B) {
	data := make([]byte, 64)
	rand.Read(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.Encode(data)
	}
}

func BenchmarkEncode_vs_EncodeTo_64bytes_EncodeTo(b *testing.B) {
	data := make([]byte, 64)
	rand.Read(data)
	buf := make([]byte, base58.EncodedMaxLen(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base58.Bitcoin.EncodeTo(buf[:0], data)
	}
}
