package subs

import (
	"container/heap"
	"fmt"
)

type Row1337 struct {
	RowNumber    int
	SoldierCount int
}

type RowHeap1337 []Row1337

func (r RowHeap1337) Len() int { return len(r) }
func (r RowHeap1337) Less(i, j int) bool {
	if r[i].SoldierCount < r[j].SoldierCount {
		return true
	}
	if r[i].SoldierCount > r[j].SoldierCount {
		return false
	}
	return r[i].RowNumber < r[j].RowNumber
}
func (r RowHeap1337) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func (r *RowHeap1337) Push(x interface{}) {
	*r = append(*r, x.(Row1337))
}

func (r *RowHeap1337) Pop() interface{} {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])

	rh := &RowHeap1337{}

	for i := 0; i < m; i++ {
		row := mat[i]

		j, k := 0, n-1
		for j <= k {
			p := (j + k) / 2
			v := row[p]
			if v == 0 {
				k = p - 1
			} else {
				j = p + 1
			}
		}
		rh.Push(Row1337{RowNumber: i, SoldierCount: k + 1})
	}
	heap.Init(rh)
	// fmt.Println(rh)
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(rh).(Row1337).RowNumber)
	}
	return res
}

func Test1337() {
	mat := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {0, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	fmt.Println(kWeakestRows(mat, 3))
}
