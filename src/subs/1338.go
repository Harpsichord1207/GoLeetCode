package subs

import (
	"fmt"
	"sort"
)

func minSetSize(arr []int) int {
	l := len(arr)
	m := make(map[int]int, l)
	for i:=0;i<l;i++ {
		m[arr[i]] ++
	}

	var values []int
	for _, v := range m {
		values = append(values, v)
	}

	sort.Ints(values)
	res := 0
	acc := 0

	for i:=len(values)-1;i>=0;i--{
		acc += values[i]
		res += 1
		if acc >= l / 2 {break}
	}

	return res
}

func Test1338()  {
	fmt.Println(minSetSize([]int{3,3,3,3,5,5,5,2,2,7}))
}