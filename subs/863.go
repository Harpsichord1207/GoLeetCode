package subs

import "fmt"

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	// get nodes with distance k from subtrees
	var _l1 []int
	helper1 := func(_node *TreeNode, _k int) {}
	helper1 = func(_node *TreeNode, _k int) {
		if _node == nil {
			return
		}
		if _k == 0 {
			_l1 = append(_l1, _node.Val)
			return
		}
		helper1(_node.Left, _k-1)
		helper1(_node.Right, _k-1)
	}

	// get distance from _node to _target
	helper2 := func(_node *TreeNode, _k int) int { return 0 }
	helper2 = func(_node *TreeNode, _k int) int {
		if _node == nil {
			return -1
		}

		if _node == target {
			helper1(_node, k)
			return 0
		}

		dl := helper2(_node.Left, _k)
		// found target in _node.Left, distance = dl
		if dl != -1 {
			// if the distance from _node to _target == k
			if dl+1 == _k {
				_l1 = append(_l1, _node.Val)
			} else if dl+1 > _k {
				// 距离超过了 应该返回-1？
				return -1
			} else {
				// 距离不够，右子树来补
				helper1(_node.Right, _k-dl-2)
			}
			return dl + 1
		}

		dr := helper2(_node.Right, _k)
		// found target in _node.Left, distance = dl
		if dr != -1 {
			// if the distance from _node to _target == k
			if dr+1 == _k {
				_l1 = append(_l1, _node.Val)
			} else if dr+1 > _k {
				// 距离超过了 应该返回-1？
				return -1
			} else {
				// 距离不够，右子树来补
				helper1(_node.Left, _k-dr-2)
			}
			return dr + 1
		}
		return -1
	}
	helper2(root, k)
	return _l1
}

func Test863() {
	n0 := TreeNode{Val: 0}
	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 3}
	n4 := TreeNode{Val: 4}
	n5 := TreeNode{Val: 5}
	n6 := TreeNode{Val: 6}
	n7 := TreeNode{Val: 7}
	n8 := TreeNode{Val: 8}

	n2.Left = &n7
	n2.Right = &n4
	n5.Left = &n6
	n5.Right = &n2
	n1.Left = &n0
	n1.Right = &n8
	n3.Left = &n5
	n3.Right = &n1

	fmt.Println(distanceK(&n3, &n5, 2))

}
