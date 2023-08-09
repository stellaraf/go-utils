package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

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

// FindProjectRoot finds the root of a Go project.
func FindProjectRoot(depth int) (root string, err error) {
	start, err := os.Getwd()
	if err != nil {
		return
	}
	loops := 0
	for {
		loops++
		if loops > depth {
			err = fmt.Errorf("failed to locate project root, exceeded depth of %d", depth)
			return
		}
		dir, err := findGoMod(start)
		if err != nil {
			return "", err
		}
		if dir == "" {
			start, err = filepath.Abs(filepath.Dir(start))
			if err != nil {
				return "", err
			}
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
