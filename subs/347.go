package subs

import "container/heap"

type NumFreq struct {
	Num int
	Freq int
}

type Heap347 []NumFreq

func (h Heap347) Len() int {
	return len(h)
}

func (h Heap347) Less(i, j int) bool {
	return h[i].Freq > h[j].Freq
}

func (h Heap347) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap347) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(NumFreq))
}

func (h *Heap347) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Same as 692
func topKFrequent347(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	heap347 := &Heap347{}

	for num, freq := range m {
		*heap347 = append(*heap347, NumFreq{Num: num, Freq: freq})
	}

	heap.Init(heap347)

	var res []int

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(heap347).(NumFreq).Num)
	}
	return res
}

