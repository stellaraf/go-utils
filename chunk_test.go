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
	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type S struct {
			One int
			Two int
		}
		in := []S{{1, 2}, {3, 4}, {5, 6}}
		out := utils.ChunkSlice(in, 2)
		assert.Len(t, out, 2)
		assert.Equal(t, 1, out[0][0].One)
		assert.Equal(t, 3, out[0][1].One)
		assert.Equal(t, 5, out[1][0].One)
	})

}
