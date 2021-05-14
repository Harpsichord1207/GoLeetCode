package subs

func minMoves(nums []int) int {
	l := len(nums)
	min := int(^uint(0) >> 1)

	total := 0
	for _, n := range nums {
		total += n
		if min == -1 {
			min = n
		} else if min >n {
			min = n
		}
	}
	return total - min * l
}