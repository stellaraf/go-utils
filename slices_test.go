package utils_test

import (
	"testing"

	"github.com/stellaraf/go-utils"
	"github.com/stretchr/testify/assert"
)

func Test_FilterPartialDuplicates(t *testing.T) {
	t.Run("sequential", func(t *testing.T) {
		t.Parallel()
		sl := []string{"thing1", "thing that has thing1", "thing2", "thing that has thing2", "thing3"}
		exp := []string{"thing1", "thing2", "thing3"}
		result := utils.FilterPartialDuplicates(sl)
		assert.ElementsMatch(t, exp, result)
	})
	t.Run("non-sequential", func(t *testing.T) {
		t.Parallel()
		sl := []string{"thing that has thing1", "thing2", "thing1", "thing that has thing2", "thing3"}
		exp := []string{"thing1", "thing2", "thing3"}
		result := utils.FilterPartialDuplicates(sl)
		assert.ElementsMatch(t, exp, result)
	})

}
