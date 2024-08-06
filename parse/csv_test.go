package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/parse"
)

func Test_CommaSeparatedNumbers(t *testing.T) {
	t.Run("all numbers", func(t *testing.T) {
		t.Parallel()
		result := parse.CommaSeparatedNumbers("1,2,3,4")
		assert.Equal(t, []int{1, 2, 3, 4}, result)
	})
	t.Run("mixed", func(t *testing.T) {
		t.Parallel()
		result := parse.CommaSeparatedNumbers("1,2,three,4")
		assert.Equal(t, []int{1, 2, 4}, result)
	})
}

func Test_CommaSeparatedValues(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		result := parse.CommaSeparatedValues("one,two,three")
		assert.Equal(t, []string{"one", "two", "three"}, result)
	})
	t.Run("with spaces", func(t *testing.T) {
		t.Parallel()
		result := parse.CommaSeparatedValues("one, two, three")
		assert.Equal(t, []string{"one", "two", "three"}, result)
	})
}
