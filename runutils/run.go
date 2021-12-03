package runutils

import (
	"context"
	"log"
	"runtime/debug"
	"time"
)

type Runner struct {
	start       func(ctx context.Context) error
	startCtx    context.Context
	stop        func(ctx context.Context) error
	stopCtx     context.Context
	waitTime    time.Duration
	recoverFunc func(v interface{}, stack []byte)
}

type Option func(r *Runner)

func NewRunner(start func(ctx context.Context) error, opts ...Option) *Runner {
	r := &Runner{
		start:    start,
		startCtx: context.Background(),
		stop: func(ctx context.Context) error {
			return nil
		},
		stopCtx:  context.Background(),
		waitTime: 0,
		recoverFunc: func(v interface{}, stack []byte) {
			log.Printf("panic: %s\n", v)
			log.Println(string(stack))
		},
	}
	for _, option := range opts {
		option(r)
	}
	return r
}

func (r *Runner) Run() error {
	errC := make(chan error, 1)

	go func() {
		if v := recover(); v != nil {
			r.recoverFunc(v, debug.Stack())
		}
		if err := r.start(r.startCtx); err != nil {
			errC <- err
		}
		close(errC)
	}()

	select {
	case <-time.After(r.waitTime):
	case err := <-errC:
		if err != nil {
			return err
		}
	}

	return r.stop(r.stopCtx)
}
