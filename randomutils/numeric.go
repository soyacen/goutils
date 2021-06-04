package randomutils

import (
	"math/rand"
	"time"

	"github.com/soyacen/goutils/sliceutils"
)

// NumericPermutation Generate a random number sequence of a given length
func NumericPermutation(length int) string {
	rand.Seed(time.Now().UnixNano())
	return sliceutils.JoinInt(rand.Perm(length), "")
}
