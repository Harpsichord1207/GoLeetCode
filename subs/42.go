package subs

import "fmt"

func trap(height []int) int {
	r := 0
	right := len(height) - 1
	left := 0

	if right < 0 {
		return r
	}

	rightMax := height[right]
	leftMax := height[left]

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				r += leftMax - height[left]
			}
			left += 1
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				r += rightMax - height[right]
			}
			right -= 1
		}
	}
	return r
}

func Test42() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
