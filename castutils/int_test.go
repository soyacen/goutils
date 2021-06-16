package castutils

import "testing"

func TestToInt(t *testing.T) {
	res := ToInt(1.3)
	t.Log(res)
}
