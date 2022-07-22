package subs

import (
	"fmt"
	"math"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	if k <= 0 || t < 0 {
		return false
	}

	abs := func(_num int64) int64 {
		if _num > 0 {
			return _num
		}
		return -_num
	}

	buckets := make(map[int64]int64)

	// 确保没有负数
	nums64 := make([]int64, len(nums))
	for i, num := range nums {
		nums64[i] = int64(num) - math.MinInt32
	}

	for i, num := range nums64 {
		m := num / (int64(t) + 1) // num will put into bucket[m], ie: t=2,  0, 1, 2 will be in the same bucket

		// num对应的bucket已经有数，则必定num - bucket[m] <= t
		// 并且距离i超过k距离的元素在下面的逻辑已经删除，此时仍存在bucket中的数一定符合条件
		if _, ok := buckets[m]; ok {
			return true
		}

		// 与num差小于等于t的数还可能在bucket[m-1]和bucket[m+1]中，不可能在别的bucket里了
		if v, ok := buckets[m-1]; ok && abs(v-num) <= int64(t) {
			return true
		}

		if v, ok := buckets[m+1]; ok && abs(v-num) <= int64(t) {
			return true
		}

		buckets[m] = num
		fmt.Println(buckets)
		if i >= k {
			delete(buckets, nums64[int64(i)-int64(k)]/(int64(t)+1)) // 删除与i距离k的元素
		}

	}

	return false
}

func Test220() {
	nums := []int{2147483647, -1, 2147483647}
	k := 1
	t := 2147483647
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}
