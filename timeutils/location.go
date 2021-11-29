package timeutils

import "time"

var BeijingLocation = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))
