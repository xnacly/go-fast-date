# go-fast-date

> 3-6x faster than [`time`](https://pkg.go.dev/time), see [Benchmarks](#Benchmarks)

Go port of _A Very Fast 64–Bit Date Algorithm_, based on [A Very Fast 64–Bit
Date Algorithm: 30–40% faster](https://www.benjoffe.com/fast-date-64) by Ben
Joffe. Also uses Parts of _The Julian Map: A Faster technique for Gregorian
Date Conversion_, see [Improvements to the Inverse
Function](https://www.benjoffe.com/fast-date#inverse). Special thanks to
@benjoffe for writing such easy to follow blog posts and implementations.

## Usage

```shell
go get github.com/xnacly/go-fast-date
```

```go
package main

import (
	"fmt"
	"time"

	"github.com/xnacly/go-fast-date"
)

func main() {
	date := gofastdate.FromUnix(time.Now().Unix())
	fmt.Println(date)

	fmt.Println(date.Time())
	fmt.Println(date.Unix(), date.Time().Unix())
}
```

## Features

`go-fast-date` implements Ben Joffe's very fast 64-bit date
algorithm and provides date round trip functions for
[`time.Time`](https://pkg.go.dev/time) compatibility:

- `gofastdate.FromUnix`
- `gofastdate.FromTime`
- `gofastdate.Date.Time`
- `gofastdate.Date.Unix`

While not depending on any dependency and striving to match
Joffe's implementation, where applicable.

## Benchmarks

| kind   | benchmark                 | `time`       | `go-fast-date` | speedup | runtime reduction |
| ------ | ------------------------- | ------------ | -------------- | ------- | ----------------- |
| micro  | BenchmarkDateToUnix       | 194.6 ns/op  | 63.59 ns/op    | 3.06x   | −67.3%            |
|        | BenchmarkUnixToDate       | 582.9 ns/op  | 90.23 ns/op    | 6.46x   | −84.5%            |
| stress | BenchmarkStressDateToUnix | 8.697 ms/op  | 2.856 ms/op    | 3.05x   | −67.2%            |
|        | BenchmarkStressUnixToDate | 27.316 ms/op | 4.286 ms/op    | 6.37x   | −84.3%            |

Benchmarks focus on comparing `gofastdate.FromUnix` with
`time.Unix`. Run with `go test ./... -bench=.`

```text
goos: linux
goarch: amd64
pkg: github.com/xnacly/go-fast-date
cpu: AMD Ryzen 7 PRO 7840U w/ Radeon 780M Graphics
BenchmarkTimeDateToUnix-16                       6215389               194.6 ns/op             0 B/op          0 allocs/op
BenchmarkGoFastDateToUnix-16                    18198014                63.59 ns/op            0 B/op          0 allocs/op
BenchmarkTimeUnixToDate-16                       2046949               582.9 ns/op             0 B/op          0 allocs/op
BenchmarkGoFastDateFromUnix-16                  12165228                90.23 ns/op            0 B/op          0 allocs/op
BenchmarkStressTimeDateToUnix-16                     135           8696891 ns/op               0 B/op          0 allocs/op
BenchmarkStressGoFastDateToUnix-16                   391           2855936 ns/op               0 B/op          0 allocs/op
BenchmarkStressTimeUnixToDate-16                      37          27315987 ns/op               0 B/op          0 allocs/op
BenchmarkStressGoFastDateFromUnix-16                 282           4285671 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/xnacly/go-fast-date  13.927s
```
