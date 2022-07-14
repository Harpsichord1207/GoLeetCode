package subs

import (
	"fmt"
	"strings"
)

// 100% beat!
func repeatedStringMatch(a string, b string) int {
	la := len(a)
	lb := len(b)
	c := lb / la
	if lb%la != 0 {
		c++
	}

	ar := strings.Repeat(a, c)

	if strings.Contains(ar, b) {
		return c
	}

	// why here ?
	if strings.Contains(ar+a, b) {
		return c + 1
	}

	return -1
}

func Test686() {
	a := "a"
	b := "aa"
	fmt.Println(repeatedStringMatch(a, b))
}
