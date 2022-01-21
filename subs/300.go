package subs

import "fmt"

func lengthOfLIS(nums []int) int {
	res := 1
	dp := make([]int, len(nums))

	for i, n := range nums {
		dp[i] = 1
		l := 1
		for j := i - 1; j >= 0; j-- {
			if dp[j]+1 > l && n > nums[j] {
				l = dp[j] + 1
			}
		}
		dp[i] = l
		if res < l {
			res = l
		}
	}
	return res
}

func Test300() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
