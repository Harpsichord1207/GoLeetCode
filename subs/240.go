package subs

import "fmt"

func binarySearch240(list []int, target int) bool {
	left, right := 0, len(list)-1
	var mid int
	var value int
	for left <= right {
		mid = (left + right) / 2
		value = list[mid]
		if value == target {
			return true
		} else if value > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		first, last := matrix[i][0], matrix[i][n-1]
		if target >= first && target <= last {
			if binarySearch240(matrix[i], target) {
				return true
			}
		}
	}
	return false
}

func Test240() {
	//matrix := [][]int{{1,4,7,11,15}, {2,5,8,12,19}, {3,6,9,16,22}, {10,13,14,17,24}, {18,21,23,26,30}}
	//fmt.Println(searchMatrix(matrix, 5))
	fmt.Println(binarySearch240([]int{-5}, -5))
}
