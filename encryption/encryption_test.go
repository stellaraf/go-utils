package encryption_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/encryption"
)

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Test_encrypt(t *testing.T) {
	t.Run("encrypts and decrypts value", func(t *testing.T) {
		key := randomString(32)
		value := randomString(128)
		encrypted, err := encryption.Encrypt(key, value)
		assert.NoError(t, err)
		decrypted, err := encryption.Decrypt(key, encrypted)
		assert.NoError(t, err)
		t.Logf("key=%s, value=%s, encrypted=%s, decrypted=%s", key, value, encrypted, decrypted)
		t.Logf("equal %t", decrypted == value)
		assert.Equal(t, value, decrypted)
	})
}
