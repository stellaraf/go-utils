package environment_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-utils/environment"
)

func Test_Load(t *testing.T) {
	t.Run("load without options", func(t *testing.T) {
		type Env struct {
			Key string `env:"KEY"`
		}
		var env Env
		err := environment.Load(&env, nil)
		require.NoError(t, err)
		assert.Equal(t, "value", env.Key)
	})
	t.Run("load with options", func(t *testing.T) {
		type Env struct {
			Test string `env:"TEST_VARIABLE"`
		}
		var env Env
		options := &environment.EnvironmentOptions{
			DotEnv:           false,
			ProjectRootDepth: 1,
		}
		value := "value"
		os.Setenv("TEST_VARIABLE", value)
		err := environment.Load(&env, options)
		require.NoError(t, err)
		assert.Equal(t, value, env.Test)
	})
	t.Run("load from non-default file", func(t *testing.T) {
		type Env struct {
			TestKey string `env:"TESTKEY"`
		}
		var env Env
		options := &environment.EnvironmentOptions{
			DotEnv:    true,
			FileNames: []string{".env.test"},
		}
		err := environment.Load(&env, options)
		require.NoError(t, err)
		assert.Equal(t, "TESTVAL", env.TestKey)
	})
}

func ExampleLoad() {
	type Env struct {
		Key string `env:"KEY"`
	}
	var env Env
	err := environment.Load(&env)
	if err != nil {
		panic(err)
	}
	fmt.Println(env.Key)
	// Output: value
}
