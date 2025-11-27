# go-fast-date

Go port of _A Very Fast 64–Bit Date Algorithm_, based on [A
Very Fast 64–Bit Date Algorithm: 30–40%
faster](https://www.benjoffe.com/fast-date-64)

Special thanks to [Ben Joffe](https://www.benjoffe.com/) for
writing such an easy to follow blog post and implementation

## Features

`go-fast-date` implements Ben Joffes very fast 64-bit date
algorithm and provides date round trip functions for
[`time.Time`](https://pkg.go.dev/time) compability:

- `gofastdate.FromUnix`
- `gofastdate.FromTime`
- `gofastdate.Date.ToTime`

## Benchmarks
