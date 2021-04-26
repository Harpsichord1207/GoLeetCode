package subs

import "fmt"

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	la := len(a)
	lb := len(b)
	if la > lb {
		return la
	}
	return lb
}


func Test521(){
	fmt.Println(findLUSlength("abc", "bc"))
}
