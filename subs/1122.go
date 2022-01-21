package subs

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m2 := make(map[int]int)

	for i, a := range arr2 {
		m2[a] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		v1 := arr1[i]
		v2 := arr1[j]
		i1, ok1 := m2[v1]
		i2, ok2 := m2[v2]

		if ok1 && ok2 {
			return i1 < i2
 		} else if ok1 && !ok2 {
			 return true
		} else if !ok1 && ok2 {
			return false
		} else {
			return v1 < v2
		}
	})
	return arr1
}

