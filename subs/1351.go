package subs

import "fmt"

func countNegatives(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	r, c, cnt := m-1, 0, 0

	for r >= 0 && c < n {
		if grid[r][c] < 0 {
			cnt += n - c
			r -= 1
		} else {
			c += 1
		}
	}
	return cnt
}

func Test1351() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}}
	grid = [][]int{{3, 2}, {1, 0}}
	fmt.Println(countNegatives(grid))
}
