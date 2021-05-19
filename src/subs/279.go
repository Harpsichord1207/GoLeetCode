package subs

import (
	"fmt"
	"math"
)

func minInt(a int, b int) int {
	if a < b {return a}
	return b
}

func numSquares(n int) int {
	if n <= 2 {return n}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i:=3;i<=n;i++{
		s := int(math.Sqrt(float64(i)))
		if s * s == i {
			dp[i-1] = 1
		} else {
			r := math.MaxInt32
			for j:=1;j<i/2+1;j++{
				r = minInt(r, dp[i-j-1]+dp[j-1])
			}
			dp[i-1] = r
		}
	}
	return dp[n-1]
}

func Test279()  {
	fmt.Println(numSquares(13))
}