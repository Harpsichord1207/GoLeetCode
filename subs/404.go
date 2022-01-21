package subs

import "fmt"

func sumOfLeftLeaves(root *TreeNode) int {
	helper := func(r *TreeNode) {}
	s := 0
	helper = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil && r.Left.Left == nil && r.Left.Right == nil {
			s += r.Left.Val
		}
		helper(r.Left)
		helper(r.Right)
	}
	helper(root)
	return s
}

func Test404() {
	l := TreeNode{9, nil, nil}
	r := TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}
	fmt.Println(sumOfLeftLeaves(&TreeNode{3, &l, &r}))
}
