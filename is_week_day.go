package dateutil

import "time"

func IsWeekDay() bool {
	return IsWeekDayAt(time.Now())
}

func IsWeekDayAt(t time.Time) bool {
	weekday := t.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}
