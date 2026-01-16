package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(2009, 11, 17, 20, 34, 58, 65138723, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	p(then.Sub(now))
	p(now.Sub(then))
	p(now.Add(time.Hour))
	p(now.Add(-time.Hour))
}
