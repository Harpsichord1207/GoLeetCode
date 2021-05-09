package subs

func invertTree(root *TreeNode) *TreeNode {
	helper := func(_root *TreeNode) {}
	helper = func(_root *TreeNode) {
		if _root == nil {return}
		_root.Left, _root.Right = _root.Right, _root.Left
		helper(_root.Left)
		helper(_root.Right)
	}
	helper(root)
	return root
}

