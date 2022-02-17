package subs

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	l := len(s1)

	diff_cnt := 0

	var a byte
	var b byte
	for i := 0; i < l; i++ {

		if s1[i] != s2[i] {
			diff_cnt += 1

			if a == 0 {
				a = s1[i]
				b = s2[i]
			} else {
				if diff_cnt > 2 {
					return false
				}
				if a != s2[i] || b != s1[i] {
					return false
				}
			}
		}
	}
	return diff_cnt == 0 || diff_cnt == 2
}

func Test1790() {
	s1 := "qgqeg"
	s2 := "gqgeq"
	fmt.Println(areAlmostEqual(s1, s2))
}
