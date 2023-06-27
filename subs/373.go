package subs

import (
	"container/heap"
	"fmt"
)

type pair373 struct {
	i int
	j int
}

type ele373 struct {
	n1   int
	n2   int
	pair pair373
}

type heap373 []ele373

func (h heap373) Len() int { return len(h) }
func (h heap373) Less(i, j int) bool {
	s1 := h[i].n1 + h[i].n2
	s2 := h[j].n1 + h[j].n2
	if s1 < s2 {
		return true
	} else if s1 > s2 {
		return false
	}
	if h[i].pair.i < h[j].pair.i {
		return true
	} else if h[i].pair.i > h[j].pair.i {
		return false
	}
	return h[i].pair.j < h[j].pair.j
}
func (h heap373) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *heap373) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(ele373))
}
func (h *heap373) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var nums [][]int
	visited := make(map[pair373]struct{})
	l1, l2 := len(nums1), len(nums2)
	h := heap373{ele373{nums1[0], nums2[0], pair373{0, 0}}}
	for k > 0 && len(h) > 0 {
		// 上一个最小数的i或者j +1
		ele := heap.Pop(&h).(ele373)
		nums = append(nums, []int{ele.n1, ele.n2})

		pair := pair373{ele.pair.i + 1, ele.pair.j}
		if _, ok := visited[pair]; !ok && pair.i < l1 {
			heap.Push(&h, ele373{nums1[pair.i], nums2[pair.j], pair})
			visited[pair] = struct{}{}
		}
		pair = pair373{ele.pair.i, ele.pair.j + 1}
		if _, ok := visited[pair]; !ok && pair.j < l2 {
			heap.Push(&h, ele373{nums1[pair.i], nums2[pair.j], pair})
			visited[pair] = struct{}{}
		}
		k -= 1
	}
	return nums
}

func Test373() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	fmt.Println(kSmallestPairs([]int{1, 2}, []int{3}, 3))
}
