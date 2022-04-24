package subs

func search33(nums []int, target int) int {

	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2

		p, v, q := nums[i], nums[m], nums[j]

		if v == target {
			return m
		}

		// v和q都在同一个递增的半边序列中
		if v <= q {
			// 如果target也在两者之间
			if v < target && target <= q {
				i = m + 1
			} else {
				j = m - 1
			}
			// v和q分属两个半边序列，即p和v在同一个递增的半边序列中
		} else {

			if p <= target && target < v {
				j = m - 1
			} else {
				i = m + 1
			}
		}
	}
	return -1
}
