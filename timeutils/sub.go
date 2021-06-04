package timeutils

import "time"

func SubAbsolute(x, y time.Time) time.Duration {
	if x.Before(y) {
		return y.Sub(x)
	}
	return x.Sub(y)
}
