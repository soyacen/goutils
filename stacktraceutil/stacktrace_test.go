package stacktraceutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTakeStacktrace(t *testing.T) {
	trace := TakeStacktrace(0)
	lines := strings.Split(trace, "\n")
	require.NotEmpty(t, lines, "Expected stacktrace to have at least one frame.")
	assert.Contains(
		t,
		lines[0],
		"github.com/soyacen/goutils/stacktraceutil.TestTakeStacktrace",
		"Expected stacktrace to start with the test.",
	)
}

func TestTakeStacktraceWithSkip(t *testing.T) {
	trace := TakeStacktrace(1)
	lines := strings.Split(trace, "\n")
	require.NotEmpty(t, lines, "Expected stacktrace to have at least one frame.")
	assert.Contains(
		t,
		lines[0],
		"testing.",
		"Expected stacktrace to start with the test runner (skipping our own frame).",
	)
}

func TestTakeStacktraceWithSkipInnerFunc(t *testing.T) {
	var trace string
	func() {
		trace = TakeStacktrace(2)
	}()
	lines := strings.Split(trace, "\n")
	require.NotEmpty(t, lines, "Expected stacktrace to have at least one frame.")
	assert.Contains(
		t,
		lines[0],
		"testing.",
		"Expected stacktrace to start with the test function (skipping the test function).",
	)
}

func BenchmarkTakeStacktrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TakeStacktrace(0)
	}
}
