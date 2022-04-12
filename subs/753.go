package subs

import "fmt"

func search753(nums []int, target int) int {

	i, j := 0, len(nums)-1
	for i <= j {
		k := (i + j) / 2
		v := nums[k]
		// fmt.Println(i, k, j, v)
		if v == target {
			return k
		} else if v > target {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	return -1

}

func Test753() {
	nums := []int{5}
	target := 5
	fmt.Println(search753(nums, target))
}
