package utils_test

import (
	"fmt"
	"net/url"
	"regexp"
	"testing"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func Test_RandomBytes(t *testing.T) {
	t.Run("random bytes", func(t *testing.T) {
		length := 32
		result, err := utils.RandomBytes(length)
		assert.NoError(t, err)
		assert.IsType(t, []byte{}, result)
		assert.Len(t, result, length)
	})
	t.Run("random bytes errors when length is negative", func(t *testing.T) {
		length := -1
		_, err := utils.RandomBytes(length)
		assert.Error(t, err)
	})
}

func Test_RandomStringFromLetterSet(t *testing.T) {
	ls := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	p := regexp.MustCompile(fmt.Sprintf("[%s]+", ls))
	t.Run("random string from letter set", func(t *testing.T) {
		length := 128
		result, err := utils.RandomStringFromLetterSet(length, ls)
		assert.NoError(t, err)
		assert.IsType(t, "", result)
		assert.Len(t, result, length)
		assert.True(t, p.MatchString(result))
	})

	t.Run("random string errors when length is negative", func(t *testing.T) {
		length := -1
		_, err := utils.RandomStringFromLetterSet(length, ls)
		assert.Error(t, err)
	})
}

func Test_RandomString(t *testing.T) {
	t.Run("random string", func(t *testing.T) {
		length := 32
		result, err := utils.RandomString(length)
		assert.NoError(t, err)
		assert.IsType(t, "", result)
		assert.Len(t, result, length)
	})

}

func Test_RandomStringURLSafe(t *testing.T) {
	t.Run("random url safe string", func(t *testing.T) {
		length := 32
		result, err := utils.RandomStringURLSafe(length)
		assert.NoError(t, err)
		assert.IsType(t, "", result)
		assert.Len(t, result, length)
		encoded := url.QueryEscape(result)
		assert.Equal(t, result, encoded)
	})
}
