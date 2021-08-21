package randomutils

import (
	"strconv"
)

func Intn(n int) int {
	return r.Intn(n)
}

// NumericPermString Generate a random number sequence of a given length
func NumericPermString(length int) string {
	var bs []byte
	for i := 0; i < length; i++ {
		bs = strconv.AppendInt(bs, int64(Intn(10)), 10)
	}
	return string(bs)
}

func PickInt32(n ...int32) int32 {
	return n[Intn(len(n))]
}
