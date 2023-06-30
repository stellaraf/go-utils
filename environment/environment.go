package environment

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type EnvironmentOptions struct {
	// Options from github.com/caarlos0/env/v9.
	env.Options

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

func findGoMod(start string) (dir string, err error) {
	err = filepath.Walk(start, func(path string, file fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(file.Name(), "go.mod") {
			dir = filepath.Dir(path)
			return nil
		}
		return nil
	})
	return
}

func findProjectRoot(depth int) (root string, err error) {
	start := "."
	for {
		if len(start) > depth {
			err = fmt.Errorf("failed to locate project root, exceeded depth of %d", depth)
			return
		}
		dir, err := findGoMod(start)
		if err != nil {
			return "", err
		}
		if dir == "" {
			start += "."
			continue
		} else {
			root, err = filepath.Abs(dir)
			if err != nil {
				return "", err
			}
			break
		}
	}
	return
}

func loadDotEnv(depth int) (err error) {
	projectRoot, err := findProjectRoot(depth)
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
	libOpts := env.Options{
		Environment:           opts.Environment,
		TagName:               opts.TagName,
		RequiredIfNoDef:       opts.RequiredIfNoDef,
		OnSet:                 opts.OnSet,
		Prefix:                opts.Prefix,
		UseFieldNameByDefault: opts.UseFieldNameByDefault,
		FuncMap:               opts.FuncMap,
	}
	err = env.ParseWithOptions(ref, libOpts)
	return

}
