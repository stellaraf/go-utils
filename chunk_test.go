package utils_test

import (
	"testing"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func Test_ChunkSlice(t *testing.T) {
	t.Run("primitive", func(t *testing.T) {
		t.Parallel()
		in := []string{"1", "2", "3", "4", "5", "6"}
		out := utils.ChunkSlice(in, 2)
		assert.Len(t, out, 3)
	})
	t.Run("uneven", func(t *testing.T) {
		t.Parallel()
		in := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		out := utils.ChunkSlice(in, 4)
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
		out := utils.ChunkSlice(in, 2)
		assert.Len(t, out, 3)
	})
}
