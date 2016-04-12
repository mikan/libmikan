// Copyright 2016 mikan.

package data

import (
	"testing"
	"time"
)

func TestMinuteAdd(t *testing.T) {
	var counter MinuteHourCounter
	counter.Add(1)
	check(t, 1, counter.MinuteCount())
	check(t, 1, counter.HourCount())
	counter.Add(10)
	check(t, 11, counter.MinuteCount())
	check(t, 11, counter.HourCount())
	timeShift(&counter, 2)
	check(t, 0, counter.MinuteCount())
	check(t, 11, counter.HourCount())
	counter.Add(1)
	check(t, 1, counter.MinuteCount())
	check(t, 12, counter.HourCount())
	timeShift(&counter, 61)
	check(t, 0, counter.MinuteCount())
	check(t, 0, counter.HourCount())
}

func check(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func timeShift(c *MinuteHourCounter, minutes int) {
	// CAUTION: Access to implementatation details
	shift := -time.Duration(minutes) * time.Minute
	for e := c.minuteEvents.Front(); e != nil; e = e.Next() {
		e.Value.(*event).t = e.Value.(*event).t.Add(shift)
	}
	for e := c.hourEvents.Front(); e != nil; e = e.Next() {
		e.Value.(*event).t = e.Value.(*event).t.Add(shift)
	}
}
