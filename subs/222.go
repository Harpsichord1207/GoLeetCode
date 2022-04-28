package subs

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := getHeight(root.Left, true)
	rightDepth := getHeight(root.Right, false)

	if leftDepth == rightDepth {
		return int(math.Pow(2, float64(leftDepth))) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func getHeight(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if isLeft {
		return 1 + getHeight(root.Left, isLeft)
	}
	return 1 + getHeight(root.Right, isLeft)
}
