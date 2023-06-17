package subs

import (
	"fmt"
	"math"
	"sort"
)

func makeArrayIncreasing(arr1 []int, arr2 []int) int {

	sort.Ints(arr2)
	l2 := len(arr2)

	// 从排好序的arr2中找到>_v的最小值
	bisectRight := func(_v int) int {
		_a, _b := 0, l2
		for _a < _b {
			_m := (_a + _b) / 2
			_p := arr2[_m]
			if _p <= _v {
				_a = _m + 1
			} else {
				_b = _m
			}
		}
		return _b
	}

	memo := make(map[string]int)

	// dfs(i, prev) arr中第i个元素排好序的最小次数
	dfs := func(_i int, _pre int) int { return 0 }
	dfs = func(_i int, _pre int) int {
		if _i == len(arr1) {
			return 0
		}

		key := fmt.Sprintf("%d_%d", _i, _pre)
		if value, ok := memo[key]; ok {
			return value
		}

		cost := math.MaxInt64

		// arr1[_i] 比前一个数大，可以不做任何操作，dfs(i) = dfs(i+1)
		if arr1[_i] > _pre {
			cost = dfs(_i+1, arr1[_i])
		}

		// 也可以从arr2中找到比pre大的最小的一个数，做一次更新
		idx := bisectRight(_pre)
		if idx < l2 && dfs(_i+1, arr2[idx]) != math.MaxInt64 {
			newCost := 1 + dfs(_i+1, arr2[idx])
			if newCost < cost {
				cost = newCost
			}
		}
		//fmt.Printf("Cost of i = %d, pre = %d is %d.\n", _i, _pre, cost)
		memo[key] = cost
		return cost
	}
	r := dfs(0, -1)
	if r < math.MaxInt64 {
		return r
	}
	return -1
}

func Test1187() {
	//fmt.Println(makeArrayIncreasing([]int{1, 5, 2, 6, 7}, []int{1, 3, 2, 4}))
	//fmt.Println(makeArrayIncreasing([]int{1, 5, 2, 6, 7}, []int{1, 3, 4}))
	fmt.Println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}))
}
