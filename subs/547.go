package subs

import "fmt"

func findCircleNum(isConnected [][]int) int {
	var mapSlice []map[int]struct{}

	checkOrAdd := func(_i int, _j int) {

		indexOfi := -1
		indexOfj := -1

		for ind, provinceMap := range mapSlice {
			if _, ok := provinceMap[_i]; ok {
				indexOfi = ind
			}
			if _, ok := provinceMap[_j]; ok {
				indexOfj = ind
			}
		}

		if indexOfi == -1 && indexOfj == -1 {
			m := make(map[int]struct{})
			m[_i] = struct{}{}
			m[_j] = struct{}{}
			mapSlice = append(mapSlice, m)
			return
		}

		if indexOfi == -1 {
			mapSlice[indexOfj][_i] = struct{}{}
			return
		}

		if indexOfj == -1 {
			mapSlice[indexOfi][_j] = struct{}{}
			return
		}

		if indexOfi == indexOfj {
			mapSlice[indexOfi][_i] = struct{}{}
			mapSlice[indexOfi][_j] = struct{}{}
			return
		}

		// clear indexOfj map
		for _v, _ := range mapSlice[indexOfj] {
			mapSlice[indexOfi][_v] = struct{}{}
		}
		mapSlice[indexOfj] = make(map[int]struct{})

	}

	for i, row := range isConnected {
		for j, is := range row {
			if i == j {
				continue
			}
			if is == 1 {
				checkOrAdd(i, j)
			}

		}
	}
	cnt1 := 0
	cnt2 := 0
	for _, m := range mapSlice {
		if len(m) > 0 {
			cnt1 += 1
			cnt2 += len(m)
		}
	}
	return cnt1 + len(isConnected) - cnt2
}

func Test547() {
	isConnected := [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1}}
	fmt.Println(findCircleNum(isConnected))
}
