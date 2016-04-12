// Copyright 2016 mikan.

package data

import (
	"container/list"
	"time"
)

type event struct {
	count int
	t     time.Time
}

// MinuteHourCounter counts events by recent minutes and hours.
type MinuteHourCounter struct {
	minuteEvents *list.List
	hourEvents   *list.List // 直近1分間のイベントは含まれて「いない」

	minuteCount int
	hourCount   int // 直近1時間の「すべて」のイベントをカウントしている
}

// Add adds event to counter
func (c *MinuteHourCounter) Add(count int) {
	if c.minuteEvents == nil {
		c.minuteEvents = list.New()
	}
	if c.hourEvents == nil {
		c.hourEvents = list.New()
	}
	nowSecs := time.Now()
	c.shiftOldEvents(nowSecs)

	// 1分間のリストに流し込む (1時間のリストはあとから)
	c.minuteEvents.PushBack(&event{count, nowSecs})
	c.minuteCount += count
	c.hourCount += count
}

// MinuteCount returns number of events by recent a minute.
func (c *MinuteHourCounter) MinuteCount() int {
	c.shiftOldEvents(time.Now())
	return c.minuteCount
}

// HourCount returns number of events by recent a minute.
func (c *MinuteHourCounter) HourCount() int {
	c.shiftOldEvents(time.Now())
	return c.hourCount
}

func (c *MinuteHourCounter) shiftOldEvents(nowSecs time.Time) {
	minuteAgo := nowSecs.Add(-time.Minute)
	hourAgo := nowSecs.Add(-time.Hour)

	// 1分以上経過したイベントを 'minuteEvents' から 'hourEvents' へと移動する。
	for c.minuteEvents.Len() > 0 && c.minuteEvents.Front().Value.(*event).t.Before(minuteAgo) {
		front := c.minuteEvents.Front()
		c.hourEvents.PushBack(front.Value)

		c.minuteCount -= c.minuteEvents.Front().Value.(*event).count
		c.minuteEvents.Remove(front)
	}

	// 1時間以上経過した古いイベントを 'hourevents' から削除する
	for c.hourEvents.Len() > 0 && c.hourEvents.Front().Value.(*event).t.Before(hourAgo) {
		c.hourCount -= c.hourEvents.Front().Value.(*event).count
		c.hourEvents.Remove(c.hourEvents.Front())
	}
}
