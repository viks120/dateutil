package dateutil

import "time"

func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func YesterdayString() string {
	return Yesterday().Format("2006-01-02")
}
