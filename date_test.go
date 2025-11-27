package gofastdate

import (
	"testing"
	"time"
)

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

func BenchmarkTimeUnix(b *testing.B) {
	for b.Loop() {
		for _, u := range sampleUnix {
			_ = time.Unix(u, 0).UTC().Unix()
		}
	}
}

func BenchmarkGoFastDateUnix(b *testing.B) {
	for b.Loop() {
		for _, d := range sampleDates {
			_ = d.Unix()
		}
	}
}

const stressSize = 1_000_000

var stressUnixTimestamps []int64

func init() {
	start := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(2100, 12, 31, 0, 0, 0, 0, time.UTC).Unix()
	stressUnixTimestamps = make([]int64, stressSize)

	for i := 0; i < stressSize; i++ {
		stressUnixTimestamps[i] = start + int64(i)*(end-start)/stressSize
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
			_ = FromUnix(u)
		}
	}
}
