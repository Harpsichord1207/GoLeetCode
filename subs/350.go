package subs

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	l1, l2 := len(nums1), len(nums2)
	var shortNums, longNums []int
	if l1 > l2 {
		shortNums = nums2
		longNums = nums1
	} else {
		shortNums = nums1
		longNums = nums2
	}
	l3 := len(longNums)

	helper := func(_s int, _e int, _t int) int {
		_i, _j := _s, _e
		for _i <= _j {
			_m := (_i + _j) / 2
			_v := longNums[_m]
			if _v >= _t {
				_j = _m - 1
			} else {
				_i = _m + 1
			}
		}
		_p := _j + 1
		if _p < l3 && longNums[_p] == _t {
			return _p
		}
		return -1
	}

	start := 0
	var r []int
	for _, n := range shortNums {
		p := helper(start, l3-1, n)
		if p != -1 {
			start = p + 1
			r = append(r, n)
		}
	}
	return r
}
