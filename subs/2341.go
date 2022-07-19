package subs

import "fmt"

func numberOfPairs(nums []int) []int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n] += 1
	}

	r := make([]int, 2)
	for _, v := range m {
		r[1] += v % 2
		r[0] += v / 2
	}

	return r
}

func Test2341() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	fmt.Println(numberOfPairs(nums))
}
