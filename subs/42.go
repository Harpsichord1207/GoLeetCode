package subs

import "fmt"

func trap(height []int) int {
	l := len(height)
	r := 0
	if l == 0 {
		return 0
	}

	leftHighArr := make([]int, l) // leftHighArr[i]表示i左边最高的柱子
	rightHighArr := make([]int, l)

	for p := 0; p < len(height); p++ {
		if p == 0 || height[p] > leftHighArr[p-1] {
			leftHighArr[p] = height[p]
		} else {
			leftHighArr[p] = leftHighArr[p-1]
		}
	}

	for p := l - 1; p >= 0; p-- {
		if p == l-1 || height[p] > rightHighArr[p+1] {
			rightHighArr[p] = height[p]
		} else {
			rightHighArr[p] = rightHighArr[p+1]
		}
	}

	for i := 0; i < l; i++ {
		if leftHighArr[i] > rightHighArr[i] {
			r += rightHighArr[i]
		} else {
			r += leftHighArr[i]
		}
		r -= height[i]
	}

	return r
}

func Test42() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
