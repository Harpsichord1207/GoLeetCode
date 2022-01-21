package subs

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var output []int
	if root == nil {
		return output
	}
	output = append(output, root.Val)
	for _, child := range root.Children {
		output = append(output, preorder(child)...)
	}
	return output
}
