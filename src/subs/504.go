package subs

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	res := ""
	if num < 0{
		res = "-"
		num = -num
	}
	for {
		fmt.Println(num, num%7, num/7)
		res = strconv.Itoa(num%7) + res
		num = num/7
		if num == 0 {
			break
		}
	}
	return res
}

func Test504()  {
	fmt.Println(convertToBase7(7))
}
