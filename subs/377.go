package subs

import "fmt"

func combinationSum4(nums []int, target int) int {
	m := make(map[int]int)
	dfs := func(_t int) int { return 0 }
	dfs = func(_t int) int {
		if v, ok := m[_t]; ok {
			return v
		}
		if _t == 0 {
			return 1
		}
		_r := 0
		for _, _n := range nums {
			if _n > _t {
				continue
			}
			_r += dfs(_t - _n)
		}
		m[_t] = _r
		return _r
	}
	dfs(target)
	return m[target]
}

func Test377() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
