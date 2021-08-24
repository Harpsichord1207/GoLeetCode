package subs

func projectionArea(grid [][]int) int {
	res := 0
	maxList := make([]int, len(grid))
	for _, line := range grid {
		lineMax := 0
		for i, v := range line {
			if v != 0 {
				res += 1
				if v > lineMax {
					lineMax = v
				}
				if v > maxList[i] {
					maxList[i] = v
				}
			}
		}
		res += lineMax
	}
	for _, j := range maxList {
		res += j
	}
	return res
}

