package subs

import "fmt"

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	for i:=1; i<=l/2+1; i++ {
		if l % i != 0 || l == i {
			continue
		}
		isBreak := false
		for j:=i;j<l;j+=i{
			a := s[j:i+j]
			b := s[:i]
			if a != b {
				isBreak = true
				break
			}
		}
		if !isBreak{return true}
	}
	return false
}

func Test459() {
	fmt.Println(repeatedSubstringPattern("ab"))
}
