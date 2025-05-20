package dateutil

import (
	"testing"
	"time"
)

func TestGetWeekOfMonth(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected int
	}{
		{time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC), 1},
		{time.Date(2025, 5, 5, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(2025, 5, 12, 0, 0, 0, 0, time.UTC), 3},
		{time.Date(2025, 5, 19, 0, 0, 0, 0, time.UTC), 4},
		{time.Date(2025, 5, 26, 0, 0, 0, 0, time.UTC), 5},
	}

	for _, tt := range tests {
		got := GetWeekOfMonth(tt.date)
		if got != tt.expected {
			t.Errorf("GetWeekOfMonth(%v): expected %d, got %d", tt.date.Format("2006-01-02"), tt.expected, got)
		}
	}
}
