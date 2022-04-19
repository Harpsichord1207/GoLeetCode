package subs

func twoSum(numbers []int, target int) []int {
	q := len(numbers) - 1
	for i, n := range numbers {
		m := target - n
		p := i + 1

		for p <= q {
			o := (p + q) / 2
			v := numbers[o]
			if v == m {
				return []int{i + 1, o + 1}
			} else if v > m {
				q = o - 1
			} else {
				p = o + 1
			}
		}
	}
	return nil
}
