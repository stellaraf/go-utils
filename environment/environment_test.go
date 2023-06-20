package environment_test

import (
	"testing"

	"github.com/stellaraf/go-utils/environment"
	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	t.Run("load", func(t *testing.T) {
		type Env struct {
			Key string `env:"KEY"`
		}
		var env *Env
		err := environment.Load(&env, nil)
		assert.NoError(t, err)
		assert.Equal(t, "value", env.Key)
	})
}
