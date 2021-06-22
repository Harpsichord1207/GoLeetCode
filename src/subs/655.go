package subs

import (
	"fmt"
	"strconv"
)

func getTreeDepth(root *TreeNode) int {
	if root == nil {return 0}
	left := getTreeDepth(root.Left)
	right := getTreeDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}


func printTree(root *TreeNode) [][]string {

	// len(output)
	depth := getTreeDepth(root)
	// len(output[n])
	length := 1<<depth-1

	output := make([][]string, depth)

	for i:=0;i<depth;i++ {
		output[i] = make([]string, length)
		for j:=0;j<length;j++ {
			output[i][j] = ""
		}
	}

	_helper := func(_node *TreeNode, _level int, _start int, _end int){}
	_helper = func(_node *TreeNode, _level int, _start int, _end int) {
		if _node == nil {
			return
		}
		mid := (_start + _end) / 2
		output[_level][mid] = strconv.Itoa(_node.Val)
		_helper(_node.Left, _level+1, _start, mid-1)
		_helper(_node.Right, _level+1, mid+1, _end)
	}
	_helper(root, 0, 0, length)

	return output
}

func Test655()  {
	root := TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(printTree(&root))
}


