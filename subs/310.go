package subs

import (
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {

	if n == 1 {
		return []int{0}
	}

	pathMap := make(map[int][]int)

	// pathMap: key = 每个节点的值, value = 与之相连的节点的值
	for _, edge := range edges {
		pathMap[edge[0]] = append(pathMap[edge[0]], edge[1])
		pathMap[edge[1]] = append(pathMap[edge[1]], edge[0])
	}

	// 只有一个相连节点的节点，一定是叶子节点
	var leaves []int
	for k, vs := range pathMap {
		if len(vs) == 1 {
			leaves = append(leaves, k)
		}
	}

	// 最后的结果一定不超过2个，循环去除叶子节点，最后剩下的就是根节点
	for n > 2 {
		leafSize := len(leaves)
		var newLeaves []int
		n -= leafSize
		fmt.Println(leafSize)
		for _, lv := range leaves {
			// 全部遍历然后删除叶子节点当然可行，但会超时
			// 更好的做法是找到叶子节点连接的另一个节点，从另一个节点的关联节点数组中删除这个叶子节点
			node := pathMap[lv][0]
			le := len(pathMap[node])
			for i := 0; i < le; i++ {
				if pathMap[node][i] == lv {
					pathMap[node][i] = pathMap[node][le-1]
					break
				}
			}
			pathMap[node] = pathMap[node][:le-1]

			if le-1 == 1 {
				newLeaves = append(newLeaves, node)
			}
		}
		leaves = newLeaves
	}
	return leaves
}

func Test310() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	fmt.Println(findMinHeightTrees(n, edges))
}
