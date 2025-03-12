package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/02/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return "You have an appointment on " + t.Weekday().String() + ", " + t.Month().String() + " " + t.Format("2") + ", " + t.Format("2006") + ", at " + t.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t, _ := time.Parse("1/02/2006 15:04:05", "9/15/"+time.Now().Format("2006")+" 00:00:00")
	return t
	panic("Please implement the AnniversaryDate function")
}
