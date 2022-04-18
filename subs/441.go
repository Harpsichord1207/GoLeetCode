package subs

func arrangeCoins(n int) int {
	i, j := 1, n
	for i <= j {
		m := (i + j) / 2
		v := m * (m + 1) / 2 // 等差数列求和
		if v == n {
			return m
		} else if v > n {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return j
}
