package subs

import "fmt"

func isBadVersion(n int) bool {
	return n == 4
}

func firstBadVersion(n int) int {
	i, j := 1, n
	for i <= j {
		m := (i + j) / 2
		fmt.Println(i, m, j)
		if isBadVersion(m) {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return i
}

func Test278() {
	fmt.Println(firstBadVersion(5))
}
