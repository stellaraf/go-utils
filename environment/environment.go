package environment

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type EnvironmentOptions struct {
	DotEnv           bool
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

func findProjectRoot() (root string, err error) {
	start := "."
	for {
		if len(start) > 4 {
			err = fmt.Errorf("failed to locate project root, exceeded depth of 4")
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

func loadDotEnv() (err error) {
	projectRoot, err := findProjectRoot()
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

func processOptions(optionsIn *EnvironmentOptions) (options *EnvironmentOptions) {
	if optionsIn != nil {
		options = optionsIn
		if options.ProjectRootDepth == 0 {
			options.ProjectRootDepth = 4
		}
	}
	if optionsIn == nil {
		options = &EnvironmentOptions{
			DotEnv:           true,
			ProjectRootDepth: 4,
		}
	}
	return
}

func Load(ref any, options *EnvironmentOptions) (err error) {
	opts := processOptions(options)
	if opts.DotEnv {
		err = loadDotEnv()
		if err != nil {
			return
		}
	}
	data := make(map[string]any)
	t := reflect.TypeOf(ref).Elem().Elem()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if tag, ok := f.Tag.Lookup("env"); ok {
			value := os.Getenv(tag)
			data[f.Name] = value
		}
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataJSON, ref)
	return
}
