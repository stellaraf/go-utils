package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

var ErrNoValues = errors.New("no values provided")

// HashFromStrings creates a SHA256 checksum from any number of input string values.
func HashFromStrings(values ...string) (string, error) {
	if len(values) == 0 {
		return "", ErrNoValues
	}
	src := strings.Join(values, "__")
	hash := sha256.New()
	_, err := hash.Write([]byte(src))
	if err != nil {
		return "", err
	}
	b := hash.Sum(nil)
	sum := hex.EncodeToString(b[:])
	return sum, nil
}
