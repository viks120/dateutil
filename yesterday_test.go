package dateutil

import (
	"testing"
	"time"
)

func TestYesterday(t *testing.T) {
	now := time.Now()
	yest := Yesterday()
	expected := now.AddDate(0, 0, -1)

	if yest.Year() != expected.Year() || yest.Month() != expected.Month() || yest.Day() != expected.Day() {
		t.Errorf("Expected yesterday to be %v, got %v", expected.Format("2006-01-02"), yest.Format("2006-01-02"))
	}
}

func TestYesterdayString(t *testing.T) {
	expected := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	result := YesterdayString()

	if result != expected {
		t.Errorf("Expected YesterdayString to return %v, got %v", expected, result)
	}
}
