package checksum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-utils/checksum"
)

func Test_HashFromStrings(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		expected := "e12dbfcbca890c979e7d6758bcff21a0fb0451a9a690be167d5fca07ef685228"
		result, err := checksum.FromStrings("this", "is", "a", "test")
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("no values", func(t *testing.T) {
		result, err := checksum.FromStrings()
		require.Error(t, err)
		assert.ErrorIs(t, err, checksum.ErrNoValues)
		assert.Equal(t, "", result)
	})
	t.Run("same hash for same values", func(t *testing.T) {
		values := []string{"this", "is", "a", "test"}
		result1, err := checksum.FromStrings(values...)
		require.NoError(t, err)
		result2, err := checksum.FromStrings(values...)
		require.NoError(t, err)
		assert.Equal(t, result1, result2)
	})
}
