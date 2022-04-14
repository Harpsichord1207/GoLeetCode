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
		if arr2[0] >= _num {
			return arr2[0]-_num > d
		}

		if arr2[j] <= _num {
			return _num-arr2[j] > d
		}

		for i <= j {
			k := (i + j) / 2
			v := arr2[k]
			if v == _num {
				return false
			}
			if v < _num {
				i = k + 1
			} else {
				j = k - 1
			}
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
