package subs

import "fmt"

func min72(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		for j := range dp[i] {
			dp[i][j] = 502
		}
		dp[i][0] = i
	}

	// dp[i][j] 表示word1[:i] 与 word2[:j]的距离
	// 状态转移方程:
	// dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j], dp[i-1][j-1]+int(word1[i-1]!=word2[j-1]))

	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			r := dp[i][j-1] + 1
			s := dp[i-1][j] + 1
			t := -1
			if word1[i-1] == word2[j-1] {
				t = dp[i-1][j-1]
			} else {
				t = dp[i-1][j-1] + 1
			}
			dp[i][j] = min72(r, s, t)
		}
	}
	// fmt.Println(dp)
	return dp[l1][l2]
}

func Test72() {
	fmt.Println(minDistance("horse", "ros"))
}
