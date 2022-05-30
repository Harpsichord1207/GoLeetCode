package subs

func smallestDivisor(nums []int, threshold int) int {

	div := func(a int, b int) int {
		r := a / b
		if a%b != 0 {
			return r + 1
		}
		return r
	}

	i := 1
	j := -1
	for _, n := range nums {
		if j < n {
			j = n
		}
	}

	for i < j {
		k := i + (j-i)/2
		v := 0
		for _, n := range nums {
			v += div(n, k)
		}
		if v > threshold {
			i = k + 1
		} else {
			j = k
		}

	}
	return i

}
