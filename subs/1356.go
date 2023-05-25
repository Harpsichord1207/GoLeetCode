package subs

import "sort"

type MySlice1356 []int

func (s MySlice1356) Len() int {
	return len(s)
}

func (s MySlice1356) Less(i int, j int) bool {
	i1 := count1(s[i])
	j1 := count1(s[j])
	if i1 < j1 {
		return true
	}
	if i1 == j1 {
		return s[i] < s[j]
	}
	return false
}

func (s MySlice1356) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func count1(n int) int {
	// Brian Kernighan
	a := 0
	for n > 0 {
		a++
		n &= n - 1
	}
	return a
}

func sortByBits(arr []int) []int {
	s1356 := MySlice1356(arr)
	sort.Sort(s1356)
	return s1356
}
