package dateutil

import (
	"testing"
	"time"
)

func TestFirstDayOfWeek(t *testing.T) {
	tests := []struct {
		input    time.Time
		expected time.Time
	}{
		{time.Date(2025, 5, 20, 0, 0, 0, 0, time.UTC), time.Date(2025, 5, 18, 0, 0, 0, 0, time.UTC)}, // Tue -> Sun
		{time.Date(2025, 5, 18, 0, 0, 0, 0, time.UTC), time.Date(2025, 5, 18, 0, 0, 0, 0, time.UTC)}, // Sun -> Sun
		{time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC), time.Date(2025, 5, 18, 0, 0, 0, 0, time.UTC)}, // Sat -> Sun
	}

	for _, tt := range tests {
		got := FirstDayOfWeek(tt.input)
		if !got.Equal(tt.expected) {
			t.Errorf("FirstDayOfWeek(%v): expected %v, got %v", tt.input.Format("2006-01-02"), tt.expected.Format("2006-01-02"), got.Format("2006-01-02"))
		}
	}
}
