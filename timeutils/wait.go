package timeutils

import "time"

// WaitTill waits until toTime
// Params:
//     - toTime: time to wait until. the number of seconds elapsed since January 1, 1970 UTC.
func WaitTill(toTime time.Time) {
	waitDuration := toTime.Sub(time.Now())
	if waitDuration > 0 {
		time.Sleep(waitDuration)
	}
}
