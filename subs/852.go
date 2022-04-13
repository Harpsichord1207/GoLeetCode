package subs

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	i, j := 0, len(arr)-1

	for i < j {
		k := (i + j) / 2
		// 如果k到k+1是上升的，说明最高点在k的右边
		if arr[k] < arr[k+1] {
			i = k + 1
		} else {
			j = k
		}

	}
	return i
}

func Test852() {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	fmt.Println(peakIndexInMountainArray(arr))
}
