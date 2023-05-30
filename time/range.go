package time

import (
	"errors"
	"time"
)

type Range struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (r Range) PeriodsOverlap(t Range) bool {
	return periodsOverlap(r, t)
}

// PeriodsOverlap 判断两个时间段是否有重叠
func PeriodsOverlap(st Range, et Range) bool {
	return periodsOverlap(st, et)
}

func periodsOverlap(st Range, et Range) bool {
	if et.StartTime.Before(st.EndTime) && st.StartTime.Before(et.EndTime) {
		return true
	}
	return false
}

func NewRange(st time.Time, et time.Time) (Range, error) {
	if !st.Before(et) {
		return Range{}, errors.New("time range error")
	}
	return Range{StartTime: st, EndTime: et}, nil
}

func NewRangeFromString(st, et, layout string) (_ Range, err error) {
	var t1 time.Time
	var t2 time.Time
	t1, err = time.Parse(layout, st)
	if err != nil {
		return
	}
	t2, err = time.Parse(layout, et)
	if err != nil {
		return
	}
	return NewRange(t1, t2)
}
