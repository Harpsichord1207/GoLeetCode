package subs

import "fmt"

func minOperations(nums []int) int {

	r := 0

	var f1 func(_nums []int)
	f1 = func(_nums []int) {
		for i, n := range _nums {
			if n%2 == 1 {
				_nums[i] -= 1
				r += 1
			}
		}
	}

	f2 := func(_nums []int) bool {
		all_zero := true
		for i, n := range _nums {
			if n != 0 {
				_nums[i] = _nums[i] / 2
				all_zero = false
			}
		}
		return all_zero
	}

	for {
		f1(nums)
		if f2(nums) {
			break
		}
		r += 1
	}

	return r
}

func Test1558() {
	nums := []int{1, 5}
	fmt.Println(minOperations(nums))
}
