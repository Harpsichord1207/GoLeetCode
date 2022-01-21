package subs

import "sort"

func distanceP2(p1 []int, p2 []int) int {
	a := p1[0] - p2[0]
	b := p1[1] - p2[1]
	return a * a + b * b
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p1p2 := distanceP2(p1, p2)
	p1p3 := distanceP2(p1, p3)
	p1p4 := distanceP2(p1, p4)

	p2p3 := distanceP2(p2, p3)
	p2p4 := distanceP2(p2, p4)

	p3p4 := distanceP2(p3, p4)

	var distanceP2s = []int{p1p2, p1p3, p1p4, p2p3, p2p4, p3p4}
	sort.Ints(distanceP2s)
	if distanceP2s[0] == 0 {
		return false
	}
	if !(distanceP2s[0] == distanceP2s[1] && distanceP2s[1] == distanceP2s[2] && distanceP2s[2] == distanceP2s[3]) {
		return false
	}

	if distanceP2s[4] != distanceP2s[5] {
		return false
	}
	if distanceP2s[0] * 2 != distanceP2s[5] {
		return false
	}
	return true
}
