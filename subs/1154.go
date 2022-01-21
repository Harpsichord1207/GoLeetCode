package subs

import (
	"fmt"
	"time"
)

func dayOfYear(date string) int {
	year := date[:4]
	start := year + "-01-01"
	s, _ := time.Parse("2006-01-02", start)
	e, _ := time.Parse("2006-01-02", date)
	return int(e.Sub(s).Hours()/24 + 1)
}

func Test1154() {
	fmt.Println(dayOfYear("2019-02-10"))
}
