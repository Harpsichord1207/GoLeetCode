package subs

func findKthPositive(arr []int, k int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		m := (i + j) / 2
		v := arr[m] - m - 1 // m左边有几个缺失的数字

		// 找arr中第一个index，其左边的缺失数大于等于k
		if v >= k {
			// 等于的情况不能直接返回，因此可能不是等一个符合情况的数字
			j = m - 1
		} else {
			i = m + 1
		}
	}
	// 此时j + 1 是arr中第一个index，其左边的缺失数大于等于k，加上k之后就是目标值了
	return j + k + 1
}
