package subs

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	half := int(math.Sqrt(float64(c)))
	for i := 0; i < half+1; i++ {
		j := c - i*i
		k := int(math.Sqrt(float64(j)))
		if k*k == j {
			return true
		}
	}
	return false
}

func Test633() {
	fmt.Println(judgeSquareSum(2147483646))
}
