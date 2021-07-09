package bitutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBitValue(t *testing.T) {
	value := GetBitValue(10, 0)
	assert.Equal(t, 0, value)

	value = GetBitValue(10, 1)
	assert.Equal(t, 1, value)

	value = GetBitValue(10, 2)
	assert.Equal(t, 0, value)

	value = GetBitValue(10, 3)
	assert.Equal(t, 1, value)
}

func TestSetBitValue(t *testing.T) {
	value := SetBitValue(10, 0, 1)
	assert.Equal(t, 11, value)

	value = SetBitValue(10, 1, 0)
	assert.Equal(t, 8, value)

	value = SetBitValue(10, 2, 1)
	assert.Equal(t, 14, value)

	value = SetBitValue(10, 3, 0)
	assert.Equal(t, 2, value)
}

func TestReverseBitValue(t *testing.T) {
	value := ReverseBitValue(10, 0)
	assert.Equal(t, 11, value)
	value = ReverseBitValue(10, 1)
	assert.Equal(t, 8, value)
	value = ReverseBitValue(10, 2)
	assert.Equal(t, 14, value)
	value = ReverseBitValue(10, 3)
	assert.Equal(t, 2, value)
}
