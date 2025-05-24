package dateutil

import "time"

func GetDaysInQuarter() int {
	now := time.Now()
	year := now.Year()
	month := now.Month()

	quarterStartMonth := time.Month(((int(month)-1)/3)*3 + 1)
	firstDayOfQuarter := time.Date(year, quarterStartMonth, 1, 0, 0, 0, 0, time.UTC)

	firstDayNextQuarter := firstDayOfQuarter.AddDate(0, 3, 0)
	lastDayOfQuarter := firstDayNextQuarter.AddDate(0, 0, -1)

	return int(lastDayOfQuarter.Sub(firstDayOfQuarter).Hours()/24) + 1
}
