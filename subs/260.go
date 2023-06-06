package subs

func singleNumber(nums []int) []int {
	v := 0

	// v是nums中只出现1次的2个数的异或结果
	for _, n := range nums {
		v = v ^ n
	}

	mask := 1
	for (v & mask) == 0 {
		// mask表示v二进制形式中第一个1的位置，例如mask=4 --> 100，则v二进制也是以1xx结尾
		mask = mask << 1
	}

	// 通过mask把nums分为2组，一组也是1xx结尾，另一组则数0xx
	//  - 两个只出现了一次的数必定分到了2个组
	//  - 两个组的不一定是平均分的，甚至有可能一个组只有唯一的一个数

	r1, r2 := 0, 0
	for _, n := range nums {
		if n&mask == 0 {
			r1 = r1 ^ n
		} else {
			r2 = r2 ^ n
		}
	}
	return []int{r1, r2}
}
