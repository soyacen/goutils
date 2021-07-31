package backoffutils

import (
	"context"
	"math/rand"
	"time"
)

// JitterUp adds random jitter to the duration.
//
// This adds or subtracts time from the duration within a given jitter fraction.
// For example for 10s and jitter 0.1, it will return a time within [9s, 11s])
func JitterUp(backoff BackoffFunc, jitter float64) BackoffFunc {
	return func(ctx context.Context, attempt uint) time.Duration {
		interval := backoff(ctx, attempt)
		return jitterUp(interval, jitter)
	}
}

func jitterUp(duration time.Duration, jitter float64) time.Duration {
	multiplier := jitter * (rand.Float64()*2 - 1)
	return time.Duration(float64(duration) * (1 + multiplier))
}
