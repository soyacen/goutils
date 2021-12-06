package runutils

import (
	"log"
	"runtime/debug"
	"time"
)

type Runner struct {
	execute  func() error
	waitTime time.Duration
	recover  func(v interface{})
}

type Option func(r *Runner)

func Recover(method func(v interface{})) Option {
	return func(r *Runner) {
		r.recover = method
	}
}

func NewRunner(execute func() error, waitTime time.Duration, opts ...Option) *Runner {
	r := &Runner{
		execute:  execute,
		waitTime: waitTime,
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
		if err := r.execute(); err != nil {
			errC <- err
		}
		close(errC)
	}()

	var err error
	select {
	case <-time.After(r.waitTime):
	case err = <-errC:
	}
	return err
}
