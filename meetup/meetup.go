package meetup

import (
	"fmt"
	"strconv"
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

// https://golangbyexample.com/iota-in-golang/
const (
    First = iota
	// Find first occurrence after firstOfMonth
	Second
	// First occurence + 7
	Third
	// First occurence + 14
	Fourth 
	// First occurence + 21 (must exist since we don't return errors)
	Teenth
	// Between 13th to 19th only (inclusive)
	Last 
	// Start at lastOfMonth and work backwards
)

func incrementDate(start, end int, date time.Time, weekday time.Weekday, increment int) time.Time {
	for i := start; i <= end; i++ {
		if date.Weekday() != weekday {
			date = date.AddDate(0, 0, increment)
		} else {
			break
		}
	}
	return date
}

func findDate(firstOfMonth, lastOfMonth time.Time, week WeekSchedule, weekday time.Weekday) int {
	FirstOccurence:= incrementDate(1, 7, firstOfMonth, weekday, 1)
	switch week {
	case First:
		return FirstOccurence.Day()
	case Second:
		return FirstOccurence.AddDate(0, 0, 7).Day()
	case Third:
		return FirstOccurence.AddDate(0, 0, 14).Day()
	case Fourth:
		return FirstOccurence.AddDate(0, 0, 21).Day()
	case Teenth:
		return incrementDate(1, 7, firstOfMonth.AddDate(0, 0, 12), weekday, 1).Day()
	case Last:
		return incrementDate(1, 7, lastOfMonth, weekday, -1).Day()
	}
	return -1 // Stub
}

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	// Calculate first and last of month
	firstOfMonth, err := time.Parse("January-02-2006", fmt.Sprintf("%s-01-%s", month, strconv.Itoa(year)))
	if err != nil {
		panic(err)
	}
    lastOfMonth := firstOfMonth.AddDate(0, 1, -1) // + 1 month, -1 day
	// Return desired date
	return findDate(firstOfMonth, lastOfMonth, week, weekday)
}
