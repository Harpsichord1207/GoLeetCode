package subs

import (
	"fmt"
	"strconv"
)

func helper842(num string, index int, nums []int, res *[]int) {
	if len(*res) != 0 {
		return
	}
	if index >= len(num) && len(nums) >= 3 {
		*res = nums
		return
	}

	for i := index; i < len(num); i++ {
		part := num[index : i+1]
		if (part[0] == '0' && len(part) > 1) || len(part) > 10 {
			break
		}
		longPart, _ := strconv.ParseInt(part, 10, 64)
		if longPart > 2147483647 {
			break
		}
		ln := len(nums)
		if (ln > 1) && (nums[ln-1]+nums[ln-2] != int(longPart)) {
			continue
		}
		nums = append(nums, int(longPart))
		helper842(num, i+1, nums, res)
		nums = nums[:len(nums)-1]
	}

}

func splitIntoFibonacci(num string) []int {
	var res []int
	helper842(num, 0, []int{}, &res)
	return res
}

func Test842() {
	fmt.Println(splitIntoFibonacci("1011"))
}
