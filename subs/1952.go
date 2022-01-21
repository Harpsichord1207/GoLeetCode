package subs

import "math"

func isThree(n int) bool {
	h := int(math.Sqrt(float64(n)))
	r := 0
	m := make(map[int]bool)

	for i := 1; i <= h; i++ {
		j := n / i
		if m[j] || i*j != n {
			continue
		}
		m[i] = true
		m[j] = true
		if i == j {
			r++
		} else {
			r += 2
		}
		if r > 3 {
			return false
		}
	}
	return r == 3
}
