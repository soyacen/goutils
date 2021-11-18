package sortutils_test

import (
	"testing"

	"github.com/soyacen/goutils/sortutils"
)

func TestIntsDesc(t *testing.T) {
	arr := []int{4, 2, 6, 3, 8, 9, 1}
	sortutils.IntsDesc(arr)
	t.Log(arr)
}

func TestIntsAsc(t *testing.T) {
	arr := []int{4, 2, 6, 3, 8, 9, 1}
	sortutils.IntsAsc(arr)
	t.Log(arr)
}
