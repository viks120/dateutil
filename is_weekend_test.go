package dateutil

import (
	"testing"
	"time"
)

func TestIsWeekendAt(t *testing.T) {
	sat := time.Date(2025, 5, 17, 0, 0, 0, 0, time.UTC) // Saturday
	sun := time.Date(2025, 5, 18, 0, 0, 0, 0, time.UTC) // Sunday
	mon := time.Date(2025, 5, 19, 0, 0, 0, 0, time.UTC) // Monday

	if !IsWeekendAt(sat) {
		t.Error("Expected Saturday to be a weekend")
	}
	if !IsWeekendAt(sun) {
		t.Error("Expected Sunday to be a weekend")
	}
	if IsWeekendAt(mon) {
		t.Error("Expected Monday to not be a weekend")
	}
}

func TestIsWeekend(t *testing.T) {
	_ = IsWeekend() // Just confirm it runs without panic; logic is tested in IsWeekendAt
}
