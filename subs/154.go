package subs

import "fmt"

func findMin(nums []int) int {
	i, j := 0, len(nums)-1

	for i < j {
		m := (i + j) / 2
		v := nums[m]
		//fmt.Println(i, m, j)
		if v == nums[j] {
			j--
		} else if v > nums[j] {
			i = m + 1
		} else {
			// nums[m] may be the result, so we could not use `m - 1` here
			j = m
		}

	}

	return nums[i]
}

func Test154() {
	nums := []int{2, 2, 2, 0, 1}
	nums = []int{3, 1, 3}
	fmt.Println(findMin(nums))
}
