package subs

import "fmt"

//Definition for a binary tree node.

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func inOrder(r *TreeNode) []int {
	var list []int
	if r.Left != nil {
		list = append(list, inOrder(r.Left)...)

	}
	list = append(list, r.Val)
	if r.Right != nil {
		list = append(list, inOrder(r.Right)...)
	}
	return list
}



func findMode(root *TreeNode) []int {
	var maxCountNum []int
	maxCount := 0
	currentCount := 0
	currentNum := 0
	inOrderList := inOrder(root)
	for _, v := range inOrderList {
		if currentNum != v {
			currentNum = v
			currentCount = 0
		}
		currentCount ++
		if currentCount > maxCount {
			maxCount = currentCount
			maxCountNum = []int{v}
		} else if currentCount == maxCount {
			maxCountNum = append(maxCountNum, v)
		}
	}
	return maxCountNum
}

func Test501()  {
	r := TreeNode{1, nil, &TreeNode{2, &TreeNode{2, nil, nil}, nil}}
	r = TreeNode{1, nil, &TreeNode{2, nil, nil}}
	fmt.Println(findMode(&r))
}
