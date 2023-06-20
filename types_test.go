package utils_test

import (
	"testing"
	"time"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func Test_Guards(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		assert.True(t, utils.IsBool(true))
		assert.False(t, utils.IsBool(""))
	})
	t.Run("int", func(t *testing.T) {
		assert.True(t, utils.IsInt(0))
		assert.False(t, utils.IsInt(""))
	})
	t.Run("int8", func(t *testing.T) {
		var v int8 = -127
		assert.True(t, utils.IsInt8(v))
		assert.False(t, utils.IsInt8(""))
	})
	t.Run("int16", func(t *testing.T) {
		var v int16 = 32627
		assert.True(t, utils.IsInt16(v))
		assert.False(t, utils.IsInt16(""))
	})
	t.Run("int32", func(t *testing.T) {
		var v int32 = 2147483647
		assert.True(t, utils.IsInt32(v))
		assert.False(t, utils.IsInt32(""))
	})
	t.Run("int64", func(t *testing.T) {
		var v int64 = 9223372036854775807
		assert.True(t, utils.IsInt64(v))
		assert.False(t, utils.IsInt64(""))
	})
	t.Run("float32", func(t *testing.T) {
		var v float32 = 420.69
		assert.True(t, utils.IsFloat32(v))
		assert.False(t, utils.IsFloat32(""))
	})

	t.Run("float64", func(t *testing.T) {
		var v float64 = 420.69
		assert.True(t, utils.IsFloat64(v))
		assert.False(t, utils.IsFloat64(""))
	})

	t.Run("map", func(t *testing.T) {
		m := make(map[string]string)
		assert.True(t, utils.IsMap(m))
		assert.False(t, utils.IsMap(""))
	})

	t.Run("slice", func(t *testing.T) {
		s := []byte{}
		assert.True(t, utils.IsSlice(s))
		assert.False(t, utils.IsSlice(""))
	})
	t.Run("time", func(t *testing.T) {
		n := time.Now()
		assert.True(t, utils.IsTime(n))
		assert.False(t, utils.IsTime(""))
	})
	t.Run("uint", func(t *testing.T) {
		var v uint = 0
		assert.True(t, utils.IsUint(v))
		assert.False(t, utils.IsUint(""))
	})
	t.Run("uint8", func(t *testing.T) {
		var v uint8 = 255
		assert.True(t, utils.IsUint8(v))
		assert.False(t, utils.IsUint8(""))
	})
	t.Run("uint16", func(t *testing.T) {
		var v uint16 = 32627
		assert.True(t, utils.IsUint16(v))
		assert.False(t, utils.IsUint16(""))
	})
	t.Run("uint32", func(t *testing.T) {
		var v uint32 = 2147483647
		assert.True(t, utils.IsUint32(v))
		assert.False(t, utils.IsUint32(""))
	})
	t.Run("uint64", func(t *testing.T) {
		var v uint64 = 9223372036854775807
		assert.True(t, utils.IsUint64(v))
		assert.False(t, utils.IsUint64(""))
	})
}

func Test_Contains(t *testing.T) {
	t.Run("slice contains", func(t *testing.T) {
		s := []string{"value"}
		assert.True(t, utils.SliceContains[string](s, "value"))
		assert.False(t, utils.SliceContains[string](s, "not here"))
	})
	t.Run("map has key", func(t *testing.T) {
		m := make(map[string]string)
		m["key"] = "value"
		assert.True(t, utils.MapHasKey[string](m, "key"))
		assert.False(t, utils.MapHasKey[string](m, "not here"))
	})
}
