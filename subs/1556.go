package subs

import (
	"fmt"
	"strconv"
)

func thousandSeparator(n int) string {
	var s string
	var cnt3 int

	for {
		i := n % 10

		if cnt3 > 0 && cnt3%3 == 0 {
			s = strconv.Itoa(i) + "." + s
		} else {
			s = strconv.Itoa(i) + s
		}

		cnt3++

		n = n / 10
		if n == 0 {
			break
		}
	}
	return s
}

func Test1556() {
	fmt.Println(thousandSeparator(123))
	fmt.Println(thousandSeparator(1234))
	fmt.Println(thousandSeparator(12345))
}
