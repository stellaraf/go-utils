package checksum

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

var ErrNoValues = errors.New("no values provided")

var Separator string = "__"

// HashFromStrings creates a SHA256 checksum from any number of input string values.
func FromStrings(values ...string) (string, error) {
	if len(values) == 0 {
		return "", ErrNoValues
	}
	src := strings.Join(values, Separator)
	hash := sha256.New()
	_, err := hash.Write([]byte(src))
	if err != nil {
		return "", err
	}
	b := hash.Sum(nil)
	sum := hex.EncodeToString(b[:])
	return sum, nil
}

// MustHashFromStrings creates a SHA256 checksum from any number of input string values and panics on error.
func MustFromStrings(values ...string) string {
	result, err := FromStrings(values...)
	if err != nil {
		panic(err)
	}
	return result
}

// ShouldHashFromStrings creates a SHA256 checksum from any number of input string values.
// On error, a `__` joined string of input values is returned.
func ShouldFromStrings(values ...string) string {
	result, err := FromStrings(values...)
	if err != nil {
		joined := strings.Join(values, Separator)
		return joined
	}
	return result
}
