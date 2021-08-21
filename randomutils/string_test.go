package randomutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/soyacen/goutils/randomutils"
)

func TestNumericPermString(t *testing.T) {
	permString := randomutils.NumericString(10)
	assert.Len(t, permString, 10)
	permString = randomutils.NumericString(15)
	assert.Len(t, permString, 15)
	permString = randomutils.NumericString(20)
	assert.Len(t, permString, 20)
	permString = randomutils.NumericString(21)
	assert.Len(t, permString, 21)
	permString = randomutils.NumericString(30)
	assert.Len(t, permString, 30)
	permString = randomutils.NumericString(39)
	assert.Len(t, permString, 39)
}

func TestWordString(t *testing.T) {
	permString := randomutils.WordString(10)
	assert.Len(t, permString, 10)
	permString = randomutils.WordString(15)
	assert.Len(t, permString, 15)
	permString = randomutils.WordString(20)
	assert.Len(t, permString, 20)
	permString = randomutils.WordString(21)
	assert.Len(t, permString, 21)
	permString = randomutils.WordString(30)
	assert.Len(t, permString, 30)
	permString = randomutils.WordString(39)
	assert.Len(t, permString, 39)
}
