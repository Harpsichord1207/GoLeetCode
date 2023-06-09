package subs

import (
	"fmt"
)

func longestNiceSubstring(s string) string {

	l := len(s)
	r := ""
	if l == 0 {
		return r
	}

	m := make(map[uint8]bool, l)

	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}

	for i := 0; i < len(s); i++ {
		v := s[i]
		if v < 'a' {
			v += 32
		} else {
			v -= 32
		}
		if !m[v] {
			s1 := longestNiceSubstring(s[:i])
			s2 := longestNiceSubstring(s[i+1:])
			if len(s1) >= len(s2) {
				return s1
			}
			return s2
		}
	}
	return s
}

func Test1763() {
	fmt.Println(longestNiceSubstring("dDzeE"))
}
