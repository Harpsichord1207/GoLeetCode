package subs

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	s, c := 0, 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i+2]-nums[i+1] == nums[i+1]-nums[i] {
			c++
			s += c
		} else {
			c = 0
		}
	}
	return s
}

func Test413() {
	fmt.Println(numberOfArithmeticSlices([]int{7, 7, 7, 7}))
}
