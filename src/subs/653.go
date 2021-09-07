package subs

import "fmt"

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	memo := make(map[int]bool)
	res := false

	_helper := func(_root *TreeNode) {}
	_helper = func(_root *TreeNode) {
		if _root == nil {
			return
		}
		_helper(_root.Left)
		if memo[k-_root.Val] {
			res = true
			return
		}
		memo[_root.Val] = true
		_helper(_root.Right)
	}
	_helper(root)

	return res
}

func Test653() {
	root := TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(findTarget(&root, 9))
}
