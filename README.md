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
`time.Unix`:
