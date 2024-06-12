package utils_test

import (
	"math/rand"
	"testing"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func uniqueRand(min, max int, bad map[int]bool) int {
	n := rand.Intn(max-min) + min
	if bad[n] {
		return uniqueRand(min, max, bad)
	}
	return n
}

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
	t.Run("large unique", func(t *testing.T) {
		sl := make([]int, 100)
		m := map[int]bool{}
		for i := range sl {
			sl[i] = uniqueRand(1, 41713, m)
		}
		set := utils.Set(sl)
		assert.Equal(t, len(sl), len(set))
		assert.ElementsMatch(t, sl, set)
	})
}
