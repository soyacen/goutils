package signalutils

import (
	"fmt"
	"os"
)

type SignalError struct {
	Signal os.Signal
}

func (e SignalError) Error() string {
	return fmt.Sprintf("received signal %s", e.Signal)
}
