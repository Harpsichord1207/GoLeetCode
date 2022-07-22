package subs

func findTilt(root *TreeNode) int {

	abs := func(_num int) int {
		if _num > 0 {
			return _num
		}
		return -_num
	}
	var r int
	// post order traversal
	var helper func(_root *TreeNode) int
	helper = func(_root *TreeNode) int {
		if _root == nil {
			return 0
		}
		rightSum := helper(_root.Right)
		leftSum := helper(_root.Left)
		r += abs(leftSum - rightSum)
		return leftSum + rightSum + _root.Val
	}
	helper(root)

	return r
}
