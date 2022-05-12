package subs

import "fmt"

func minimumSize(nums []int, maxOperations int) int {
	// 最后的结果一定在1和max(nums)之间
	minSize := 1
	maxSize := -1
	for _, n := range nums {
		if maxSize < n {
			maxSize = n
		}
	}

	for minSize < maxSize {
		midSize := minSize + (maxSize-minSize)/2
		opCount := 0
		for _, n := range nums {
			opCount += (n - 1) / midSize // 当n == midSize，操作数是0，因此需要n-1
		}
		// opCount = 当每个bag最大size为midSzie时需要的操作数
		// 需要的操作数大于允许的最大操作数，则需要加大size
		if opCount > maxOperations {
			minSize = midSize + 1
		} else {
			maxSize = midSize
		}
	}
	return minSize
}

func Test1760() {
	nums := []int{9}
	fmt.Println(minimumSize(nums, 2))
}
