package subs

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := -1

	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	helper := func(s int) int {
		r := 0
		for _, p := range piles {
			r += p / s
			if p%s > 0 {
				r += 1
			}

		}
		return r
	}

	for min < max {
		mid := min + (max-min)/2
		if helper(mid) > h {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func Test875() {
	piles := []int{3, 6, 7, 11}
	fmt.Println(minEatingSpeed(piles, 8))
}
