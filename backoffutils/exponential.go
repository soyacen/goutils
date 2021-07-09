package backoffutils

import (
	"context"
	"math"
	"time"
)

// Exponential it waits for "delta * e^attempts" time between calls.
func Exponential(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return exponential(attempt, delta)
	}
}

// ExponentialWithJitter creates an exponential backoff like Exponential does,
// but adds jitter (fractional adjustment).
func ExponentialWithJitter(delta time.Duration, jitterFraction float64) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return JitterUp(exponential(attempt, delta), jitterFraction)
	}
}

// exponential return "delta * e^attempts" time.duration
func exponential(attempt uint, delta time.Duration) time.Duration {
	return delta * time.Duration(math.Exp(float64(attempt)))
}
