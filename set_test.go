package utils_test

import (
	"testing"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func Test_Set(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		result := utils.Set([]string{"one", "two", "two"})
		assert.ElementsMatch(t, []string{"one", "two"}, result)
	})
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		result := utils.Set([]int{0, 1, 2, 3, 0, 1, 2, 3})
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, result)
	})
}
