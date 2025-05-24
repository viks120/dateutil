package dateutil

import "time"

func IsWeekend() bool {
	return IsWeekendAt(time.Now())
}

func IsWeekendAt(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
