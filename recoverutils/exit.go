package recoverutils

import (
	"fmt"
	"os"
	"runtime/debug"
)

func ExitOnPanic() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "panic: %s\n", r)
		debug.PrintStack()
		os.Exit(1)
	}
}
