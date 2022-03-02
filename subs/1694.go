package subs

import "fmt"

func reformatNumber(number string) string {
	var chars []int32
	for _, n := range number {
		if n != '-' && n != ' ' {
			chars = append(chars, n)
		}
	}
	size := len(chars)
	d := size % 3

	var res []int32

	stop := false
	for i, n := range chars {
		res = append(res, n)

		if stop {
			continue
		}

		if i%3 == 2 && i != size-1 {
			res = append(res, '-')
		} else if d == 1 && i == size-3 {
			res = append(res, '-')
			stop = true
		}
	}
	return string(res)

}

func Test1694() {
	fmt.Println(reformatNumber("1-23-45 6"))
}
