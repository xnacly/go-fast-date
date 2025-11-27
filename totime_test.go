package gofastdate

import (
	"testing"
	"time"
)

func TestDateUnixAndTime(t *testing.T) {
	tests := []struct {
		year, month, day uint64
	}{
		{1970, 1, 1},
		{1970, 1, 2},
		{1969, 12, 31},
		{2000, 1, 1},
		{1999, 12, 31},
		{2020, 2, 29},
		{2025, 11, 27},
	}

	for _, tt := range tests {
		t.Run(
			time.Date(int(tt.year), time.Month(tt.month), int(tt.day), 0, 0, 0, 0, time.UTC).Format("2006-01-02"),
			func(t *testing.T) {
				d := Date{day: tt.day, month: tt.month, year: tt.year}
				gotTime := d.Time()
				expected := time.Date(int(tt.year), time.Month(tt.month), int(tt.day), 0, 0, 0, 0, time.UTC)
				if !gotTime.Equal(expected) {
					t.Errorf("Date.Time() = %v; want %v", gotTime, expected)
				}

				fromTime := FromTime(expected)
				if fromTime != d {
					t.Errorf("FromTime round-trip failed: got %+v, want %+v", fromTime, d)
				}
			},
		)
	}
}

func TestFromUnixToDate(t *testing.T) {
	unixTests := []int64{
		0,
		86400,
		-86400,
		946684800,
		915148800,
		1582934400,
		1761820800,
	}

	for _, u := range unixTests {
		t.Run(time.Unix(u, 0).UTC().Format("2006-01-02"), func(t *testing.T) {
			d := FromUnix(u)
			t0 := time.Unix(u, 0).UTC()
			if d.year != uint64(t0.Year()) || d.month != uint64(t0.Month()) || d.day != uint64(t0.Day()) {
				t.Errorf("FromUnix(%d) -> Date %+v; want Y/M/D %04d/%02d/%02d",
					u, d, t0.Year(), t0.Month(), t0.Day())
			}
		})
	}
}
