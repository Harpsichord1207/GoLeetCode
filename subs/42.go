package subs

import "fmt"

func trap(height []int) int {
	i := 0
	j := len(height) - 1
	r := 0
	for ; i < j; i++ {
		leftHigh := 0
		rightHigh := 0
		for ii := 0; ii <= i; ii++ {
			if height[ii] > leftHigh {
				leftHigh = height[ii]
			}
		}
		for jj := i; jj <= j; jj++ {
			if height[jj] > rightHigh {
				rightHigh = height[jj]
			}
		}
		if leftHigh > rightHigh {
			r += rightHigh

		} else {
			r += leftHigh
		}
		r -= height[i]
	}
	return r
}

func Test42() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
