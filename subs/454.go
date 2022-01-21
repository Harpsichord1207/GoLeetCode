package subs

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var res int
	m := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			res += m[-n3-n4]
		}
	}
	return res
}

func Test454() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{3, 4}, []int{-1, -3}, []int{-2, -4}))
}
