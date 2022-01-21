package subs

import "fmt"

func findMiddleIndex(nums []int) int {
	var sum = 0
	for _, n := range nums {
		sum += n
	}
	var currSum = 0
	for i, n := range nums {
		if sum-n-currSum == currSum {
			return i
		}
		currSum += n
	}
	return -1
}

func Test1991()  {
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4}))
}