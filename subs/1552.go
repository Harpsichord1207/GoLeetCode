package subs

import "sort"

func maxDistance1552(position []int, m int) int {
	sort.Ints(position)
	size := len(position)
	min := 0
	max := position[size-1] / (m - 1)

	// check if a force is valid
	helper := func(f int) bool {
		_m := m
		last_num := position[0]
		for i := 1; i < len(position); i++ {
			// f是最小值，因此间隔要大于等于f
			if position[i]-last_num >= f {
				last_num = position[i]
				_m--
			}
		}
		return _m <= 1
	}

	for min <= max {
		mid := min + (max-min)/2
		if helper(mid) {
			min = mid + 1
		} else {
			max = mid - 1
		}

	}
	return max
}
