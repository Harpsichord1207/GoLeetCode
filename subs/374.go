package subs

func guess(num int) int { return -1 }

func guessNumber(n int) int {
	i, j := 1, n
	for i <= j {
		k := (i + j) / 2
		v := guess(k)
		if v == 0 {
			return k
		} else if v == 1 {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	return -1
}
