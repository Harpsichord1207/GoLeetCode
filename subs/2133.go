package subs

import (
	"fmt"
)

func checkValid(matrix [][]int) bool {
	length1 := len(matrix)
	length2 := len(matrix[0])
	if length1 != length2 {
		return false
	}
	m := make(map[int][]int)
	for i := 0; i < length1; i++ {
		if !checkArrayIsComplete(matrix[i], length1) {
			return false
		}
		for j := 0; j < length2; j++ {
			m[j] = append(m[j], matrix[i][j])
		}
	}
	for _, arr := range m {
		if !checkArrayIsComplete(arr, length1) {
			return false
		}
	}

	return true
}

func checkArrayIsComplete(arr []int, length int) bool {
	if len(arr) != length {
		return false
	}
	m := make(map[int]int)
	for _, i := range arr {
		m[i] += 1
	}
	for i := 1; i < length+1; i++ {
		if m[i] != 1 {
			return false
		}
	}
	return true
}

func Test2133() {
	matrix := [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}
	fmt.Println(checkValid(matrix))
}
