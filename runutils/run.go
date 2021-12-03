package runutils

import (
	"context"
	"log"
	"runtime/debug"
	"time"
)

type Runner struct {
	start    func(ctx context.Context) error
	startCtx context.Context
	stop     func(ctx context.Context) error
	stopCtx  context.Context
	waitTime time.Duration
	recover  func(v interface{})
}

type Option func(r *Runner)

func StartCtx(ctx context.Context) Option {
	return func(r *Runner) {
		r.startCtx = ctx
	}
}

func Stop(method func(ctx context.Context) error) Option {
	return func(r *Runner) {
		r.stop = method
	}
}

func StopCtx(ctx context.Context) Option {
	return func(r *Runner) {
		r.stopCtx = ctx
	}
}

func WaitTime(t time.Duration) Option {
	return func(r *Runner) {
		r.waitTime = t
	}
}

func Recover(method func(v interface{})) Option {
	return func(r *Runner) {
		r.recover = method
	}
}

func NewRunner(start func(ctx context.Context) error, opts ...Option) *Runner {
	r := &Runner{
		start:    start,
		startCtx: context.Background(),
		stop: func(ctx context.Context) error {
			return nil
		},
		stopCtx:  context.Background(),
		waitTime: 0,
		recover: func(v interface{}) {
			log.Printf("panic: %s\n", v)
			debug.PrintStack()
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
		defer func() {
			if v := recover(); v != nil {
				r.recover(v)
			}
		}()
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
