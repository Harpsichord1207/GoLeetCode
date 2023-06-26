package subs

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap := &minHeap{}
	rightHeap := &minHeap{}
	i, j := 0, len(costs)-1

	buildLeftHeap := func() {
		for {
			if leftHeap.Len() >= candidates || i > j {
				break
			}
			heap.Push(leftHeap, costs[i])
			i++
		}
	}
	buildRightHeap := func() {
		for {
			if rightHeap.Len() >= candidates || i > j {
				break
			}
			heap.Push(rightHeap, costs[j])
			j--
		}
	}
	buildLeftHeap()
	buildRightHeap()

	r, s := 0, 0
	for s < k {
		if (leftHeap.Len() == 0) || (rightHeap.Len() > 0 && (*leftHeap)[0] > (*rightHeap)[0]) {
			r += heap.Pop(rightHeap).(int)
			buildRightHeap()
		} else {
			r += heap.Pop(leftHeap).(int)
			buildLeftHeap()
		}
		s++
	}
	return int64(r)
}

func Test2462() {
	//fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	//fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))
	fmt.Println(totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2))
}
