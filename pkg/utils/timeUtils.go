package utils

import (
	"fmt"
	"github.com/divan/num2words"
	"strings"
	"time"
)

func GetCurrentTime() string {
	return GetHumanFriendlyTime(time.Now())
}
func GetTimeByString(parse string) (string, error) {
	value, err := parseTimeString(parse)
	return value, err
}

func parseTimeString(timeString ...string) (string, error) {
	parse, err := time.Parse("15:04", timeString[0])

	return GetHumanFriendlyTime(parse), err
}

// Convert given time to human friendly text.
func GetHumanFriendlyTime(timeInput time.Time) string {
	hour := convertTo12Hour(timeInput.Hour())
	minute := timeInput.Minute()
	timeString := ""

	if hour == 0 {
		hour = 12
	}

	switch minute {
	case 00:
		timeString = fmt.Sprint(strings.Title(num2words.Convert(hour)), " o'clock")
	case 15:
		timeString = fmt.Sprint("Quarter past ", num2words.Convert(hour))
	case 30:
		timeString = fmt.Sprint("Half past ", num2words.Convert(hour))
	case 45:
		hour = convertTo12Hour(addHour(timeInput))
		timeString = fmt.Sprint("Quarter to ", num2words.Convert(hour))
	default:
		meridian := ""
		if minute < 30 {
			meridian = " past "
		} else {
			hour = convertTo12Hour(addHour(timeInput))
			meridian = " to "
			minute = 60 - minute
		}

		timeString = fmt.Sprint(
			strings.Title(num2words.Convert(minute)),
			meridian,
			num2words.Convert(hour),
		)
	}
	return timeString
}

// Adds an hour to a time.Time object
func addHour(timeInput time.Time) int {
	return timeInput.Add(time.Hour * 1).Hour()
}

// converts time from 24H to 12H format
func convertTo12Hour(hour int) int {
	if hour >= 13 {
		hour = hour - 12
	}
	return hour
}
