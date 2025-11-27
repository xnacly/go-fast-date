package gofastdate

import (
	"fmt"
	"testing"
	"time"
)

func TestReadme(t *testing.T) {
	date := FromUnix(time.Now().Unix())
	fmt.Println(date)

	fmt.Println(date.Time())
	fmt.Println(date.Unix(), date.Time().Unix())
}
