package retryutils

import (
	"context"
	"time"

	"github.com/soyacen/goutils/backoffutils"
)

func Call(ctx context.Context, maxAttempts uint, backoffFunc backoffutils.BackoffFunc, method func() error) error {
	var err error
	max := int(maxAttempts)
	for i := 0; i <= max; i++ {
		// call method
		err = method()

		// if method not return error, no need to retry
		if err == nil {
			break
		}

		// If the maximum number of attempts is exceeded, no need to retry
		if i >= max {
			break
		}

		// sleep and wait retry
		time.Sleep(backoffFunc(ctx, uint(i+1)))
	}
	return err
}
