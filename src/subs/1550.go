package subs

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, a := range arr {
		if a % 2 == 1 {
			count += 1
		} else {
			count = 0
		}
		if count == 3 {
			return true
		}
	}
	return false
}
