package subs

import "fmt"

var numsByte = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

type position struct {
	i int
	j int
}

func (p position) String() string {
	return fmt.Sprintf("Position(%d, %d)", p.i, p.j)
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
	dotsMap := make(map[position]bool)

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
				dotsMap[position{i, j}] = true
			}
		}
	}
	
	var positionStack []position
	failedValueMap := make(map[byte]bool)

	for {
		var firstPos position
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

		//printSudoku(board)
		//print("++++++++++++++++++++++\n")
		//fmt.Println(firstPos)
		
		i := firstPos.i
		j := firstPos.j

		gridNum := getGridNum(i, j)
		var candidates []byte
		for _, n := range numsByte {
			if rowsMap[i][n] {
				fmt.Printf("number %c in rowsMap[%d]\n", n, i)
				continue
			}
			if colsMap[j][n] {
				fmt.Printf("number %c in colsMap[%d]\n", n, j)
				continue
			}
			if gridsMap[gridNum][n] {
				fmt.Printf("number %c in gridsMap[%d]\n", n, gridNum)
				continue
			}
			if failedValueMap[n]{
				fmt.Printf("number %c has failed before\n", n)
				continue
			}
			fmt.Printf("number %c has been candidated\n", n)
			candidates = append(candidates, n)
		}

		if len(candidates) + len(positionStack) == 0 {
			panic("YOU WROTE A BUG!")
		}

		if len(candidates) == 0 {
			prePos := positionStack[len(positionStack)-1]
			positionStack = positionStack[:len(positionStack)-1]
			preValue := board[prePos.i][prePos.j]
			failedValueMap[preValue] = true
			board[prePos.i][prePos.j] = '.'
			colsMap[prePos.j][preValue] = false
			rowsMap[prePos.i][preValue] = false
			gridsMap[getGridNum(prePos.i, prePos.j)][preValue] = false
			dotsMap[prePos] = true
			fmt.Printf("stack pop, pre : %s\n", prePos)
		} else {
			board[i][j] = candidates[0]
			rowsMap[i][candidates[0]] = true
			colsMap[j][candidates[0]] = true
			gridsMap[gridNum][candidates[0]] = true
			dotsMap[firstPos] = false
			positionStack = append(positionStack, firstPos)
			failedValueMap = make(map[byte]bool)  // clear map

			fmt.Printf("try value %s -> %c\n", firstPos, candidates[0])
			println("---------------")
			printSudoku(board)
			println("===============")
		}
		
	}
}

func printSudoku(board [][]byte)  {
	for _, row := range board {
		println(string(row))
	}
}

func Test37() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	printSudoku(board)
}

