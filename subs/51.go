package subs

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	queens := make([][]string, n)
	for i := 0; i < n; i++ {
		queens[i] = make([]string, n)
		for j := 0; j < n; j++ {
			queens[i][j] = "."
		}
	}

	var res [][]string
	solveRow := func(_i int, _queens [][]string) {}
	solveRow = func(_i int, _queens [][]string) {
		if _i == n {
			res = append(res, queen2string(_queens))
			return
		}

		for _j := 0; _j < n; _j++ {
			if isSafe51(_i, _j, _queens) {
				_queens[_i][_j] = "Q"
				solveRow(_i+1, queens)
				queens[_i][_j] = "."
			}

		}
	}
	solveRow(0, queens)
	return res
}

func queen2string(queens [][]string) []string {
	var res []string
	for _, row := range queens {
		res = append(res, strings.Join(row, ""))
	}
	return res
}

func isSafe51(i int, j int, queens [][]string) bool {
	s := len(queens)
	for k, row := range queens {

		if k >= i {
			return true
		}

		if row[j] == "Q" {
			return false
		}

		m := j - (i - k) // 左上->右下
		if m >= 0 && m <= s-1 && row[m] == "Q" {
			return false
		}

		n := j + (i - k) // 右上->左下
		if n >= 0 && n <= s-1 && row[n] == "Q" {
			return false
		}

	}
	return true
}

func Test51() {
	for _, q := range solveNQueens(9) {
		fmt.Println(q)
	}
}
