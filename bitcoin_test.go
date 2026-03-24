package base58_test

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/KarpelesLab/base58"
)

func TestNewEncodingNonASCII(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("NewEncoding with non-ASCII: expected panic, got none")
		}
		msg, ok := r.(string)
		if !ok || !strings.Contains(msg, "non-ASCII") {
			t.Fatalf("NewEncoding with non-ASCII: expected descriptive panic, got: %v", r)
		}
	}()
	// Replace first char with a non-ASCII byte (0x80)
	base58.NewEncoding("\x80" + "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
}

func TestDecodeEmptyString(t *testing.T) {
	_, err := base58.Bitcoin.Decode("")
	if err == nil {
		t.Error("Decode empty string: expected error, got nil")
	}
}

func TestDecodeInvalidCharacter(t *testing.T) {
	_, err := base58.Bitcoin.Decode("0OIl") // characters not in bitcoin alphabet
	if err == nil {
		t.Error("Decode invalid characters: expected error, got nil")
	}
}

func TestEncodeDecodeLeadingZeros(t *testing.T) {
	// Encoding with leading zero bytes should preserve them
	input := []byte{0, 0, 0, 1, 2, 3}
	encoded := base58.Bitcoin.Encode(input)
	decoded, err := base58.Bitcoin.Decode(encoded)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !bytes.Equal(input, decoded) {
		t.Errorf("leading zeros not preserved: got %x, want %x", decoded, input)
	}
}

func TestBase58Vectors(t *testing.T) {
	vecs := []string{
		"1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq:00fe7bd0e0032b8d2c1156841fa0601456aaac8f3c0ef16d8c",
		"1DhRmSGnhPjUaVPAj48zgPV9e2oRhAQFUb:008b46d254a083d10ce3f12f5e9543ba731f21f2a96feb2a60",
		"17LN2oPYRYsXS9TdYdXCCDvF2FegshLDU2:00457a36bb6beee4ead3609537da658c02623ebe88086d18c7",
		"14h2bDLZSuvRFhUL45VjPHJcW667mmRAAn:00287a57cdbe7b5cf80f76309b29756d258660072b30da677b",
	}

	buf := make([]byte, 64)

	for _, v := range vecs {
		vA := strings.SplitN(v, ":", 2)
		in := vA[0]
		out, _ := hex.DecodeString(vA[1])

		res, err := base58.Bitcoin.Decode(in)
		if err != nil {
			t.Errorf("Failed to decode %s: %s", in, err)
			continue
		}
		if !bytes.Equal(res, out) {
			t.Errorf("Bad decode %s, got %x instead of %x", in, res, out)
		}

		final := base58.Bitcoin.Encode(res)
		if final != in {
			t.Errorf("Bad encode %s, got %s instead", in, final)
		}
		final2 := base58.Bitcoin.EncodeTo(buf[:0], res)
		if string(final2) != in {
			t.Errorf("Bad encodeTo %s, got %s instead", in, final2)
		}
	}
}
