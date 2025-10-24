package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// handle Teenth case first.
	if wSched == Teenth {
		// start at 13th day of the given month.
		d := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		// increment until got the needed day.
		for d.Weekday() != wDay {
			d = d.AddDate(0, 0, 1)
		}
		return d.Day()
	}

	// for the last day, start at the end of the month.
	if wSched == Last {
		// to do that, go to zeroth day of the next month.
		d := time.Date(year, month + 1, 0, 0, 0, 0, 0, time.UTC)	
		// decrement until get the right day.
		for {
			if d.Weekday() == wDay {
				return d.Day()
			}
			d = d.AddDate(0, 0, -1)
		} 
	}

	// for other weekdays, start with the first day of the month.
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	// keep track of weeks.
	weeks := 0
	for {
		if d.Weekday() == wDay {
			weeks++
		}
		if weeks == int(wSched) {
			return d.Day()
		}
		d = d.AddDate(0, 0, 1)
	}
}
