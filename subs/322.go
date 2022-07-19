package subs

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {

	var m = map[int]int{0: 0}

	helper := func(_amount int) int { return 0 }
	helper = func(_amount int) int {
		if v, ok := m[_amount]; ok {
			return v
		}

		r := -2
		for _, coin := range coins {

			if coin > _amount {
				continue
			}
			v := helper(_amount - coin)
			if v == -1 {
				continue
			}

			if r == -2 || r > v {
				r = v
			}

		}
		m[_amount] = r + 1
		return r + 1

	}

	return helper(amount)
}

func Test322() {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 100))
}
