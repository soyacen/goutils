package backoffutils

import (
	"context"
	"math"
	"time"
)

// Exponential2 it waits for "delta * 2^attempts" time between calls.
func Exponential2(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return exponential2(attempt, delta)
	}
}

// Exponential2WithJitter creates an exponential2 backoff like Exponential2 does,
// but adds jitter (fractional adjustment).
func Exponential2WithJitter(delta time.Duration, jitterFraction float64) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return JitterUp(exponential2(attempt, delta), jitterFraction)
	}
}

// exponential return "delta * 2^attempts" time.duration
func exponential2(attempt uint, delta time.Duration) time.Duration {
	return delta * time.Duration(math.Exp2(float64(attempt)))
}
