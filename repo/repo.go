package repo

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ErrNoGoMod = errors.New("failed to locate go.mod")

// FindGoMod finds a go.mod file in the provided directory and returns the path to the directory (not the go.mod file).
func FindGoMod(start string) (string, error) {
	dir := ""
	err := filepath.Walk(start, func(path string, file fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(file.Name(), "go.mod") {
			dir = filepath.Dir(path)
			return nil
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", ErrNoGoMod
	}
	return dir, nil
}

// Root finds the root of a Go project from the caller's working directory.
func Root(depth int) (string, error) {
	start, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return RootFrom(start, depth)
}

// Root finds the root of a Go project from a given working directory.
func RootFrom(wd string, depth int) (string, error) {
	if !filepath.IsAbs(wd) {
		var err error
		wd, err = filepath.Abs(wd)
		if err != nil {
			return "", err
		}
	}
	for loops := 0; loops <= depth; loops++ {
		dir, err := FindGoMod(wd)
		if err != nil && !errors.Is(err, ErrNoGoMod) {
			return "", err
		}
		if errors.Is(err, ErrNoGoMod) {
			wd, err = filepath.Abs(filepath.Dir(wd))
			if err != nil {
				return "", err
			}
			continue
		}
		if dir != "" {
			root, err := filepath.Abs(dir)
			if err != nil {
				return "", err
			}
			return root, nil
		}
	}
	return "", fmt.Errorf("failed to locate project root, exceeded depth of %d", depth)
}
