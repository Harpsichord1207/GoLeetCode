package subs

import "fmt"

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i:=l-2;i>=0;i--{
		d := l - i - 1
		if nums[i] >= d {
			dp[i] = 1
		} else {
			k := 10001
			for j:=i+1;j<=nums[i]+i && j<=l-1;j++{
				if k > dp[j] + 1 && dp[j] > 0 {
					k = dp[j] + 1
				}
			}
			dp[i] = k
		}
	}
	return dp[0]
}

func Test45() {
	fmt.Println(jump([]int{0}))
}
