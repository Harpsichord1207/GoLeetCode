package subs

func chalkReplacer(chalk []int, k int) int {
	s := 0
	for _, c := range chalk {
		s += c
	}
	k = k % s
	for i, c := range chalk {
		if k >= c {
			k -= c
		} else {
			return i
		}
	}
	return 0
}
