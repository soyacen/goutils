package randomutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soyacen/goutils/randomutils"
)

func TestNumericPermString(t *testing.T) {
	permString := randomutils.NumericPermString(10)
	assert.Len(t, permString, 10)
	permString = randomutils.NumericPermString(15)
	assert.Len(t, permString, 15)
	permString = randomutils.NumericPermString(20)
	assert.Len(t, permString, 20)
	permString = randomutils.NumericPermString(21)
	assert.Len(t, permString, 21)
	permString = randomutils.NumericPermString(30)
	assert.Len(t, permString, 30)
	permString = randomutils.NumericPermString(39)
	assert.Len(t, permString, 39)
}
