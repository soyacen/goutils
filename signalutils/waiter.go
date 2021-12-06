package signalutils

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SignalHook = map[os.Signal]func()

type SignalWaiter struct {
	signals        []os.Signal
	signalC        chan os.Signal
	incomingSignal os.Signal
	hooks          []func(os.Signal)
	waitTimeout    time.Duration
	stopC          chan interface{}
	sync.RWMutex
}

func NewSignalWaiter(signals []os.Signal, waitTimeout time.Duration) *SignalWaiter {
	w := &SignalWaiter{
		signals:        signals,
		signalC:        make(chan os.Signal),
		incomingSignal: nil,
		hooks:          make([]func(os.Signal), 0),
		waitTimeout:    waitTimeout,
		stopC:          make(chan interface{}, 1),
		RWMutex:        sync.RWMutex{},
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
	return w
}

func (w *SignalWaiter) WaitHooksAsyncInvoked() *SignalWaiter {
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
	case <-time.After(w.waitTimeout):
		return
	}
}
