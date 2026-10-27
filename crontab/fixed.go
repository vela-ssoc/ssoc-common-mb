package crontab

import (
	"slices"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

// FixedTime 定点任务，在指定的几个时间点执行。
//
// 例如仅在：
//
//	2025-01-01 00:00:00
//	2025-03-01 00:00:00
//	2025-03-15 00:00:00
//
// 执行三次就完事的任务。
type FixedTime struct {
	points []time.Time
}

func NewFixedTime(points []time.Time) FixedTime {
	points = slices.Clone(points)
	slices.SortFunc(points, func(a, b time.Time) int { return a.Compare(b) })

	return FixedTime{points: points}
}

func (ft FixedTime) Next(t time.Time) time.Time {
	for _, p := range ft.points {
		if p.After(t) {
			return p
		}
	}

	return time.Time{}
}

func NewSecondOptional() cron.Parser {
	op := cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor

	return cron.NewParser(op)
}

func Stringify(s cron.Schedule) string {
	switch v := s.(type) {
	case nil:
		return "<nil>"
	case FixedTime:
		ss := make([]string, 0, len(v.points))
		for _, p := range v.points {
			str := p.Format(time.RFC3339)
			ss = append(ss, str)
		}
		return "@fixed " + strings.Join(ss, " ")
	case cron.ConstantDelaySchedule:
		return "@every " + v.Delay.String()
	}

	return "<unknown>"
}
