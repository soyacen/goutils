package errorutils

// PanicHandler is called for recovering from panics spawned internally to the library (and thus
// not recoverable by the caller's goroutine). Defaults to nil, which means panics are not recovered.
var PanicHandler func(interface{})

func WithRecover(fn func()) {
	defer func() {
		handler := PanicHandler
		if handler != nil {
			if err := recover(); err != nil {
				handler(err)
			}
		}
	}()
	fn()
}
