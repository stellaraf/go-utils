package slice_test

import (
	"testing"

	"math/rand"

	"github.com/stellaraf/go-utils/slice"
	"github.com/stretchr/testify/assert"
)

func uniqueRand(min, max int, bad map[int]bool) int {
	n := rand.Intn(max-min) + min
	if bad[n] {
		return uniqueRand(min, max, bad)
	}
	return n
}

func Test_Contains(t *testing.T) {
	t.Run("slice contains", func(t *testing.T) {
		assert.True(t, slice.Contains([]string{"value"}, "value"))
		assert.False(t, slice.Contains([]string{"value"}, "not here"))
		assert.True(t, slice.Contains([]int{0, 1, 2}, 1))
		assert.False(t, slice.Contains([]int{0, 1, 2}, 5))
	})
}

func Test_FilterPartialDuplicates(t *testing.T) {
	t.Run("sequential", func(t *testing.T) {
		t.Parallel()
		sl := []string{"thing1", "thing that has thing1", "thing2", "thing that has thing2", "thing3"}
		exp := []string{"thing1", "thing2", "thing3"}
		result := slice.FilterPartialDuplicates(sl)
		assert.ElementsMatch(t, exp, result)
	})
	t.Run("non-sequential", func(t *testing.T) {
		t.Parallel()
		sl := []string{"thing that has thing1", "thing2", "thing1", "thing that has thing2", "thing3"}
		exp := []string{"thing1", "thing2", "thing3"}
		result := slice.FilterPartialDuplicates(sl)
		assert.ElementsMatch(t, exp, result)
	})

}

func Test_Dedup(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		result := slice.Dedup([]string{"one", "two", "two"})
		assert.ElementsMatch(t, []string{"one", "two"}, result)
	})
	t.Run("int", func(t *testing.T) {
		t.Parallel()
		result := slice.Dedup([]int{0, 1, 2, 3, 0, 1, 2, 3})
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, result)
	})
	t.Run("large unique", func(t *testing.T) {
		sl := make([]int, 100)
		m := map[int]bool{}
		for i := range sl {
			sl[i] = uniqueRand(1, 41713, m)
		}
		set := slice.Dedup(sl)
		assert.Equal(t, len(sl), len(set))
		assert.ElementsMatch(t, sl, set)
	})
}

func Test_Chunk(t *testing.T) {
	t.Run("primitive", func(t *testing.T) {
		t.Parallel()
		in := []string{"1", "2", "3", "4", "5", "6"}
		out := slice.Chunk(in, 2)
		assert.Len(t, out, 3)
	})
	t.Run("uneven", func(t *testing.T) {
		t.Parallel()
		in := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		out := slice.Chunk(in, 4)
		assert.Len(t, out, 3)
		assert.Equal(t, [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}, {"9"}}, out)
	})
	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type S struct {
			One int
			Two int
		}
		in := []S{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}}
		out := slice.Chunk(in, 2)
		assert.Len(t, out, 3)
	})
}
