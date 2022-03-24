package subs

import "fmt"

func numRescueBoats(people []int, limit int) int {
	d := make(map[int]int)
	for _, p := range people {
		d[p] += 1
	}
	r := 0

	_f := func(_k int) int {
		for i := _k; i > 0; i-- {
			_v, _e := d[i]
			if _v <= 0 {
				continue
			}
			if !_e {
				continue
			}
			return i
		}
		return -1
	}

	for {
		all_is_zero := true
		for k, v := range d {
			if v <= 0 {
				delete(d, k)
				continue
			}
			all_is_zero = false
			if k == limit {
				d[k] -= 1
				r += 1
			} else if k < limit {
				diff := limit - k
				d[k] -= 1
				target := _f(diff)
				if target != -1 {
					d[target] -= 1
				}
				r += 1
			}
		}
		if all_is_zero {
			break
		}
	}

	return r
}

func Test881() {
	people := []int{2, 4}
	limit := 5
	fmt.Println(numRescueBoats(people, limit))
}
