package breaker

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ redis.Hook = &Hook{}

const defaultInterval = time.Duration(0) * time.Second
const defaultTimeout = time.Duration(60) * time.Second

func defaultReadyToTrip(counts Counts) bool {
	return counts.ConsecutiveFailures > 5
}

func defaultIsSuccessful(errs []error) bool {
	for _, err := range errs {
		if err == nil {
			continue
		}
		if errors.Is(err, redis.Nil) {
			continue
		}
		return false
	}
	return true
}

func defaultOnStateChange(name string, from State, to State) {}

// State is a type that represents a state of Hook.
type State int

// These constants are states of Hook.
const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

var (
	// ErrTooManyRequests is returned when the CB state is half open and the requests count is over the cb maxRequests
	ErrTooManyRequests = errors.New("too many requests")
	// ErrOpenState is returned when the CB state is open
	ErrOpenState = errors.New("circuit breaker is open")
)

// String implements stringer interface.
func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateHalfOpen:
		return "half-open"
	case StateOpen:
		return "open"
	default:
		return fmt.Sprintf("unknown state: %d", s)
	}
}

// Counts holds the numbers of requests and their successes/failures.
// Hook clears the internal Counts either
// on the change of the state or at the closed-state intervals.
// Counts ignores the results of the requests sent before clearing.
type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

func (c *Counts) onRequest() {
	c.Requests++
}

