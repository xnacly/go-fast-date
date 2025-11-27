# go-fast-date

Go port of _A Very Fast 64–Bit Date Algorithm_, based on [A
Very Fast 64–Bit Date Algorithm: 30–40%
faster](https://www.benjoffe.com/fast-date-64) by Ben Joffe.
Special thanks to [Ben Joffe](https://www.benjoffe.com/) for
writing such an easy to follow blog post and implementation.

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
- `gofastdate.Date.ToTime`

While not depending on any dependency and striving to match
Joffe's implementation, where applicable.

## Benchmarks

Benchmarks focus on comparing `gofastdate.FromUnix` with
`time.Unix`. Run with `go test ./... -bench=.`

### Low end hardware:

- micro benchmark: fastdate `~13%` faster
- stress: fastdate `~4.2x` faster (8.98ms vs 37.38ms)

```text
goos: linux
goarch: amd64
pkg: github.com/xnacly/go-fast-date
cpu: Intel(R) Core(TM) i5-3210M CPU @ 2.50GHz
BenchmarkTimeUnix-4                     21141045                56.21 ns/op
BenchmarkGoFastDateUnix-4               24253924                48.90 ns/op
BenchmarkStressTimeUnixToDate-4               32             37382839 ns/op
BenchmarkStressGoFastDateFromUnix-4          132              8977516 ns/op
PASS
ok      github.com/xnacly/go-fast-date  6.369s
```

### Mid end hardware:

- micro benchmark: fastdate `~32%` faster
- stress: fastdate `~4.1x` faster (5.16ms vs 21.41ms)

```text
goos: linux
goarch: amd64
pkg: github.com/xnacly/go-fast-date
cpu: AMD Ryzen 7 3700X 8-Core Processor
BenchmarkTimeUnix-16                            35375085                37.20 ns/op
BenchmarkGoFastDateUnix-16                      47718576                25.21 ns/op
BenchmarkStressTimeUnixToDate-16                      56             21414260 ns/op
BenchmarkStressGoFastDateFromUnix-16                 232              5158183 ns/op
PASS
ok      github.com/xnacly/go-fast-date  5.463s
```
