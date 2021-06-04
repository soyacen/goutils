package timeutils

import "time"

func TimeToNanosecond(t time.Time) int64 {
	return t.UnixNano()
}

func TimeToMicrosecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

func TimeToMillisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func TimeToSecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Second)
}

func TimeToMinute(t time.Time) int64 {
	return t.UnixNano() / int64(time.Minute)
}

func TimeToHour(t time.Time) int64 {
	return t.UnixNano() / int64(time.Hour)
}
