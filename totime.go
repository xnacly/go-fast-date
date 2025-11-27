package gofastdate

import "time"

// totime exists for gofastdate->time interop

func FromTime(t time.Time) Date {
	return FromUnix(t.Unix())
}

func toRataDie(year, month, day uint64) int32 {
	y := int64(year)
	m := int64(month)
	d := int64(day)

	// shift Jan/Feb to end of previous year
	if m <= 2 {
		y--
		m += 12
	}

	// compute days from years, including leap years
	days := 365*y + y/4 - y/100 + y/400

	// days from months
	days += (153 * (m + 1)) / 5

	days += d

	// subtract epoch offset (1970‑01‑01)
	return int32(days - 719468) // 719468 aligns 1970-01-01 = 0
}

func (d Date) ToTime() time.Time {
	days := toRataDie(d.year, d.month, d.day)
	seconds := int64(days) * 86400
	return time.Unix(seconds, 0).UTC()
}
