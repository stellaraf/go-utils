package is_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/is"
)

func Test_Guards(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		assert.True(t, is.Bool(true))
		assert.False(t, is.Bool(""))
	})
	t.Run("int", func(t *testing.T) {
		assert.True(t, is.Int(0))
		assert.False(t, is.Int(""))
	})
	t.Run("int8", func(t *testing.T) {
		var v int8 = -127
		assert.True(t, is.Int8(v))
		assert.False(t, is.Int8(""))
	})
	t.Run("int16", func(t *testing.T) {
		var v int16 = 32627
		assert.True(t, is.Int16(v))
		assert.False(t, is.Int16(""))
	})
	t.Run("int32", func(t *testing.T) {
		var v int32 = 2147483647
		assert.True(t, is.Int32(v))
		assert.False(t, is.Int32(""))
	})
	t.Run("int64", func(t *testing.T) {
		var v int64 = 9223372036854775807
		assert.True(t, is.Int64(v))
		assert.False(t, is.Int64(""))
	})
	t.Run("float32", func(t *testing.T) {
		var v float32 = 420.69
		assert.True(t, is.Float32(v))
		assert.False(t, is.Float32(""))
	})

	t.Run("float64", func(t *testing.T) {
		var v float64 = 420.69
		assert.True(t, is.Float64(v))
		assert.False(t, is.Float64(""))
	})

	t.Run("map", func(t *testing.T) {
		m := make(map[string]string)
		assert.True(t, is.Map(m))
		assert.False(t, is.Map(""))
	})

	t.Run("slice", func(t *testing.T) {
		s := []byte{}
		assert.True(t, is.Slice(s))
		assert.False(t, is.Slice(""))
	})
	t.Run("time", func(t *testing.T) {
		n := time.Now()
		assert.True(t, is.Time(n))
		assert.False(t, is.Time(""))
	})
	t.Run("uint", func(t *testing.T) {
		var v uint = 0
		assert.True(t, is.Uint(v))
		assert.False(t, is.Uint(""))
	})
	t.Run("uint8", func(t *testing.T) {
		var v uint8 = 255
		assert.True(t, is.Uint8(v))
		assert.False(t, is.Uint8(""))
	})
	t.Run("uint16", func(t *testing.T) {
		var v uint16 = 32627
		assert.True(t, is.Uint16(v))
		assert.False(t, is.Uint16(""))
	})
	t.Run("uint32", func(t *testing.T) {
		var v uint32 = 2147483647
		assert.True(t, is.Uint32(v))
		assert.False(t, is.Uint32(""))
	})
	t.Run("uint64", func(t *testing.T) {
		var v uint64 = 9223372036854775807
		assert.True(t, is.Uint64(v))
		assert.False(t, is.Uint64(""))
	})
}
