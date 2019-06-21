package gigasecond

import (
	"time"
)

// AddGigasecond receives a time, adds 1 gigasecond, and converts the time back to a date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
