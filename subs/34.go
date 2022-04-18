package subs

func searchRange(nums []int, target int) []int {
	r := []int{-1, -1}
	i, j := 0, len(nums)-1
	for i <= j {
		m := (i + j) / 2
		v := nums[m]
		if v == target {
			// 找右边界
			m1 := m
			for m1 <= j {
				mr := (m1 + j) / 2
				vr := nums[mr]
				if vr > target {
					j = mr - 1
				} else {
					m1 = mr + 1
				}
			}

			// 找左边界
			m2 := m
			for i <= m2 {
				ml := (i + m2) / 2
				vl := nums[ml]
				if vl < target {
					i = ml + 1
				} else {
					m2 = ml - 1
				}
			}

			return []int{i, j}
		} else if v < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return r
}
