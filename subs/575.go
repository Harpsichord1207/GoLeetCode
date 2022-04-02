package subs

func distributeCandies(candyType []int) int {
	m := make(map[int]int)

	for _, c := range candyType {
		m[c] += 1
	}

	d := len(m)
	if d < len(candyType)/2 {
		return d
	}
	return len(candyType) / 2
}
