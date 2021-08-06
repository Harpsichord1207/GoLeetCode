package subs

import (
	"fmt"
	"strings"
)

type Ele394 struct {
	Times int
	Res   string
}

func decodeString(s string) string {
	var stack []Ele394
	var times int
	var res string

	m := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			times = times*10 + m[c]
		} else if c == '[' {
			stack = append(stack, Ele394{Times: times, Res: res})
			res = ""
			times = 0
		} else if c == ']' {
			st := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = st.Res + strings.Repeat(res, st.Times)
		} else {
			res = res + string(c)
		}
	}

	return res
}

func Test394() {
	fmt.Println(decodeString("3[a]2[bc]"))
}
