package subs

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < n; i++ {
		if helper766(matrix, 0, i, m, n) {
			continue
		}
		return false
	}

	for j := 1; j < m; j++ {
		if helper766(matrix, j, 0, m, n) {
			continue
		}
		return false
	}
	return true
}

func helper766(matric [][]int, pos1 int, pos2 int, m int, n int) bool {
	v := matric[pos1][pos2]

	for {
		pos1++
		pos2++
		if pos1 >= m || pos2 >= n {
			break
		}
		if matric[pos1][pos2] != v {
			return false
		}
	}
	return true
}

func Test766() {
	m := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(m))
}
