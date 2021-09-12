package subs

import "fmt"

func findLHS(nums []int) int {
	var res = 0
	var m = make(map[int]int)
	for _, n := range nums {
		m[n]++
		if m[n+1] > 0 {
			t := m[n+1] + m[n]
			if t > res {
				res = t
			}
		}
		if m[n-1] > 0 {
			t := m[n-1] + m[n]
			if t > res {
				res = t
			}
		}
	}
	return res
}

func Test594() {
	nums := []int{1, 3, 2, 2, 5, 3, 7}
	fmt.Println(findLHS(nums))
}
