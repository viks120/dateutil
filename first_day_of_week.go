package dateutil

import "time"

func FirstDayOfWeek(date time.Time) time.Time {
	offset := int(date.Weekday()) // Sunday = 0, Monday = 1, ..., Saturday = 6
	return date.AddDate(0, 0, -offset)
}
