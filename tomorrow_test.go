package dateutil

import (
	"testing"
	"time"
)

func TestTomorrow(t *testing.T) {
	now := time.Now()
	tom := Tomorrow()
	expected := now.AddDate(0, 0, 1)

	if tom.Year() != expected.Year() || tom.Month() != expected.Month() || tom.Day() != expected.Day() {
		t.Errorf("Expected tomorrow to be %v, got %v", expected.Format("2006-01-02"), tom.Format("2006-01-02"))
	}
}

func TestTomorrowString(t *testing.T) {
	expected := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	result := TomorrowString()

	if result != expected {
		t.Errorf("Expected TomorrowString to return %v, got %v", expected, result)
	}
}
