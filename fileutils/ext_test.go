package fileutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soyacen/goutils/fileutils"
)

func TestExtName(t *testing.T) {
	name := fileutils.ExtName("")
	assert.Equal(t, "", name)

	name = fileutils.ExtName("config")
	assert.Equal(t, "", name)

	name = fileutils.ExtName("config.yaml")
	assert.Equal(t, "yaml", name)

	name = fileutils.ExtName(".conf")
	assert.Equal(t, "conf", name)
}
