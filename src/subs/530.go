package subs

import "math"

func getMinimumDifference(root *TreeNode) int {
	minNum := math.MaxInt64
	prev := math.MinInt64
	innerFunc := func(r *TreeNode) {}
	innerFunc = func(r *TreeNode) {
		if r == nil {
			return
		}
		innerFunc(r.Left)
		if prev != math.MinInt64 {
			diff := r.Val - prev
			if diff < minNum {
				minNum = diff
			}
		}
		prev = r.Val

		innerFunc(r.Right)
	}
	innerFunc(root)
	return minNum
}

func Test530(){
}
