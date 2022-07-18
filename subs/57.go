package subs

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	li := len(intervals)
	if li == 0 {
		return [][]int{newInterval}
	}

	var res [][]int

	if intervals[0][0] > newInterval[1] {
		res = append(res, newInterval)
		res = append(res, intervals...)
		return res
	}

	if intervals[li-1][1] < newInterval[0] {
		res = append(res, intervals...)
		res = append(res, newInterval)
		return res
	}

	// 使用二分找到newInterval[0]在intervals[*][0]有序数组中应该插入的位置，和35题一样
	i, j := 0, li-1
	for i <= j {
		m := (i + j) / 2
		v := intervals[m][0]
		if v > newInterval[0] {
			j = m - 1
		} else {
			i = m + 1
		}
	}

	// 把newInterval插入进来得到新的fullIntV
	var fullIntV [][]int
	a := false
	for k, intV := range intervals {
		if k == i {
			fullIntV = append(fullIntV, newInterval)
			a = true
		}
		fullIntV = append(fullIntV, intV)
	}
	if !a {
		fullIntV = append(fullIntV, newInterval)
	}

	//fmt.Println(i)
	//fmt.Println(fullIntV)

	// 再从前往后判断是否需要合并
	for _, intV := range fullIntV {
		if len(res) == 0 {
			res = append(res, intV)
			continue
		}

		if res[len(res)-1][1] > intV[1] {
			continue
		}

		if res[len(res)-1][1] < intV[0] {
			res = append(res, intV)
			continue
		}

		res[len(res)-1][1] = intV[1]

	}
	return res
}

func Test57() {
	in := [][]int{{1, 3}, {6, 9}}
	nin := []int{2, 5}
	fmt.Println(insert(in, nin))
}
