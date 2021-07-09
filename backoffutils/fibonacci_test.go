package backoffutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	duration := fibonacci(1, time.Second)
	assert.Equal(t, time.Second, duration)

	duration = fibonacci(2, time.Second)
	assert.Equal(t, time.Second, duration)

	duration = fibonacci(3, time.Second)
	assert.Equal(t, 2*time.Second, duration)

	duration = fibonacci(4, time.Second)
	assert.Equal(t, 3*time.Second, duration)

	duration = fibonacci(5, time.Second)
	assert.Equal(t, 5*time.Second, duration)

	duration = fibonacci(6, time.Second)
	assert.Equal(t, 8*time.Second, duration)
}
