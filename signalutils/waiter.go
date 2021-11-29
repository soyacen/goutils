package signalutils

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SignalHook = map[os.Signal]func()

type SignalWaiter struct {
	signals          []os.Signal
	signalC          chan os.Signal
	ctx              context.Context
	cancelFunc       context.CancelFunc
	incomingSignal   os.Signal
	hooks            []func(os.Signal)
	waitUntilTimeout time.Duration
	stopC            chan interface{}
	sync.RWMutex
}

type Option func(*SignalWaiter)

func Signals(signals ...os.Signal) Option {
	return func(waiter *SignalWaiter) {
		waiter.signals = make([]os.Signal, 0)
		for _, signal := range signals {
			waiter.signals = append(waiter.signals, signal)
		}
	}
}

func Context(parent context.Context) Option {
	return func(waiter *SignalWaiter) {
		waiter.ctx, waiter.cancelFunc = context.WithCancel(parent)
	}
}

func Hooks(hooks ...func(os.Signal)) Option {
	return func(waiter *SignalWaiter) {
		for _, hook := range hooks {
			waiter.hooks = append(waiter.hooks, hook)
		}
	}
}

func WaitUntilTimeout(d time.Duration) Option {
	return func(waiter *SignalWaiter) {
		waiter.waitUntilTimeout = d
	}
}

func NewSignalWaiter(opts ...Option) *SignalWaiter {
	ctx, cancelFunc := context.WithCancel(context.Background())
	w := &SignalWaiter{
		signals:          ShutdownSignals(),
		signalC:          make(chan os.Signal),
		ctx:              ctx,
		cancelFunc:       cancelFunc,
		incomingSignal:   nil,
		hooks:            make([]func(os.Signal), 0),
		waitUntilTimeout: 30 * time.Second,
		stopC:            make(chan interface{}, 1),
		RWMutex:          sync.RWMutex{},
	}
	for _, opt := range opts {
		opt(w)
	}
	signal.Notify(w.signalC, w.signals...)
	return w
}

func (w *SignalWaiter) AddHook(f func(os.Signal)) {
	w.Lock()
	defer w.Unlock()
	w.hooks = append(w.hooks, f)
}

func (w *SignalWaiter) Kill(signum syscall.Signal) error {
	return syscall.Kill(syscall.Getpid(), signum)
}

func (w *SignalWaiter) WaitSignals() *SignalWaiter {
	w.incomingSignal = <-w.signalC
	w.cancelFunc()
	return w
}

func (w *SignalWaiter) AsyncInvokeHooks() *SignalWaiter {
	go func(sig os.Signal) {
		w.RLock()
		defer w.RUnlock()
		defer close(w.stopC)
		for i := range w.hooks {
			f := w.hooks[len(w.hooks)-1-i]
			f(sig)
		}
	}(w.incomingSignal)
	return w
}

func (w *SignalWaiter) WaitUntilTimeout() {
	select {
	case <-w.stopC:
		return
	case w.incomingSignal = <-w.signalC:
		return
	case <-time.After(w.waitUntilTimeout):
		return
	}
}

// Context return context that cancels on the waiter's signals.
func (w *SignalWaiter) Context() context.Context {
	return w.ctx
}
