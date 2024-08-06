package format_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/format"
)

func Test_CommaSeparatedValues(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		t.Parallel()
		result := format.CommaSeparatedValues([]string{"one", "two", "three"})
		assert.Equal(t, "one,two,three", result)
	})
	t.Run("numbers", func(t *testing.T) {
		t.Parallel()
		result := format.CommaSeparatedValues([]int{1, 2, 3})
		assert.Equal(t, "1,2,3", result)
	})
}
