package subs

import (
	"fmt"
)

func nextGreaterElements(nums []int) []int {
	var res []int
	nums2 := append(nums, nums...)
	for i, n := range nums {
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > n {
				res = append(res, nums2[j])
				break
			}
			if j == len(nums2)-1 {
				res = append(res, -1)
			}
		}
	}
	return res
}

func Test503() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElements(nums))
	//fmt.Println(binarySearch(nums, 1))
}
