package mmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/mmap"
)

func Test_HasKey(t *testing.T) {
	t.Run("map has key", func(t *testing.T) {
		m := make(map[string]string)
		m["key"] = "value"
		assert.True(t, mmap.HasKey(m, "key"))
		assert.False(t, mmap.HasKey(m, "not here"))
	})
}

func Test_AssertValue(t *testing.T) {
	t.Run("has", func(t *testing.T) {
		t.Parallel()
		m := map[string]any{
			"key": "value",
		}
		val, ok := mmap.AssertValue[string](m, "key")
		assert.True(t, ok)
		assert.Equal(t, "value", val)
	})
	t.Run("does not have", func(t *testing.T) {
		t.Parallel()
		m := map[string]any{
			"key": "value",
		}
		_, ok := mmap.AssertValue[string](m, "not here")
		assert.False(t, ok)
	})
	t.Run("has but wrong type", func(t *testing.T) {
		t.Parallel()
		m := map[string]any{
			"key": 1,
		}
		_, ok := mmap.AssertValue[string](m, "key")
		assert.False(t, ok)
	})
}
