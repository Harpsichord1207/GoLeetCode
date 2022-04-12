package subs

func min583(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func minDistance583(word1 string, word2 string) int {
	// dp[i][j] 表示word1[:i] 与 word2[:j]的距离
	// 状态转移方程:
	// dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2*int(word1[i-1]!=word2[j-1]))
	// 基本和72题是一样的

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
				t = dp[i-1][j-1] + 2
			}
			dp[i][j] = min583(r, s, t)
		}
	}
	// fmt.Println(dp)
	return dp[l1][l2]
}

func Test583() {

}
