package subs

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {

	le := len(intervals)
	startMap := make(map[int]int, le)
	sortedStart := make([]int, le)

	for i, intV := range intervals {
		startMap[intV[0]] = i
		sortedStart[i] = intV[0]
	}

	sort.Ints(sortedStart)

	helper := func(_v int) (int, error) {
		_i, _j := 0, le-1
		for _i <= _j {
			_k := (_i + _j) / 2
			if sortedStart[_k] >= _v {
				_j = _k - 1
			} else {
				_i = _k + 1
			}
		}
		if _i < le {
			return sortedStart[_i], nil
		}

		return -1, fmt.Errorf("no value bigger than %d", _v)
	}

	res := make([]int, le)
	for i, intV := range intervals {
		v := intV[1]
		// 找到第一个比v大的数
		if f, e := helper(v); e == nil {
			res[i] = startMap[f]
		} else {
			res[i] = f
		}

	}

	return res

}

func Test436() {
	intervals := [][]int{{3, 4}, {2, 3}, {1, 2}}
	fmt.Println(findRightInterval(intervals))
}
