package backoffutils

import (
	"context"
	"time"
)

// Fibonacci it waits for "delta * fibonacci(attempt)" time between calls.
func Fibonacci(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return fibonacci(attempt, delta)
	}
}

// FibonacciWithJitter creates an fibonacci backoff like Fibonacci does,
// but adds jitter (fractional adjustment).
func FibonacciWithJitter(delta time.Duration, jitterFraction float64) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return JitterUp(fibonacci(attempt, delta), jitterFraction)
	}
}

func fibonacci(attempt uint, delta time.Duration) time.Duration {
	var (
		pre int64
		cur int64
		i   uint
	)
	for pre, cur, i = 0, 1, 0; i < attempt; i++ {
		pre, cur = cur, pre+cur
	}
	return delta * time.Duration(pre)
}
