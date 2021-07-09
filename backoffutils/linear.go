package backoffutils

import (
	"context"
	"time"
)

// Linear it waits for "delta * attempt" time between calls.
func Linear(delta time.Duration) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return linear(attempt, delta)
	}
}

// LinearWithJitter creates an linear backoff like Linear does,
// but adds jitter (fractional adjustment).
func LinearWithJitter(delta time.Duration, jitterFraction float64) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		return JitterUp(delta, jitterFraction)
	}
}

func linear(attempt uint, delta time.Duration) time.Duration {
	return delta * time.Duration(attempt)
}
