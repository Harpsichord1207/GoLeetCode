package subs


func postorderTraversal(root *TreeNode) []int {
	var res []int
	helper := func(_root *TreeNode) {}
	helper = func(_root *TreeNode) {
		if _root == nil {return}
		res = append(res, _root.Val)
		helper(_root.Right)
		helper(_root.Left)
	}
	helper(root)
	return res
}