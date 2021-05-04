package subs

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	ans := 0
	hasOdd := false
	for i:=0;i<len(s);i++{
		m[s[i]] += 1
	}
	for _, v := range m {
		if v % 2 == 0 {
			ans += v
		} else {
			hasOdd = true
			ans += v - 1
		}
	}
	if hasOdd {ans += 1}
	return ans
}

func Test409()  {
	fmt.Println(longestPalindrome("abccccdd"))
}