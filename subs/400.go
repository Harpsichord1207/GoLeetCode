package subs

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	m, k, d := 1, 9, 1
	for n > k*m {
		m, k, n, d = m+1, k*10, n-k*m, d*10
	}
	// now the target number is the nth number in [d, ~)
	n = n - 1
	number := (n / m) + d
	ns := strconv.Itoa(number)
	return int(ns[n%m] - '0')

}

func Test400() {
	fmt.Println(findNthDigit(12))
}
