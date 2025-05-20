package dateutil

import "time"

func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

func TomorrowString() string {
	return Tomorrow().Format("2006-01-02")
}
