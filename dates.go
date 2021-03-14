package main

import "time"

var WDays = []string{
	"", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
}

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func Today() time.Time {
	return StartOfDay(time.Now().Local())
}

func Tomorrow() time.Time {
	return Today().Add(24 * time.Hour)
}

// Weekday adjusts time.Weekday() to have Monday be 1 nd Sunday 7.
func Weekday(t time.Time) int {
	d := t.Weekday()
	if d == 0 {
		d = 7
	}
	return int(d)
}

func TodayWeekday() int {
	return Weekday(time.Now())
}

func TodayDate() string {
	return Date(time.Now())
}

func TodayKey() int {
	return TodayWeekday() * 100
}

func Date(t time.Time) string {
	return t.Format("2006-01-02")
}
