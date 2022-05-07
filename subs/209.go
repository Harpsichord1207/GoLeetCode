package subs

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	curr_sum, l, r := 0, 0, 0
	size := len(nums)
	min := size + 1
	for r < size {
		curr_sum += nums[r]

		for curr_sum >= target {
			if min > r-l {
				min = r - l
			}
			curr_sum -= nums[l]
			l += 1
		}
		r += 1

	}
	if min == size+1 {
		return 0
	}
	return min + 1
}

func Test209() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}
