package subs

import "fmt"

func maximumRemovals(s string, p string, removable []int) int {
	isSubsequence := func(_s string, _p string) bool { return true }
	isSubsequence = func(_s string, _p string) bool {

		if len(_p) == 0 {
			return true
		}
		if len(_s) == 0 {
			return false
		}

		if _s[0] == _p[0] {
			return isSubsequence(_s[1:], _p[1:])
		}

		return isSubsequence(_s[1:], _p)
	}

	removeIndex := func(_s string, _removable []int) string {
		var _r []rune
		_m := make(map[int]bool)
		for _, _i := range _removable {
			_m[_i] = true
		}
		for _i, _c := range _s {
			if _m[_i] {
				continue
			}
			_r = append(_r, _c)
		}
		return string(_r)
	}

	i, j := 0, len(removable)
	for i <= j {
		k := i + (j-i)/2
		newS := removeIndex(s, removable[:k])
		fmt.Println(i, k, j)
		if isSubsequence(newS, p) {
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
