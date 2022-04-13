package subs

import "fmt"

func searchInsert(nums []int, target int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		k := (i + j) / 2
		v := nums[k]
		if v == target {
			return k
		}
		if v < target {
			i = k + 1
		} else {
			j = k - 1
		}
	}
	return j + 1

}

func Test35() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
