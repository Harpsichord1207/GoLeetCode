package subs

import "fmt"

func tribonacci(n int) int {
	a, b, c := 0, 1, 1

	if n < 2 {
		return n
	}

	if n == 2 {
		return 1
	}

	for i := 1; i <= n-2; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

func Test1137() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, tribonacci(i))
	}
}
