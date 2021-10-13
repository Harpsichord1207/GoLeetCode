package subs

var numsByte = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

type pos struct {
	i int
	j int
}

func getGridNum(i int, j int) int {
	if i < 3 {
		if j < 3 {
			return 0
		}
		if j < 6 {
			return 1
		}
		if j < 9 {
			return 2
		}
	}

	if i < 6 {
		if j < 3 {
			return 3
		}
		if j < 6 {
			return 4
		}
		if j < 9 {
			return 5
		}
	}

	if i < 9 {
		if j < 3 {
			return 6
		}
		if j < 6 {
			return 7
		}
		if j < 9 {
			return 8
		}
	}

	panic("Impossible!")
}

func solveSudoku(board [][]byte)  {

	var colsMap []map[byte]bool
	var rowsMap []map[byte]bool
	var gridsMap []map[byte]bool
	dotsMap := make(map[pos]bool)

	for i:=0;i<9;i++ {
		colsMap = append(colsMap, make(map[byte]bool))
		rowsMap = append(rowsMap, make(map[byte]bool))
		gridsMap = append(gridsMap, make(map[byte]bool))
	}

	for i:=0;i<9;i++ {
		for j:=0;j<9;j++ {
			value := board[i][j]
			gridNum := getGridNum(i, j)
			if value != '.' {
				rowsMap[i][value] = true
				gridsMap[gridNum][value] = true
				colsMap[j][value] = true
			} else {
				dotsMap[pos{i, j}] = true
			}
		}
	}

	failedMap := make(map[pos]map[byte]bool)
	var stack []pos

	for {
		var firstPos pos
		hasDot := false
		for pos, v := range dotsMap {
			if v {
				firstPos = pos
				hasDot = true
				break
			}
		}

		if !hasDot {
			break
		}

		i := firstPos.i
		j := firstPos.j

		gridNum := getGridNum(i, j)
		var candidates []byte
		for _, n := range numsByte {
			if rowsMap[i][n] {
				continue
			}
			if colsMap[j][n] {
				continue
			}
			if gridsMap[gridNum][n] {
				continue
			}
			if failedMap[firstPos][n] {
				continue
			}
			candidates = append(candidates, n)
		}
		if len(candidates) == 0 {
			// 上一个位置填充的数放到失败数组里，上一个位置置为.
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			preValue := board[pre.i][pre.j]
			if failedMap[pre] == nil {
				failedMap[pre] = make(map[byte]bool)
				failedMap[pre][preValue] = true
			} else {
				failedMap[pre][preValue] = true
			}
			board[pre.i][pre.j] = '.'
		} else {
			board[i][j] = candidates[0]
			rowsMap[i][candidates[0]] = true
			colsMap[j][candidates[0]] = true
			gridsMap[gridNum][candidates[0]] = true
			dotsMap[firstPos] = false
			stack = append(stack, firstPos)
		}

	}



}

func Test37() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '8', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	}
	solveSudoku(board)
}

