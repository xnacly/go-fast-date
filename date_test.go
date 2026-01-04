package gofastdate

import (
	"testing"
	"time"
)

var sinkI64 int64
var sinkDate Date

var sampleDates = []Date{
	{1970, 1, 1},
	{1970, 1, 2},
	{1969, 12, 31},
	{2000, 1, 1},
	{1999, 12, 31},
	{2020, 2, 29},
	{2025, 11, 27},
}

var sampleUnix = []int64{
	0,
	86400,
	-86400,
	946684800,
	915148800,
	1582934400,
	1761820800,
}

const stressSize = 1_000_000

var stressUnixTimestamps []int64
var stressDates []Date

func init() {
	start := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2100, 12, 31, 0, 0, 0, 0, time.UTC)

	stressUnixTimestamps = make([]int64, stressSize)
	stressDates = make([]Date, stressSize)

	startUnix := start.Unix()
	endUnix := end.Unix()

	for i := 0; i < stressSize; i++ {
		u := startUnix + int64(i)*(endUnix-startUnix)/stressSize
		stressUnixTimestamps[i] = u

		t := time.Unix(u, 0).UTC()
		stressDates[i] = Date{
			year:  uint64(t.Year()),
			month: uint8(t.Month()),
			day:   uint8(t.Day()),
		}
	}
}

func BenchmarkTimeDateToUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, d := range sampleDates {
			sinkI64 = time.Date(
				int(d.year),
				time.Month(d.month),
				int(d.day),
				0, 0, 0, 0,
				time.UTC,
			).Unix()
		}
	}
}

func BenchmarkGoFastDateToUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, d := range sampleDates {
			sinkI64 = d.Unix()
		}
	}
}

func BenchmarkTimeUnixToDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, u := range sampleUnix {
			t := time.Unix(u, 0).UTC()
			_ = t.Year()
			_ = t.Month()
			_ = t.Day()
		}
	}
}

func BenchmarkGoFastDateFromUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, u := range sampleUnix {
			sinkDate = FromUnix(u)
		}
	}
}

func BenchmarkStressTimeDateToUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, d := range stressDates {
			sinkI64 = time.Date(
				int(d.year),
				time.Month(d.month),
				int(d.day),
				0, 0, 0, 0,
				time.UTC,
			).Unix()
		}
	}
}

func BenchmarkStressGoFastDateToUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, d := range stressDates {
			sinkI64 = d.Unix()
		}
	}
}

func BenchmarkStressTimeUnixToDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, u := range stressUnixTimestamps {
			t := time.Unix(u, 0).UTC()
			_ = t.Year()
			_ = t.Month()
			_ = t.Day()
		}
	}
}

func BenchmarkStressGoFastDateFromUnix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, u := range stressUnixTimestamps {
			sinkDate = FromUnix(u)
		}
	}
}
