package subs

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	r := 0
	for i, n1 := range nums1 {
		j, k := i, len(nums2)-1
		if j > k || n1 > nums2[j] {
			continue
		}
		for j <= k {
			p := (j + k) / 2
			v := nums2[p]
			// last num in nums2[j:] larger than or equal to n1
			if v >= n1 {
				j = p + 1
			} else {
				k = p - 1
			}
		}
		if k-i > r {
			r = k - i
		}
	}
	return r
}

func Test1855() {
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}
	fmt.Println(maxDistance(nums1, nums2))
}
