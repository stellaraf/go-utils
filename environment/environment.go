package environment

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"go.stellar.af/go-utils/repo"
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

	// FileNames defines the .env file names to look for.
	//
	// Default:
	//    []string{".env"}
	FileNames []string
}

func loadDotEnv(depth int, fileNames []string) error {
	projectRoot, err := repo.Root(depth)
	if err != nil {
		return err
	}
	exists := make([]string, 0, len(fileNames))
	for _, fileName := range fileNames {
		envFile := filepath.Join(projectRoot, fileName)
		if _, err := os.Stat(envFile); err == nil {
			exists = append(exists, envFile)
		}
	}
	err = godotenv.Load(exists...)
	if err != nil {
		return err
	}
	return nil
}

func getOptions(optionsIn []*EnvironmentOptions) *EnvironmentOptions {
	defaultOptions := &EnvironmentOptions{
		DotEnv:           true,
		ProjectRootDepth: 4,
	}
	if len(optionsIn) == 0 {
		return defaultOptions
	}
	firstOptions := optionsIn[0]
	if firstOptions != nil {
		if firstOptions.ProjectRootDepth == 0 {
			firstOptions.ProjectRootDepth = 4
		}
		return firstOptions
	}
	return defaultOptions
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
func Load(ref any, options ...*EnvironmentOptions) error {
	opts := getOptions(options)
	var fileNames []string
	if len(opts.FileNames) == 0 {
		fileNames = []string{".env"}
	} else {
		fileNames = opts.FileNames
	}
	if opts.DotEnv {
		if err := loadDotEnv(opts.ProjectRootDepth, fileNames); err != nil {
			return err
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
	return env.ParseWithOptions(ref, libOpts)
}
