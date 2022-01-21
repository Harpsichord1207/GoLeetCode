package subs

import "fmt"

func dfs(nodes []*TreeNode, out *[]int) {

	if len(nodes) == 0 {
		return
	}

	n := nodes[len(nodes)-1]
	*out = append(*out, n.Val)

	var children []*TreeNode

	for _, n := range nodes {
		if n.Left != nil {
			children = append(children, n.Left)
		}
		if n.Right != nil {
			children = append(children, n.Right)
		}
	}
	dfs(children, out)
}

func rightSideView(root *TreeNode) []int {
	var out []int
	if root == nil {
		return out
	}
	dfs([]*TreeNode{root}, &out)
	return out
}

func Test199() {
	//root := TreeNode{}
	fmt.Println(rightSideView(nil))
}
