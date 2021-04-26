package subs

import (
	"fmt"
)

func max(a int,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindromeSubseq(s string) int {
	l := len(s)
	var m = make(map[int]map[int]int)
	for i:=l-1;i>=0;i--{
		m[i] = make(map[int]int)
		m[i][i] = 1
		for j:=i+1;j<l;j++{
			if s[i]==s[j]{
				m[i][j] = m[i+1][j-1] + 2
			} else {
				m[i][j] = max(m[i][j-1], m[i+1][j])
			}
		}
	}
	return m[0][l-1]
}

func Test516()  {
	fmt.Println(longestPalindromeSubseq("dabbcbba"))
	fmt.Println(longestPalindromeSubseq("a"))
}