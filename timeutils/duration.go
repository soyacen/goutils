package timeutils

import (
	"strconv"
	"time"
)

func ParseDurationFromIntString(str string) (time.Duration, error) {
	duration, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(duration), nil
}
