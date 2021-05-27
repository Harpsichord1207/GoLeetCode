package subs

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 == 0 { return s2 == s3 }
	if l2 == 0 { return s1 == s3 }
	if l1 + l2 != len(s3) { return false }

	e1 := s1[0] == s3[0]
	e2 := s2[0] == s3[0]
	if e1 && e2 {
		return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
	} else if e1 {
		return isInterleave(s1[1:], s2, s3[1:])
	} else if e2 {
		return isInterleave(s1, s2[1:], s3[1:])
	}
	return false
}

func isInterleaveDp(s1 string, s2 string, s3 string) bool {
	var dp [][]bool
	l1 := len(s1)
	l2 := len(s2)
	if l1 + l2 != len(s3){return false}
	for i:=0;i<=l1;i++{
		dp = append(dp, make([]bool, l2+1))
	}
	dp[0][0] = true
	for i:=1;i<=l1;i++{
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j:=1;j<=l2;j++{
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i:=1;i<=l1;i++{
		for j:=1;j<=l2;j++{
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[l1][l2]
}

func Test97()  {
	fmt.Println(isInterleave("aaaa", "aa", "aaa"))
	fmt.Println(isInterleaveDp("aabcc", "dbbca", "aadbbcbcac"))
}
