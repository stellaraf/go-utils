package random_test

import (
	"fmt"
	"net/url"
	"regexp"
	"testing"

	"github.com/stellaraf/go-utils/random"
	"github.com/stretchr/testify/assert"
)

func Test_RandomBytes(t *testing.T) {
	t.Run("random bytes", func(t *testing.T) {
		length := 32
		result, err := random.Bytes(length)
		assert.NoError(t, err)
		assert.IsType(t, []byte{}, result)
		assert.Len(t, result, length)
	})
	t.Run("random bytes errors when length is negative", func(t *testing.T) {
		length := -1
		_, err := random.Bytes(length)
		assert.Error(t, err)
	})
}

func Test_RandomStringFromLetterSet(t *testing.T) {
	ls := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	p := regexp.MustCompile(fmt.Sprintf("[%s]+", ls))
	t.Run("random string from letter set", func(t *testing.T) {
		length := 128
		result, err := random.StringFromLetterSet(length, ls)
		assert.NoError(t, err)
		assert.IsType(t, "", result)
		assert.Len(t, result, length)
		assert.True(t, p.MatchString(result))
	})

	t.Run("random string errors when length is negative", func(t *testing.T) {
		length := -1
		_, err := random.StringFromLetterSet(length, ls)
		assert.Error(t, err)
	})
}

func Test_RandomString(t *testing.T) {
	t.Run("random string", func(t *testing.T) {
		length := 32
		result, err := random.String(length)
		assert.NoError(t, err)
		assert.IsType(t, "", result)
		assert.Len(t, result, length)
	})

}

func Test_RandomStringURLSafe(t *testing.T) {
	t.Run("random url safe string", func(t *testing.T) {
		length := 32
		result, err := random.StringURLSafe(length)
		assert.NoError(t, err)
		assert.IsType(t, "", result)
		assert.Len(t, result, length)
		encoded := url.QueryEscape(result)
		assert.Equal(t, result, encoded)
	})
}
