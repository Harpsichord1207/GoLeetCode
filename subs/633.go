package subs

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	half := int(math.Sqrt(float64(c)))
	for i := 0; i < half+1; i++ {
		j := c - i*i

		_a, _b := 0, j
		for _a <= _b {
			_m := (_a + _b) / 2
			_v := _m * _m
			if _v == j {
				return true
			} else if _v > j {
				_b = _m - 1
			} else {
				_a = _m + 1
			}
		}

	}
	return false
}

func Test633() {
	fmt.Println(judgeSquareSum(2147483646))
}
