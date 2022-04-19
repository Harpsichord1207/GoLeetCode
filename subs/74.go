package subs

import "fmt"

func searchMatrix74(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	if matrix[0][0] > target || matrix[m-1][n-1] < target {
		return false
	}

	// step1: find the first row that the last element larger than or euqal to target
	i, j := 0, m-1
	target_row := -1
	for i <= j {
		k := (i + j) / 2
		v := matrix[k][n-1]
		if v == target {
			target_row = k
			break
		}
		if v > target {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	if target_row == -1 {
		target_row = j + 1
	}

	fmt.Printf("Target Row %d = %v\n", target_row, matrix[target_row])

	// step2: find target in target row
	i, j = 0, n-1
	for i <= j {
		k := (i + j) / 2
		v := matrix[target_row][k]
		if v == target {
			return true
		}
		if v > target {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	return false
}

func Test74() {
	grid := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix74(grid, 3))
}
