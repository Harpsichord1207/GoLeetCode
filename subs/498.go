package subs

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	var res []int
	m := len(mat)
	n := len(mat[0])

	if m == 0 || n == 0 {
		return res
	}
	row := 0
	col := 0
	direction := true

	for row < m && col < n {
		res = append(res, mat[row][col])
		if direction {
			if col == n-1 {
				row += 1
				direction = !direction
			} else if row == 0 {
				col += 1
				direction = !direction
			} else {
				row -= 1
				col += 1
			}
		} else {
			if row == m-1 {
				col += 1
				direction = !direction
			} else if col == 0 {
				row += 1
				direction = !direction
			} else {
				row += 1
				col -= 1
			}
		}
	}

	return res
}

func Test498() {
	var mat = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(findDiagonalOrder(mat))
}
