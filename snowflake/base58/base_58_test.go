package base58_test

import (
	"testing"

	"github.com/PaulShpilsher/snowflake-id/snowflake/base58"
)

func TestEncodeInvalidInput(t *testing.T) {
	_, err := base58.Encode(-1)
	if err != nil {
		t.Fatal("expected to fail")
	}
}

func TestDecodeInvalidInput(t *testing.T) {
	var inputs = []string{"", "$"}
	for _, s := range inputs {
		_, err := base58.Decode("@#")
		if err == nil {
			t.Fatalf("expected to fail with \"%s\" argument", s)
		}
	}
}

func TestEncoderValidAlphabet(t *testing.T) {

	validCharacters := []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

	for i := 0; i < 10000; i++ {
		s, err := base58.Encode(int64(i))
		if err != nil {
			t.Fatalf("encoding error: %v", err)
		}

		if s == "" {
			t.Fatal("empty string")
		}

		for _, x := range []byte(s) {
			if !contains(x, validCharacters) {
				t.Fatalf("invalid encoded string: %s", s)
			}
		}
	}
}

// ----------------------------------------------------------------
// helpers
// ----------------------------------------------------------------
func contains(b byte, arr []byte) bool {
	for _, x := range arr {
		if x == b {
			return true
		}
	}
	return false
}
