package subs

import (
	"sort"
)

func abs1385(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	le2 := len(arr2)

	helper := func(_num int) bool {
		i, j := 0, le2-1
		for i <= j {
			k := (i + j) / 2
			v := nums[k]
			if v == _num {
				return false
			}
			if v < _num {
				i = k + 1
			} else {
				j = k - 1
			}
		}

		if i == le2 {
			return abs1385(arr2[i-1]-_num) > d
		} else if i == 0 {
			return abs1385(arr2[i]-_num) > d
		}
		return abs1385(arr2[i]-_num) > d && abs1385(arr2[i-1]-_num) > d
	}

	r := 0
	for _, a := range arr1 {
		if helper(a) {
			r++
		}
	}
	return r
}
