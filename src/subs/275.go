package subs

func hIndex275(citations []int) int {
	res := 0
	le := len(citations)
	l, r := 0, le - 1

	for l < r {
		m := l + (r - l) / 2
		curr := le - m
		if curr > citations[m] {
			curr = citations[m]
			l = m + 1
		} else {
			r = m - 1
		}
		if curr > res {
			res = curr
		}
	}
	return res
}

