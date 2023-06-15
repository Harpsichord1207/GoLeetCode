package subs

import "math"

func maxLevelSum(root *TreeNode) int {
	m := make(map[int]int)

	f := func(_r *TreeNode, _d int) {}
	f = func(_r *TreeNode, _d int) {
		if _r == nil {
			return
		}

		m[_d] += _r.Val
		f(_r.Left, _d+1)
		f(_r.Right, _d+1)
	}
	f(root, 1)

	level := 1
	max := math.MinInt64
	r := 1
	for {
		v, ok := m[level]
		if !ok {
			break
		}
		if v > max {
			max = v
			r = level
		}
		level += 1
	}
	return r
}
