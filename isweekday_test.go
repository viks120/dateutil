package dateutil

import (
	"testing"
	"time"
)

func TestIsWeekDayAt(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected bool
	}{
		{time.Date(2025, 5, 19, 0, 0, 0, 0, time.UTC), true},  // Monday
		{time.Date(2025, 5, 20, 0, 0, 0, 0, time.UTC), true},  // Tuesday
		{time.Date(2025, 5, 21, 0, 0, 0, 0, time.UTC), true},  // Wednesday
		{time.Date(2025, 5, 22, 0, 0, 0, 0, time.UTC), true},  // Thursday
		{time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC), true},  // Friday
		{time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC), false}, // Saturday
		{time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC), false}, // Sunday
	}

	for _, tt := range tests {
		got := IsWeekDayAt(tt.date)
		if got != tt.expected {
			t.Errorf("IsWeekDayAt(%v): expected %v, got %v", tt.date.Weekday(), tt.expected, got)
		}
	}
}

func TestIsWeekDay(t *testing.T) {
	_ = IsWeekDay() // Ensure no panic
}
