package subs

import "fmt"

func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])
	ans := 0
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' { continue }
				if j > 0 && board[i][j-1] == 'X' { continue }
				ans += 1
			}
		}
	}
	return ans
}

func Test419()  {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	fmt.Println(countBattleships(board))
}
