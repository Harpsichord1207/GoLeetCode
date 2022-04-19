package subs

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	// fmt.Println(nums)

	helper := func(target int) int {
		i, j := 0, len(nums)-1
		for i <= j {
			k := (i + j) / 2
			// 找到第一个大于等于target的位置
			if nums[k] < target {
				i = k + 1
			} else {
				j = k - 1
			}
		}
		return j + 1
	}

	for i := 0; i <= nums[len(nums)-1]; i++ {
		q := helper(i)
		// fmt.Println(i, q)
		if len(nums)-q == i {
			return i
		}
	}
	return -1
}

func Test1608() {
	nums := []int{3, 9, 7, 8, 3, 8, 6, 6}
	fmt.Println(specialArray(nums))
}
