package dateutil

import "time"

func GetWeekOfMonth(d time.Time) int {
	firstOfMonth := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
	offset := int(firstOfMonth.Weekday())
	if offset == 0 {
		offset = 7 // Treat Sunday as 7 for ISO-like handling
	}
	day := d.Day()
	return ((day + offset - 2) / 7) + 1
}
