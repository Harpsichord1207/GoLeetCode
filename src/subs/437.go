package subs

import "fmt"

func helper(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {return res}
	if root.Val == targetSum {res += 1}
	newTarget := targetSum - root.Val
	return helper(root.Right, newTarget) + res + helper(root.Left, newTarget)
	
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {return 0}
	return helper(root, targetSum) + pathSum(root.Right, targetSum) + pathSum(root.Left, targetSum)
}

func Test437()  {
	root := &TreeNode{5, &TreeNode{4, nil, nil}, nil}
	fmt.Println(pathSum(root, 3))
}