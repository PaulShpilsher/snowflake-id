package base58

import (
	"errors"
)

var errInvalidInput = errors.New("invalid input")

////////////////////////////////////////////////////////////////
// Encoding int64 ID to base58 string
////////////////////////////////////////////////////////////////

const Alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

// Encode function converts int64 to base58 string
func Encode(n int64) (string, error) {
	if n < 58 {
		if n < 0 {
			return "", errInvalidInput
		}

		return string(Alphabet[n]), nil
	}

	buf := make([]byte, 0, 11)
	for n >= 58 {
		buf = append(buf, Alphabet[n%58])
		n /= 58
	}
	buf = append(buf, Alphabet[n])

	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf), nil
}

////////////////////////////////////////////////////////////////
// Decoding base58 string to int64
////////////////////////////////////////////////////////////////

// map to of base58 alphabet to byte value.  used to speed up decoding
var decoderMap = [...]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 1, 2, 3, 4, 5, 6, 7, 8, 255, 255, 255, 255, 255, 255, 255, 34, 35, 36, 37, 38, 39, 40, 41, 255, 42, 43, 44, 45, 46, 255, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 255, 255, 255, 255, 255, 255, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 255, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}

// Decode converts base58 string to int64 value
func Decode(s string) (int64, error) {

	if s == "" {
		return 0, errInvalidInput
	}

	var n int64

	buf := []byte(s)
	for _, x := range buf {
		v := decoderMap[x]
		if v == 0xFF {
			return 0, errInvalidInput
		}
		n = n*58 + int64(v)
	}

	return n, nil
}
