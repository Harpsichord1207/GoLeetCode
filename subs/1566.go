package subs

import "fmt"

func containsPattern(arr []int, m int, k int) bool {

	sliceEqual := func(_s1 []int, _s2 []int) bool {
		for _i := 0; _i < len(_s1); _i++ {
			if _s1[_i] != _s2[_i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(arr)-m; i++ {
		pat := arr[i : i+m]
		res := 1
		for j := i + m; ; j += m {

			if j+m > len(arr) {
				break
			}

			if sliceEqual(pat, arr[j:j+m]) {
				res++
			} else {
				break
			}
			if res == k {
				return true
			}
		}
	}
	return false
}

func Test1566() {
	//fmt.Println(containsPattern([]int{1, 2, 4, 4, 4, 4}, 1, 3))
	//fmt.Println(containsPattern([]int{2, 2}, 1, 2))
	fmt.Println(containsPattern([]int{2, 2, 1, 2, 2, 1, 1, 1, 2, 1}, 2, 2))
}
