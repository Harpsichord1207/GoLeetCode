package subs

import (
	"fmt"
	"sort"
)

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	l := len(arr)
	fmt.Println(arr)
	for i, a := range arr {
		var t int
		if a < 0 {
			if a%2 != 0 {
				continue
			}
			t = a / 2
		} else {
			t = a * 2
		}
		p, q := i+1, l-1
		for p <= q {
			m := (p + q) / 2
			v := arr[m]
			fmt.Printf("i = %d, a = %d, p = %d, q = %d, m = %d, v = %d \n", i, a, p, q, m, v)
			if v == t {
				fmt.Println(v, a, a%2)
				return true
			} else if v > t {
				q = m - 1
			} else {
				p = m + 1
			}
		}
	}
	return false
}

func Test1346() {
	arr := []int{-10, 12, -20, -8, 15}
	arr = []int{102, -592, 457, 802, 98, -132, 883, 356, -857, 461, -453, 522, 250, 476, 991, 540, -852, -485, -637, 999, -803, -691, -880, 881, -584, 750, -124, 745, -909, -892, 304, -814, 868, 665, 50, -40, 26, -242, -797, -360, -918, -741, 88, -933, -93, 360, -738, 833, -191, 563, 449, 840, 806, -87, -950, 508, 74, -448, -815, -488, 639, -334}
	fmt.Println(checkIfExist(arr))
}
