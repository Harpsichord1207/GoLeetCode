package subs

import "fmt"

type parentAndDepth struct {
	depth  int
	parent *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	getDepth := func(v int, _node *TreeNode, depth int) parentAndDepth { return parentAndDepth{0, nil} }
	getDepth = func(v int, _node *TreeNode, depth int) parentAndDepth {
		if _node == nil {
			return parentAndDepth{-1, nil}
		}

		if _node.Val == v {
			return parentAndDepth{depth: depth, parent: nil}
		}

		if _node.Left != nil && _node.Left.Val == v {
			return parentAndDepth{depth: depth + 1, parent: _node}
		}

		if _node.Right != nil && _node.Right.Val == v {
			return parentAndDepth{depth: depth + 1, parent: _node}
		}

		leftDepth := getDepth(v, _node.Left, depth+1)
		if leftDepth.depth != -1 {
			return leftDepth
		}
		rightDepth := getDepth(v, _node.Right, depth+1)
		if rightDepth.depth != -1 {
			return rightDepth
		}
		return parentAndDepth{-1, nil}
	}

	xDepth := getDepth(x, root, 0)
	yDepth := getDepth(y, root, 0)

	return xDepth.depth == yDepth.depth && xDepth.parent != yDepth.parent

}

func Test993() {
	//node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 1}
	node1.Right = node3
	node1.Left = node2
	//node3.Right = node5
	node2.Left = node4
	fmt.Println(isCousins(node1, 4, 3))
}
