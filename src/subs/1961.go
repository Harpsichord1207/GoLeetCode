package subs

import "fmt"

func isPrefixString(s string, words []string) bool {
	i, j := 0, len(s)
	for _, word := range words {
		l := len(word)

		if j < i+l {
			return false
		}
		for k := 0; k < l; k++ {
			if word[k] != s[i+k] {
				return false
			}
		}
		i += l
		if j == i {
			return true
		}
	}
	return false
}

func Test1981() {
	fmt.Println(isPrefixString("rx", []string{"om", "uwr"}))
}
