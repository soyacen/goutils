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
	signalC          chan os.Signal
	signals          []os.Signal
	incomingSignal   os.Signal
	hooks            []func(os.Signal)
	stopC            chan interface{}
	waitUntilTimeout time.Duration
	sync.RWMutex
}

type Option func(*SignalWaiter)

func Signals(signals ...os.Signal) Option {
	return func(helper *SignalWaiter) {
		helper.signals = make([]os.Signal, 0)
		for _, signal := range signals {
			helper.signals = append(helper.signals, signal)
		}
	}
}

func Hooks(hooks ...func(os.Signal)) Option {
	return func(helper *SignalWaiter) {
		for _, hook := range hooks {
			helper.hooks = append(helper.hooks, hook)
		}
	}
}

func WaitUntilTimeout(d time.Duration) Option {
	return func(helper *SignalWaiter) {
		helper.waitUntilTimeout = d
	}
}

func NewSignalWaiter(opts ...Option) *SignalWaiter {
	w := &SignalWaiter{
		signalC:          make(chan os.Signal),
		signals:          ShutdownSignals(),
		hooks:            make([]func(os.Signal), 0),
		stopC:            make(chan interface{}, 1),
		waitUntilTimeout: 30 * time.Second,
	}
	for _, opt := range opts {
		opt(w)
	}
	w.signals = append(w.signals, syscall.SIGUSR1)
	signal.Notify(w.signalC, w.signals...)
	return w
}

func (w *SignalWaiter) AddHook(f func(os.Signal)) {
	w.Lock()
	defer w.Unlock()
	w.hooks = append(w.hooks, f)
}

func (w *SignalWaiter) Leave() {
	// leave
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
}

func (w *SignalWaiter) WaitSignals() *SignalWaiter {
	// wait signals
	w.incomingSignal = <-w.signalC
	return w
}

func (w *SignalWaiter) AsyncInvokeHooks() *SignalWaiter {
	// async invoke hooks
	go func(sig os.Signal) {
		w.RLock()
		defer w.RUnlock()
		defer close(w.stopC)
		for _, f := range w.hooks {
			defer f(sig)
		}
	}(w.incomingSignal)
	return w
}

func (w *SignalWaiter) WaitUntilTimeout() {
	// wait until timeout
	select {
	case <-w.stopC:
		return
	case w.incomingSignal = <-w.signalC:
		return
	case <-time.After(w.waitUntilTimeout):
	}
}

// SignalContext creates a new context  that cancels on the waiter's signals.
func (w *SignalWaiter) SignalContext() (context.Context, context.CancelFunc) {
	return SignalContext(w.signals...)
}

func ShutdownSignals() []os.Signal {
	return []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGUSR1,
	}
}

// SignalContext creates a new context that cancels on the given signals.
func SignalContext(signals ...os.Signal) (context.Context, context.CancelFunc) {
	return ContextWithSignal(context.Background(), signals...)
}

// ContextWithSignal creates a new context that cancels on the given signals.
func ContextWithSignal(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, closer := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)

	go func() {
		select {
		case <-c:
			closer()
		case <-ctx.Done():
		}
	}()

	return ctx, closer
}
