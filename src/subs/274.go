package subs

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	l := len(citations)
	sort.Ints(citations)
	res := 0
	for i, c := range citations {
		curr := l - i
		if curr > c {
			curr = c
		}
		if curr > res {
			res = curr
		}
	}
	return res
}


func Test274()  {
	r := hIndex([]int{0, 0})
	fmt.Println(r)
}
