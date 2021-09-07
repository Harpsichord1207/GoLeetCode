package subs

import (
	"container/heap"
	"fmt"
)

type PriorityQueue295 []int

func (p PriorityQueue295) Len() int { return len(p) }

func (p PriorityQueue295) Less(i, j int) bool { return p[i] < p[j] }

func (p PriorityQueue295) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *PriorityQueue295) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*p = append(*p, x.(int))
}

func (p *PriorityQueue295) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p PriorityQueue295) Top() float64 {
	return float64(p[0])
}

type MedianFinder struct {
	Small *PriorityQueue295
	Lager *PriorityQueue295
}

func Constructor295() MedianFinder {
	m := MedianFinder{}
	m.Small = new(PriorityQueue295)
	m.Lager = new(PriorityQueue295)
	return m
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.Small, num)
	heap.Push(m.Lager, -heap.Pop(m.Small).(int))
	if m.Small.Len() < m.Lager.Len() {
		heap.Push(m.Small, -heap.Pop(m.Lager).(int))

	}
}

func (m *MedianFinder) FindMedian() float64 {
	if m.Small.Len() > m.Lager.Len() {
		return m.Small.Top()
	}
	return (m.Small.Top() - m.Lager.Top()) * 0.5
}

func Test295() {
	m := Constructor295()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.Small, m.Lager)
	m.AddNum(3)
	fmt.Println(m.Small, m.Lager)
	fmt.Println(m.Small.Top(), m.Lager.Top())
	fmt.Println(m.FindMedian())
}
