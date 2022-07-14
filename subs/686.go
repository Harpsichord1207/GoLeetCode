package subs

import (
	"fmt"
	"strings"
)

// 莫名其妙
func repeatedStringMatch(a string, b string) int {
	var r = 1
	var ar = a
	for {
		if strings.Contains(ar, b) {
			return r
		} else if len(ar) > len(b) && !strings.Contains(ar+ar, b) {
			return -1
		}
		r += 1
		ar += a

	}
}

func Test686() {
	a := "abcd"
	b := "abcdb"
	fmt.Println(repeatedStringMatch(a, b))
}
