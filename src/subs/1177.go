package subs

import "fmt"

func canMakePaliQueries(s string, queries [][]int) []bool {
	var dp [][26]int
	dp = append(dp, [26]int{})
	for i := 0; i<len(s); i++ {
		if i > 0 {
			dp = append(dp, [26]int{})
			for j:=0;j<26;j++{
				dp[i][j] = dp[i-1][j]
			}
		}
		dp[i][s[i] - 'a'] += 1
	}

	var res []bool
	for _, query := range queries{
		l, r, k := query[0], query[1], query[2]
		if l == r {
			res = append(res, true)
			continue
		}
		c := 0
		if l == 0 {
			for j:=0;j<26;j++{
				if dp[r][j] % 2 != 0 {
					c += 1
				}
			}
		} else {
			for j:=0;j<26;j++{
				if (dp[r][j] - dp[l-1][j]) % 2 != 0 {
					c += 1
				}
			}
		}
		res = append(res, k*2+1>=c)
	}
	return res
}

func Test1177()  {
	s := "abcda"
    queries:= [][]int{{3,3,0},{1,2,0},{0,3,1},{0,3,2},{0,4,1}}
    fmt.Println(canMakePaliQueries(s, queries))
}
