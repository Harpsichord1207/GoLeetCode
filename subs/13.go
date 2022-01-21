package subs

import "fmt"

var romanMap = map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	l := len(s)
	n := 0
	for i := 0; i < l-1; i++ {
		if romanMap[s[i+1]] > romanMap[s[i]] {
			n -= romanMap[s[i]]
		} else {
			n += romanMap[s[i]]
		}
	}
	n += romanMap[s[l-1]]
	return n
}

func Test13() {
	fmt.Println(romanToInt("MCMXCIV"))
}
