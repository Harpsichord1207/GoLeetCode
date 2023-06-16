package subs

type Point struct {
	x int
	y int
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	m := make(map[Point]bool, len(queens))

	for _, queen := range queens {
		p := Point{x: queen[0], y: queen[1]}
		m[p] = true
	}
	var badQueens [][]int

	// 上
	x := king[0]
	y := king[1]
	for {
		y -= 1
		if y < 0 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 下
	x = king[0]
	y = king[1]
	for {
		y += 1
		if y >= 8 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 左
	x = king[0]
	y = king[1]
	for {
		x -= 1
		if x < 0 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 右
	x = king[0]
	y = king[1]
	for {
		x += 1
		if x >= 8 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 左上
	x = king[0]
	y = king[1]
	for {
		x -= 1
		y -= 1
		if x < 0 || y < 0 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 右上
	x = king[0]
	y = king[1]
	for {
		x -= 1
		y += 1
		if x < 0 || y >= 8 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 左下
	x = king[0]
	y = king[1]
	for {
		x += 1
		y -= 1
		if x >= 8 || y < 0 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	// 右下
	x = king[0]
	y = king[1]
	for {
		x += 1
		y += 1
		if x >= 8 || y >= 8 {
			break
		}
		if m[Point{x: x, y: y}] {
			badQueens = append(badQueens, []int{x, y})
			break
		}
	}

	return badQueens
}
