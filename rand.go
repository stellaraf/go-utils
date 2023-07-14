package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const urlSafe string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.~"
const alphaNumeric string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

// RandomBytes generates cryptographically secure random bytes.
func RandomBytes(n int) ([]byte, error) {
	if n < 1 {
		return nil, fmt.Errorf("length must be a positive integer")
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// RandomStringFromLetterSet generates a cryptographically secure random string from a given set
// of letters. RandomStringFromLetterSet is used as the backend for RandomString and
// RandomStringURLSafe.
func RandomStringFromLetterSet(n int, l string) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("length must be a positive integer")
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(l))))
		if err != nil {
			return "", err
		}
		b[i] = l[num.Int64()]
	}

	return string(b), nil
}

// RandomString generates a cryptographically secure random string.
func RandomString(n int) (string, error) {
	return RandomStringFromLetterSet(n, alphaNumeric)
}

// RandomStringURLSafe generates a cryptographically secure random string that is URL safe.
func RandomStringURLSafe(n int) (string, error) {
	return RandomStringFromLetterSet(n, urlSafe)
}
