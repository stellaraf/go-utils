package str_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/go-utils/str"
)

func Test_Truncate(t *testing.T) {
	t.Run("long", func(t *testing.T) {
		t.Parallel()
		o := "Lorem ipsum odor amet, consectetuer adipiscing elit."
		tr := "Lorem ipsum odor amet, consectetuer..."
		r := str.Truncate(o, 38)
		assert.Len(t, r, 38)
		assert.Equal(t, tr, r)
	})
	t.Run("short", func(t *testing.T) {
		o := "Lorem ipsum odor amet."
		r := str.Truncate(o, 128)
		assert.Len(t, r, 22)
		assert.Equal(t, o, r)
	})
}
