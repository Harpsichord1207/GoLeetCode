package subs

func mySqrt(x int) int {
	i, j := 1, x
	for i <= j {
		m := (i + j) / 2
		v := m * m
		if v == x {
			return m
		} else if v > x {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return j
}
