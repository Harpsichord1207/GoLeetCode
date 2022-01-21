package subs

import "fmt"

func kthSmallest(root *TreeNode, k int) int {
	var nums []int
	_helper := func(_root *TreeNode) {}
	_helper = func(_root *TreeNode) {
		if _root == nil {
			return
		}

		if _root.Left != nil {
			_helper(_root.Left)
		}

		nums = append(nums, _root.Val)

		if _root.Right != nil {
			_helper(_root.Right)
		}

	}
	_helper(root)
	return nums[k-1]

}

func Test230() {
	root := TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
	fmt.Println(kthSmallest(&root, 1))
}
