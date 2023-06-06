package subs

import (
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	if len(arr) < 2 {
		return true
	}

	diff := arr[1] - arr[0]
	for i, a := range arr {
		if i == 0 {
			continue
		}
		if a-arr[i-1] != diff {
			return false
		}
	}
	return true
}

func Test1502() {
	canMakeArithmeticProgression([]int{3, 5, 1})
}
