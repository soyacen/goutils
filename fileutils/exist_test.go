package fileutils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soyacen/goutils/fileutils"
)

func TestIsExist(t *testing.T) {
	f, err := os.CreateTemp("", "_Go_ErrIsExist")
	assert.Nil(t, err)
	defer os.Remove(f.Name())
	defer f.Close()

	dir := filepath.Dir(f.Name())
	isExist := fileutils.IsExist(dir)
	assert.True(t, isExist)

	isExist = fileutils.IsExist(dir + "tmp")
	assert.False(t, isExist)

	isExist = fileutils.IsExist(f.Name())
	assert.True(t, isExist)

	isExist = fileutils.IsExist(filepath.Join(dir, "_Go_ErrIsNotExist"))
	assert.False(t, isExist)

}
