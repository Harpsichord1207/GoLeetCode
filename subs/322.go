package subs

import (
	"fmt"
	"math"
)

func coinChange_(coins []int, amount int) int {

	// 递归
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

func coinChange(coins []int, amount int) int {
	// DP
	// dp[i] 表示 amount = i时的结果，则
	// dp[0] = 0, dp[i] = for coin in coins: min(dp[i], dp[i-coin]+1)

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			a := i - coin
			if a < 0 {
				continue
			}
			r := dp[i-coin] + 1
			if r < dp[i] {
				dp[i] = r
			}
		}
	}
	r := dp[amount]
	if r == math.MaxInt32 {
		return -1
	}
	return r
}

func Test322() {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 100))
}
