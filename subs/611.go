package subs

import "sort"

func triangleNumber(nums []int) int {
	size := len(nums)
	sort.Ints(nums)
	res := 0
	for k := size - 1; k >= 0; k-- {
		i, j := 0, k-1
		for i < j {
			if nums[i]+nums[j] > nums[k] {
				res += (j - i)
				j -= 1
			} else {
				i += 1
			}
		}
	}
	return res
}
