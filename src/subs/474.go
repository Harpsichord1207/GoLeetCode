package subs

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	var dp [][]int
	for i:=0;i<=m;i++{
		dp = append(dp, make([]int, n+1))
	}
	for _, str := range strs {
		c0, c1 := 0, 0
		for _, s := range str {
			if s == '0' {c0++} else {c1++}
		}
		for i:=m;i>=c0;i--{
			for j:=n;j>=c1;j--{
				dp[i][j] = max(dp[i][j], dp[i-c0][j-c1]+1)
			}
		}
	}
	return dp[m][n]
}

func Test474() {
	fmt.Println(findMaxForm([]string{"10","0001","111001","1","0"}, 5, 3))
}
