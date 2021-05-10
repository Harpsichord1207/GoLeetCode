package subs

func rangeBitwiseAnd(left int, right int) int {
	i := 0
	for ;left != right;i++ {
		left >>= 1
		right >>= 1
	}
	return left << i
}
