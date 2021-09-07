package subs

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var res []string
	l := len(score)
	sortedScore := append([]int(nil), score...)
	sort.Ints(sortedScore)
	m := make(map[int]int)
	for i, v := range sortedScore {
		m[v] = i
	}
	for _, v := range score {
		i := l - m[v] - 1
		if i == 0 {
			res = append(res, "Gold Medal")
		} else if i == 1 {
			res = append(res, "Silver Medal")
		} else if i == 2 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(i+1))
		}
	}
	return res
}

func Test506() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
}
