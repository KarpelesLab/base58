package base58_test

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/KarpelesLab/base58"
)

func TestChunked(t *testing.T) {
	vectors := []string{
		"00",
		"00000000",
		"01234567abcbcdef",
		"01234567abcbcdef01234567abcbcdef01234567abcbcdef01234567abcbcdef",
		"0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f",
	}

	// extra test vectors
	for n := 0; n < 32; n++ {
		vectors = append(vectors,
			strings.Repeat("00", n),
			strings.Repeat("ff", n),
			strings.Repeat("55", n),
		)
	}

	for _, vechex := range vectors {
		vec := must(hex.DecodeString(vechex))
		enc := base58.Bitcoin.EncodeChunked(vec)
		dec, err := base58.Bitcoin.DecodeChunked(enc)
		if err != nil {
			t.Errorf("failed to decode encoded string: %s", err)
			continue
		}
		if !bytes.Equal(vec, dec) {
			t.Errorf("failed to decode back to input: %x → %s → %x", vec, enc, dec)
		}
	}
}

func TestDecodeChunkedInvalidBlockLength(t *testing.T) {
	// Encoded string lengths that give remainder 1, 4, or 8 mod 11 are invalid
	// and should return an error, not panic.
	invalidLengths := []int{1, 4, 8, 12, 15, 19}
	for _, l := range invalidLengths {
		input := strings.Repeat("1", l)
		_, err := base58.Bitcoin.DecodeChunked(input)
		if err == nil {
			t.Errorf("DecodeChunked(%q) (len=%d): expected error, got nil", input, l)
		}
	}
}

func TestDecodeChunkedOverflow(t *testing.T) {
	// "zzzzzzzzzzz" (11 z's) decodes to a value > 2^64-1,
	// which should return an overflow error.
	_, err := base58.Bitcoin.DecodeChunked("zzzzzzzzzzz")
	if err == nil {
		t.Error("DecodeChunked with overflowing block: expected error, got nil")
	}

	// A valid max block (0xFFFFFFFFFFFFFFFF) should still decode fine
	maxBlock := base58.Bitcoin.EncodeChunked([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	dec, err := base58.Bitcoin.DecodeChunked(maxBlock)
	if err != nil {
		t.Errorf("DecodeChunked with max valid block: unexpected error: %s", err)
	}
	expected := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	if !bytes.Equal(dec, expected) {
		t.Errorf("DecodeChunked max valid block: got %x, want %x", dec, expected)
	}
}

func TestDecodeChunkedSmallBlockOverflow(t *testing.T) {
	// For a 2-char encoded block (rawSize=1 byte, max value 255),
	// encode 1 byte of 0xff, then replace with a higher valid base58 value
	// that exceeds 255 when decoded.
	// 58^2-1 = 3363, which is > 255. "zz" in bitcoin base58 = 57*58+57 = 3363
	_, err := base58.Bitcoin.DecodeChunked("zz")
	if err == nil {
		t.Error("DecodeChunked with small block overflow: expected error, got nil")
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
