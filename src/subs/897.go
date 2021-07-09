package subs

func increasingBST(root *TreeNode) *TreeNode {

	var h *TreeNode = nil
	var t *TreeNode = nil
	helper := func(_root *TreeNode) {}
	helper = func(_root *TreeNode) {
		if _root.Left != nil {
			helper(_root.Left)
		}

		newNode := &TreeNode{_root.Val, nil, nil}

        if h == nil {h=newNode}
		if t == nil {t=newNode} else {t.Right=newNode;t=t.Right}

		if _root.Right != nil {
			helper(_root.Right)
		}
	}
	helper(root)
	return h
}


