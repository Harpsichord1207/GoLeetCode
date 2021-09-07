package subs

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}

	var left int
	if root.Left.Val == root.Val {
		left = findSecondMinimumValue(root.Left)
	} else {
		left = root.Left.Val
	}

	var right int
	if root.Right.Val == root.Val {
		right = findSecondMinimumValue(root.Right)
	} else {
		right = root.Right.Val
	}

	if left == -1 || right == -1 {
		if left > right {
			return left
		}
		return right
	} else {
		if left > right {
			return right
		}
		return left
	}

}
