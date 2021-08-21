package randomutils

import (
	"math/rand"
	"strconv"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// NumericPermString Generate a random number sequence of a given length
func NumericPermString(length int) string {
	var bs []byte
	for i := 0; i < length; i++ {
		bs = strconv.AppendInt(bs, int64(rand.Intn(10)), 10)
	}
	return string(bs)
}
