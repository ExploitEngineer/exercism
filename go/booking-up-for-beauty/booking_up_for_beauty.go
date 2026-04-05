package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
	}
	currentTime := time.Now()
	value := t.Before(currentTime)
	return value
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
	}

	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layoutIn := "1/2/2006 15:04:05"
	t, err := time.Parse(layoutIn, date)
	if err != nil {
		fmt.Println("Error parsing date: ", err)
	}

	// Format the time into the desired output string
	output := t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	return output
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
