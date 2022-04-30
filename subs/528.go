package subs

import "math/rand"

// Solution528 https://leetcode.com/problems/random-pick-with-weight/discuss/154044/Java-accumulated-freq-sum-and-binary-search

type Solution528 struct {
	wSum []int
	size int
}

func Constructor528(w []int) Solution528 {
	l := len(w)
	wSum := make([]int, l)
	for i, n := range w {
		if i == 0 {
			wSum[0] = n
		} else {
			wSum[i] = wSum[i-1] + n
		}
	}
	return Solution528{wSum: wSum, size: l}
}

func (s *Solution528) PickIndex() int {
	r := rand.Intn(s.wSum[s.size-1]) + 1 // think why +1
	i, j := 0, s.size-1

	for i <= j {
		m := i + (j-i)/2
		v := s.wSum[m]

		if v == r {
			return m
		}

		if v > r {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return j
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
