package subs

import "fmt"



func longestOnes(nums []int, k int) int {
	i, j := 0, 0
	r := 0
	c := 0
	for j < len(nums) {
		if nums[j] == 0 {
			c ++
		}
		if c <= k {
			j ++
			continue
		}
		r = max(r, j-i)
		if nums[i] == 0 {
			c --
		}
		i ++
		j ++
	}
	return max(r, j-i)
}


func Test1004() {
	nums := []int{0,0,0,1}
	fmt.Println(longestOnes(nums, 4))
}
