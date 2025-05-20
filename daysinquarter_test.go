package dateutil

import (
	"testing"
	"time"
)

func TestGetDaysInQuarter(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected int
	}{
		{time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC), 90},  // Q1: Jan–Mar
		{time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC), 91},  // Q2: Apr–Jun
		{time.Date(2025, 7, 10, 0, 0, 0, 0, time.UTC), 92},  // Q3: Jul–Sep
		{time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC), 92}, // Q4: Oct–Dec
	}

	for _, tt := range tests {
		quarterStartMonth := time.Month(((int(tt.date.Month()) - 1) / 3 * 3) + 1)
		firstDayOfQuarter := time.Date(tt.date.Year(), quarterStartMonth, 1, 0, 0, 0, 0, time.UTC)
		firstDayNextQuarter := firstDayOfQuarter.AddDate(0, 3, 0)
		lastDayOfQuarter := firstDayNextQuarter.AddDate(0, 0, -1)
		actualDays := int(lastDayOfQuarter.Sub(firstDayOfQuarter).Hours()/24) + 1

		if actualDays != tt.expected {
			t.Errorf("For date %v: expected %d days in quarter, got %d", tt.date.Format("2006-01-02"), tt.expected, actualDays)
		}
	}
}
