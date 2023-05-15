package subs

import "fmt"

func integerReplacement(n int) int {

	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	}

	a := integerReplacement(n + 1)
	b := integerReplacement(n - 1)
	if a > b {
		return a
	}
	return b
}

func Test397() {
	fmt.Println(integerReplacement(8))
}
