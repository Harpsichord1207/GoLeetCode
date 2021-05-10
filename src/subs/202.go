package subs

import "fmt"

func calc(n int) int {
	res := 0
	for n > 0 {
		a := n % 10
		res += a * a
		n /= 10
	}
	fmt.Println(res)
	return res
}

func isHappy(n int) bool {
	m := make(map[int]int)
	for {
		m[n] = 1
		n = calc(n)
		if m[n] == 1 {return n == 1}
	}
}

func Test202()  {
	fmt.Println(isHappy(2))
}