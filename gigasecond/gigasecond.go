package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond receives a time, adds 1 gigasecond, and converts the time back to a date.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}

// Do I need the math package for this?
