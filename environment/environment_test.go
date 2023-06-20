package environment_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stellaraf/go-utils/environment"
	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	t.Run("load without options", func(t *testing.T) {
		type Env struct {
			Key string `env:"KEY"`
		}
		var env *Env
		err := environment.Load(&env, nil)
		assert.NoError(t, err)
		assert.Equal(t, "value", env.Key)
	})
	t.Run("load with options", func(t *testing.T) {
		type Env struct {
			Test string `env:"TEST_VARIABLE"`
		}
		var env *Env
		options := &environment.EnvironmentOptions{
			DotEnv:           false,
			ProjectRootDepth: 1,
		}
		value := "value"
		os.Setenv("TEST_VARIABLE", value)
		err := environment.Load(&env, options)
		assert.NoError(t, err)
		assert.Equal(t, value, env.Test)
	})
}

func ExampleLoad() {
	type Env struct {
		Key string `env:"KEY"`
	}
	var env *Env
	err := environment.Load(&env)
	if err != nil {
		panic(err)
	}
	fmt.Println(env.Key)
	// Output: value
}
