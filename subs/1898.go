package subs

import "fmt"

func maximumRemovals(s string, p string, removable []int) int {

	isSubsequenceV2 := func(_s string, _p string, _m map[int]bool) bool {
		_i, _j := 0, 0
		for {
			if _m[_i] {
				_i += 1
				continue
			}
			if len(p[_j:]) == 0 {
				return true
			}
			if len(s[_i:]) == 0 {
				return false
			}
			if _s[_i] == _p[_j] {
				_i++
				_j++
				continue
			}
			_i++
		}
	}

	i, j := 0, len(removable)
	for i <= j {
		k := i + (j-i)/2
		mm := make(map[int]bool)
		for _, v := range removable[:k] {
			mm[v] = true
		}
		if isSubsequenceV2(s, p, mm) {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	return j
}

func Test1898() {
	s := "abcacb"
	p := "ab"
	r := []int{3, 1, 0}
	fmt.Println(maximumRemovals(s, p, r))
}
