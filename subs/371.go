package subs

import "fmt"

func getSum(a int, b int) int {
	for {
		if b == 0 {
			break
		}
		c := a ^ b
		b = (a & b) << 1
		a = c
	}
	return a
}

func Test371() {
	fmt.Println(getSum(7, 8))
}
