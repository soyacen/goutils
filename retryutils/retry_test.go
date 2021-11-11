package retryutils_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/soyacen/goutils/backoffutils"
	"github.com/soyacen/goutils/retryutils"
)

func TestCall(t *testing.T) {
	maxAttempts := 3
	ctx := context.Background()
	method := func(attemptTime int) error {
		fmt.Println(attemptTime)
		if attemptTime < maxAttempts {
			attemptTime++
			return errors.New("mock error")
		}
		return nil
	}
	backoffFunc := backoffutils.Constant(time.Second)
	err := retryutils.Call(ctx, uint(maxAttempts), backoffFunc, method)
	assert.Nil(t, err)
}
