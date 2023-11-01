package base58_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/PaulShpilsher/snowflake-id/snowflake/base58"
)

func TestEncodeInvalidInput(t *testing.T) {
	_, err := base58.Encode(-1)
	if err == nil {
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
	for i := 0; i < 10000; i++ {
		s, err := base58.Encode(int64(i))
		if err != nil {
			t.Fatalf("encoding error: %v", err)
		}

		if s == "" {
			t.Fatal("empty string")
		}

		for _, x := range []byte(s) {
			if !contains(x, []byte(base58.Alphabet)) {
				t.Fatalf("invalid encoded string: %s", s)
			}
		}
	}
}

func TestEncodeDecodeRoundtrip(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	for i := 0; i < 100000; i++ {

		var n int64
		if i%2 == 0 {
			n = r.Int63()
		} else {
			n = int64(r.Int31())
		}
		s, err := base58.Encode(n)
		if err != nil {
			t.Fatalf("encoding [%d] error: %v", n, err)
		}

		m, err := base58.Decode(s)
		if err != nil {
			t.Fatalf("decoding [%s] error: %v", s, err)
		}

		if n != m {
			t.Fatalf("failed roundtrip: [%d] (%s) != [%d]", n, s, m)
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
