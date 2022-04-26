package subs

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	k := -1
	var r []int
	for i <= j {
		m := i + (j-i)/2
		v := nums[m]
		if v == target {
			k = m
			break
		} else if v > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}

	if k == -1 {
		return r
	}

	s := -1
	for p := k - 1; p >= 0; p-- {
		if nums[p] == target {
			// not `r = append([]int{p}, r...)` which is very slow!
			s = p
		}
	}

	for q := s; q < len(nums); q++ {
		if nums[q] == target {
			r = append(r, q)
		}
	}

	return r
}
