package subs

func isMatch(s string, p string) bool {
	l1 := len(s)
	l2 := len(p)
	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}

	// 空vs空满足
	dp[0][0] = true

	// 如果i-2匹配，p第i个字符是'*'，则i匹配
	for i := 1; i <= l2; i++ {
		// '*'前面一定有字符，因此i-2不会越界
		if p[i-1] == '*' && dp[0][i-2] {
			dp[0][i] = true
		}
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			// p第j个字符 等于 s第i个字符 or '.'
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
				// p第j个字符为 '*'
			} else if p[j-1] == '*' {
				// p第j-1个字符 不等于 s第i个字符 or '.'，此时p第j个字符 '*' 只能匹配第j-1个字符0次
				if p[j-2] != s[i-1] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
					// 否则，'*' 可以匹配第j-1个字符0次(j-2)、1次(j-1)或者多次(i-1)
					// 多次包含i-1/i-2/i-3... 但如果i-2匹配则i-1必然匹配
				} else {
					dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i][j-2]
				}
			}
		}
	}
	return dp[l1][l2]
}
