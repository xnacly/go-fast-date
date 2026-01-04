package gofastdate

import "time"

// totime exists for gofastdate->time interop

func FromTime(t time.Time) Date {
	return FromUnix(t.Unix())
}

const (
	s          = 14700
	rata_shift = 719468 + 146097*s + 1
	// The YEAR_SHIFT comprises 14700 * 400 years, as that is the smallest
	// multiple of 400 years which exceeds 231 days.
	year_shift      = 400 * s
	unixEpochOffset = 719163
)

func (date Date) Unix() int64 {
	y := int64(date.year)
	m := int64(date.month)
	d := int64(date.day)

	// as defined in https://www.benjoffe.com/fast-date#inverse
	bump := m <= 2
	y += year_shift
	phase := int64(-2919)
	if bump {
		y--
		phase = 8829
	}
	cent := y / 100
	// overflow prevention
	y_days := y*365 + y/4 - cent + cent/4
	m_days := (979*m + phase) / 32

	days := y_days + m_days + d - rata_shift
	daysSinceEpoch := days - unixEpochOffset
	return daysSinceEpoch * 86400
}

func (d Date) Time() time.Time {
	return time.Date(int(d.year), time.Month(d.month), int(d.day), 0, 0, 0, 0, time.UTC)
}
