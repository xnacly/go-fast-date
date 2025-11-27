package gofastdate

import "time"

// totime exists for gofastdate->time interop

func FromTime(t time.Time) Date {
	return FromUnix(t.Unix())
}

func (date Date) Unix() int64 {
	y := int64(date.year)
	m := int64(date.month)
	d := int64(date.day)

	if m <= 2 {
		y--
		m += 12
	}

	days := 365*y + y/4 - y/100 + y/400
	days += (153*(m-3) + 2) / 5
	days += d - 1

	const unixEpochOffset = 719163
	daysSinceEpoch := days - unixEpochOffset

	return daysSinceEpoch * 86400
}

func (d Date) Time() time.Time {
	return time.Date(int(d.year), time.Month(d.month), int(d.day), 0, 0, 0, 0, time.UTC)
}
