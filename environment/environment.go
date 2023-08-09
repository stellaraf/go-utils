package environment

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/stellaraf/go-utils"
)

type EnvironmentOptions struct {

	// Environment keys and values that will be accessible for the service. From github.com/caarlos0/env/v9
	Environment map[string]string

	// TagName specifies another tagname to use rather than the default env. From github.com/caarlos0/env/v9
	TagName string

	// RequiredIfNoDef automatically sets all env as required if they do not
	// declare 'envDefault'. From github.com/caarlos0/env/v9
	RequiredIfNoDef bool

	// OnSet allows to run a function when a value is set. From github.com/caarlos0/env/v9
	OnSet func(tag string, value interface{}, isDefault bool)

	// Prefix define a prefix for each key. From github.com/caarlos0/env/v9
	Prefix string

	// UseFieldNameByDefault defines whether or not env should use the field
	// name by default if the `env` key is missing. From github.com/caarlos0/env/v9
	UseFieldNameByDefault bool

	// Custom parse functions for different types. From github.com/caarlos0/env/v9
	FuncMap map[reflect.Type]func(v string) (interface{}, error)

	// DotEnv determines if environment variables should be loaded from a file named `.env`.
	// If `true`, `environment.Load` will look for a file named `.env` at the project root
	// and load its contents as environment variables.
	//
	// Default: `true`
	DotEnv bool

	// ProjectRootDepth defines the number of upward directories to check for a `.env` file
	// from whichever file calls `environment.Load`.
	//
	// Default: `4`
	ProjectRootDepth int
}

func loadDotEnv(depth int) (err error) {
	projectRoot, err := utils.FindProjectRoot(depth)
	if err != nil {
		return
	}
	envFile := filepath.Join(projectRoot, ".env")
	if _, err := os.Stat(envFile); err == nil {
		err = godotenv.Load(envFile)
		if err != nil {
			return err
		}
	}
	return
}

func getOptions(optionsIn []*EnvironmentOptions) (options *EnvironmentOptions) {
	defaultOptions := &EnvironmentOptions{
		DotEnv:           true,
		ProjectRootDepth: 4,
	}
	if len(optionsIn) != 0 {
		firstOptions := optionsIn[0]
		if firstOptions != nil {
			if firstOptions.ProjectRootDepth == 0 {
				firstOptions.ProjectRootDepth = 4
			}
			options = firstOptions
		} else {
			options = defaultOptions
		}
	}
	if len(optionsIn) == 0 {
		options = defaultOptions
	}
	return
}

/*
Load environment variables and store the result in passed ref.

Usage:

	type Env struct {
		Key string `env:"KEY"`
	}

	var env Env

	err := environment.Load(&env)
	env.Key // Output: "value"
*/
func Load(ref any, options ...*EnvironmentOptions) (err error) {
	opts := getOptions(options)
	if opts.DotEnv {
		err = loadDotEnv(opts.ProjectRootDepth)
		if err != nil {
			return
		}
	}
	fm := make(map[reflect.Type]env.ParserFunc)
	for key := range opts.FuncMap {
		fm[key] = opts.FuncMap[key]
	}
	libOpts := env.Options{
		Environment:           opts.Environment,
		TagName:               opts.TagName,
		RequiredIfNoDef:       opts.RequiredIfNoDef,
		OnSet:                 opts.OnSet,
		Prefix:                opts.Prefix,
		UseFieldNameByDefault: opts.UseFieldNameByDefault,
		FuncMap:               fm,
	}
	err = env.ParseWithOptions(ref, libOpts)
	return

}
