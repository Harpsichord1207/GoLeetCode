package subs

import "fmt"

func findLength(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	dp := make([][]int, l1+1)
	for i:=0; i<l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}

	res := 0
	for i:=0; i<l1+1; i++ {
		for j:=0; j<l2+1; j++ {
			if i ==0 || j == 0 {
				dp[i][j] = 0
			} else if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if res < dp[i][j] {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}

func Test718() {
	n1 := []int{1, 2, 3, 2, 1}
	n2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(n1, n2))
}