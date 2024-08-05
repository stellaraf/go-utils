package repo_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-utils/repo"
)

func SetupDir(t *testing.T) (string, error) {
	t.Helper()
	temp, err := filepath.Abs(t.TempDir())
	if err != nil {
		return "", err
	}
	gomodPath := filepath.Join(temp, "go.mod")
	gomod, err := os.Create(gomodPath)
	if err != nil {
		return "", err
	}
	defer gomod.Close()
	_, err = gomod.Write([]byte(`module test\n\ngo 1.20\n`))
	if err != nil {
		return "", err
	}
	return temp, nil
}

func SetupSub(t *testing.T, dir string, depth int) (string, error) {
	t.Helper()
	paths := make([]string, depth+1)
	paths[0] = dir
	for i := 1; i < len(paths); i++ {
		paths[i] = fmt.Sprintf("sub%d", i)
	}

	subPath := filepath.Join(paths...)
	log.Println(subPath)
	err := os.MkdirAll(subPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return subPath, nil
}

func Test_FindGoMod(t *testing.T) {
	t.Parallel()
	dir, err := SetupDir(t)
	require.NoError(t, err)
	gomod, err := repo.FindGoMod(dir)
	require.NoError(t, err)
	assert.Equal(t, dir, gomod)
}

func Test_RootFrom(t *testing.T) {
	t.Run("exists", func(t *testing.T) {
		t.Parallel()
		root, err := SetupDir(t)
		require.NoError(t, err)
		dir, err := SetupSub(t, root, 4)
		require.NoError(t, err)
		result, err := repo.RootFrom(dir, 4)
		require.NoError(t, err)
		assert.Equal(t, root, result)
	})
	t.Run("not exists", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		dir := filepath.Join(root, "sub")
		err := os.Mkdir(dir, os.ModePerm)
		require.NoError(t, err)
		result, err := repo.RootFrom(dir, 2)
		assert.ErrorContains(t, err, "exceeded depth of 2")
		assert.Equal(t, "", result)
	})
}
