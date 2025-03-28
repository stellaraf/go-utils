package slice_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/random"
	"go.stellar.af/go-utils/slice"
)

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

func Test_HasAll(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		result := slice.HasAll([]string{"one", "two", "three"}, []string{"two", "three", "one"})
		assert.True(t, result)
	})
	t.Run("basic negative", func(t *testing.T) {
		t.Parallel()
		result := slice.HasAll([]string{"one", "two", "three"}, []string{"one", "two", "three", "four"})
		assert.False(t, result)
	})
	t.Run("longer head", func(t *testing.T) {
		t.Parallel()
		result := slice.HasAll([]string{"one", "two", "three", "four"}, []string{"one", "two"})
		assert.False(t, result)
	})
	t.Run("duplicates", func(t *testing.T) {
		t.Parallel()
		result := slice.HasAll([]int{1, 2, 3, 1, 3}, []int{1, 2, 3})
		assert.True(t, result)
	})
}

func Test_StringerStrings(t *testing.T) {
	t.Parallel()
	bufs := make([]*bytes.Buffer, 4)
	strs := make([]string, len(bufs))
	for i := 0; i < len(bufs); i++ {
		buf := new(bytes.Buffer)
		b, _ := random.Bytes(i * 128)
		buf.Write(b)
		bufs[i] = buf
		strs[i] = string(b)
	}
	result := slice.StringerStrings(bufs)
	assert.Equal(t, strs, result)
}

func Test_Merge(t *testing.T) {
	t.Parallel()
	s1 := []string{"one", "two", "three"}
	s2 := []string{"four", "five", "six"}
	s3 := []string{"seven", "eight", "nine"}
	exp := len(s1) + len(s2) + len(s3)
	result := slice.Merge(s1, s2, s3)
	assert.Len(t, result, exp)
}

func Test_DePointer(t *testing.T) {
	t.Parallel()
	in := make([]*string, 4)
	for i := 0; i < len(in); i++ {
		str, _ := random.String(16)
		in[i] = &str
	}
	out := slice.DePointer(in)
	assert.IsType(t, []string{}, out)
	assert.Len(t, out, len(in))
	for i := 0; i < len(in); i++ {
		o := out[i]
		e := in[i]
		assert.Equal(t, *e, o)
		assert.Equal(t, e, &o)
	}
}
