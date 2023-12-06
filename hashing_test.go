package utils_test

import (
	"testing"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_HashFromStrings(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		expected := "e12dbfcbca890c979e7d6758bcff21a0fb0451a9a690be167d5fca07ef685228"
		result, err := utils.HashFromStrings("this", "is", "a", "test")
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("no values", func(t *testing.T) {
		result, err := utils.HashFromStrings()
		require.ErrorIs(t, err, utils.ErrNoValues)
		assert.Equal(t, "", result)
	})
}