func (c *Counts) onSuccess() {
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

func (c *Counts) onFailure() {
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

func (c *Counts) clear() {
	c.Requests = 0
	c.TotalSuccesses = 0
	c.TotalFailures = 0
	c.ConsecutiveSuccesses = 0
	c.ConsecutiveFailures = 0
}

type Hook struct {
	name          string
	maxRequests   uint32
	interval      time.Duration
	timeout       time.Duration
	readyToTrip   func(counts Counts) bool
	isSuccessful  func(errs []error) bool
	onStateChange func(name string, from State, to State)

	mutex      sync.Mutex
	state      State
	generation uint64
	counts     Counts
	expiry     time.Time
}

type Option func(hook *Hook)

// Name is the name of the Hook.
func Name(name string) Option {
	return func(hook *Hook) {
		hook.name = name
	}
}

// MaxRequests is the maximum number of requests allowed to pass through
// when the Hook is half-open.
// If MaxRequests is 0, the Hook allows only 1 request.
func MaxRequests(maxRequests uint32) Option {
	return func(hook *Hook) {
		hook.maxRequests = maxRequests
	}
}

// Interval is the cyclic period of the closed state
// for the Hook to clear the internal Counts.
// If Interval is less than or equal to 0, the Hook doesn't clear internal Counts during the closed state.
func Interval(interval time.Duration) Option {
	return func(hook *Hook) {
		hook.interval = interval
	}
}

// Timeout is the period of the open state,
// after which the state of the Hook becomes half-open.
// If Timeout is less than or equal to 0, the timeout value of the Hook is set to 60 seconds.
func Timeout(timeout time.Duration) Option {
	return func(hook *Hook) {
		hook.timeout = timeout
	}
}

// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
// If ReadyToTrip returns true, the Hook will be placed into the open state.
// If ReadyToTrip is nil, default ReadyToTrip is used.
// Default ReadyToTrip returns true when the number of consecutive failures is more than 5.
func ReadyToTrip(fn func(counts Counts) bool) Option {
	return func(hook *Hook) {
		hook.readyToTrip = fn
	}
}

// OnStateChange is called whenever the state of the Hook changes.
func OnStateChange(fn func(name string, from State, to State)) Option {
	return func(hook *Hook) {
		hook.onStateChange = fn
	}
}

// IsSuccessful is called with the error returned from a request.
// If IsSuccessful returns true, the error is counted as a success.
// Otherwise, the error is counted as a failure.
// If IsSuccessful is nil, default IsSuccessful is used, which returns false for all non-nil errors.
func IsSuccessful(fn func(errs []error) bool) Option {
	return func(hook *Hook) {
		hook.isSuccessful = fn
	}
}

// New returns a new Hook configured with the given Settings.
func New(opts ...Option) *Hook {
	hook := &Hook{
		name:          "redis.breaker",
		maxRequests:   1,
		interval:      defaultInterval,
		timeout:       defaultTimeout,
		readyToTrip:   defaultReadyToTrip,
		isSuccessful:  defaultIsSuccessful,
		onStateChange: defaultOnStateChange,
		mutex:         sync.Mutex{},
		state:         0,
		generation:    0,
		counts:        Counts{},
		expiry:        time.Time{},
	}
	for _, opt := range opts {
		opt(hook)
	}

	hook.toNewGeneration(time.Now())
	return hook
}

type generationKey struct{}

func (hook *Hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	generation, err := hook.beforeRequest()
	ctx = context.WithValue(ctx, generationKey{}, generation)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (hook *Hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	generation, ok := ctx.Value(generationKey{}).(uint64)
	if !ok {
		return errors.New("generation not found")
	}
	var errs []error
	if cmd.Err() != nil {
		errs = append(errs, cmd.Err())
	}
	hook.afterRequest(generation, hook.isSuccessful(errs))
	return nil
}

func (hook *Hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	generation, err := hook.beforeRequest()
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, generationKey{}, generation)
	return ctx, nil
}

func (hook *Hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	generation, ok := ctx.Value(generationKey{}).(uint64)
	if !ok {
		return errors.New("generation not found")
	}
	var errs []error
	for _, cmd := range cmds {
		if cmd.Err() != nil {
			errs = append(errs, cmd.Err())
		}
	}
	hook.afterRequest(generation, hook.isSuccessful(errs))
	return nil
}

func (hook *Hook) beforeRequest() (uint64, error) {
	hook.mutex.Lock()
	defer hook.mutex.Unlock()

	now := time.Now()
	state, generation := hook.currentState(now)

	if state == StateOpen {
		return generation, ErrOpenState
	} else if state == StateHalfOpen && hook.counts.Requests >= hook.maxRequests {
		return generation, ErrTooManyRequests
	}

	hook.counts.onRequest()
	return generation, nil
}

func (hook *Hook) afterRequest(before uint64, success bool) {
	hook.mutex.Lock()
	defer hook.mutex.Unlock()

	now := time.Now()
	state, generation := hook.currentState(now)
	if generation != before {
		return
	}

	if success {
		hook.onSuccess(state, now)
	} else {
		hook.onFailure(state, now)
	}
}

func (hook *Hook) onSuccess(state State, now time.Time) {
	switch state {
	case StateClosed:
		hook.counts.onSuccess()
	case StateHalfOpen:
		hook.counts.onSuccess()
		if hook.counts.ConsecutiveSuccesses >= hook.maxRequests {
			hook.setState(StateClosed, now)
		}
	}
}

func (hook *Hook) onFailure(state State, now time.Time) {
	switch state {
	case StateClosed:
		hook.counts.onFailure()
		if hook.readyToTrip(hook.counts) {
			hook.setState(StateOpen, now)
		}
	case StateHalfOpen:
		hook.setState(StateOpen, now)
	}
}

func (hook *Hook) currentState(now time.Time) (State, uint64) {
	switch hook.state {
	case StateClosed:
		if !hook.expiry.IsZero() && hook.expiry.Before(now) {
			hook.toNewGeneration(now)
		}
	case StateOpen:
		if hook.expiry.Before(now) {
			hook.setState(StateHalfOpen, now)
		}
	}
	return hook.state, hook.generation
}

func (hook *Hook) setState(state State, now time.Time) {
	if hook.state == state {
		return
	}

	prev := hook.state
	hook.state = state

	hook.toNewGeneration(now)

	hook.onStateChange(hook.name, prev, state)
}

func (hook *Hook) toNewGeneration(now time.Time) {
	hook.generation++
	hook.counts.clear()

	var zero time.Time
	switch hook.state {
	case StateClosed:
		if hook.interval == 0 {
			hook.expiry = zero
		} else {
			hook.expiry = now.Add(hook.interval)
		}
	case StateOpen:
		hook.expiry = now.Add(hook.timeout)
	default: // StateHalfOpen
		hook.expiry = zero
	}
}
