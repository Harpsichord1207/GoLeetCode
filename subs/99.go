package subs

import (
	"math"
)

func recoverTree(root *TreeNode) {

	var node1 *TreeNode = nil
	var node2 *TreeNode = nil
	var pre = &TreeNode{Val: math.MinInt32}

	inOrder := func(_r *TreeNode) {}
	// 左中右中序遍历
	// 在遍历过程中找到位置不对的两个node，分别用node1/node2保存
	// 最后交换两个node的值
	inOrder = func(_r *TreeNode) {
		if _r == nil {
			return
		}

		inOrder(_r.Left)

		// node1为nil且上一个node的值大于当前node的值，说明上一个node是不对的，赋给node1
		if node1 == nil && pre.Val > _r.Val {
			node1 = pre
		}
		// node1已经不为nil了，再次遇到上一个node的值大于当前node的值，说明当前node的位置不对，赋给node2
		if node1 != nil && pre.Val > _r.Val {
			node2 = _r
		}

		pre = _r

		inOrder(_r.Right)
	}

	inOrder(root)

	node1.Val, node2.Val = node2.Val, node1.Val
}

func Test99() {
	r := TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}
	recoverTree(&r)
}
