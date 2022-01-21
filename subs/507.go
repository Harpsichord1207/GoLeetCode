package subs

import (
	"fmt"
)

func checkPerfectNumber(num int) bool {
	half := 28 / 2
	sum := 0
	for i := 1; i < half+1; i++ {
		if num%i == 0 {
			sum += i
			//fmt.Println(i)
		}
	}
	return sum == num
	// return num==6 || num==28 || num==496 || num==8128 || num==3355036
}

func Test507() {
	fmt.Println(checkPerfectNumber(28))
}
