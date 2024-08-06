package errs_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/errs"
)

func Test_ErrorIsOneOf(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		t.Parallel()
		err := errors.New("an error")
		err2 := errors.New("another error")
		result := errs.IsOneOf(err, err, err2)
		assert.True(t, result)
	})
	t.Run("false", func(t *testing.T) {
		t.Parallel()
		err := errors.New("an error")
		result := errs.IsOneOf(err, errors.New("another error"))
		assert.False(t, result)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		err := errors.New("an error")
		result := errs.IsOneOf(err)
		assert.False(t, result)
	})
}
