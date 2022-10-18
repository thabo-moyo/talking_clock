package utils

import (
	"testing"
	"time"
)

func TestGetHumanFriendlyTime(t *testing.T) {
	currentTime, _ := time.Parse("15:04", "15:00")

	if "Three o'clock" != GetHumanFriendlyTime(currentTime) {
		t.Fail()
	}
}
