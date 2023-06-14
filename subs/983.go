package subs

func mincostTickets(days []int, costs []int) int {
	mark := make([]int, 366)
	for _, day := range days {
		mark[day] = 1
	}

	dp := make([]int, 366)

	min := func(_i int, _j int, _k int) int {
		if _i <= _j && _i <= _k {
			return _i
		}
		if _j <= _i && _j <= _k {
			return _j
		}
		return _k
	}

	for i := 1; i < 366; i++ {
		if mark[i] == 0 {
			// 不去旅行的日子，截止到这一天的最小费用和前一天一样
			dp[i] = dp[i-1]
			continue
		}

		// 去旅行的日子，截止到这一天的最小费用 = min(1天前费用 + 1天cost, 7天前费用 + 7天cost, 30天前费用 + 30天cost)
		a := dp[i-1] + costs[0]
		b := 0
		c := 0
		if i >= 7 {
			b = dp[i-7] + costs[1]
		} else {
			b = dp[0] + costs[1]
		}
		if i >= 30 {
			c = dp[i-30] + costs[2]
		} else {
			c = dp[0] + costs[2]
		}
		dp[i] = min(a, b, c)

	}
	return dp[365]
}
