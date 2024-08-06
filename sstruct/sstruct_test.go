package sstruct_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-utils/sstruct"
)

func Test_GetTag(t *testing.T) {
	type S struct {
		Field string `tag:"FIELD"`
	}
	t.Run("pointer", func(t *testing.T) {
		t.Parallel()
		s := &S{Field: "field"}
		result, err := sstruct.GetTag(s, "Field", "tag")
		require.NoError(t, err)
		assert.Equal(t, "FIELD", result)
	})
	t.Run("real", func(t *testing.T) {
		t.Parallel()
		s := S{Field: "field"}
		result, err := sstruct.GetTag(s, "Field", "tag")
		require.NoError(t, err)
		assert.Equal(t, "FIELD", result)
	})
	t.Run("errors when field doesn't exist", func(t *testing.T) {
		t.Parallel()
		s := &S{Field: "field"}
		_, err := sstruct.GetTag(s, "NotHere", "tag")
		require.Error(t, err)
	})
	t.Run("errors when tag is empty", func(t *testing.T) {
		t.Parallel()
		type S struct {
			Field string `tag:""`
		}
		s := &S{Field: "field"}
		_, err := sstruct.GetTag(s, "NotHere", "tag")
		require.Error(t, err)
	})
}

func Test_SetValue(t *testing.T) {
	type S struct {
		Field string
	}
	t.Run("pointer", func(t *testing.T) {
		s := &S{}
		ss, err := sstruct.SetValue(s, "Field", "field")
		require.NoError(t, err)
		assert.Equal(t, "field", ss.Field)
		assert.Equal(t, "field", s.Field)
	})
	t.Run("non-pointer error", func(t *testing.T) {
		_, err := sstruct.SetValue(S{}, "Field", "field")
		assert.Error(t, err)
	})
	t.Run("non-existent field", func(t *testing.T) {
		s := &S{}
		ss, err := sstruct.SetValue(s, "NotHere", "value")
		require.NoError(t, err)
		assert.Equal(t, "", s.Field)
		assert.Equal(t, "", ss.Field)
	})
}
