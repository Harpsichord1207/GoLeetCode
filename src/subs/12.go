package subs

import "fmt"

var nums = [21]int{3000, 2000, 1000, 900, 500, 400, 300, 200, 100, 90, 50, 40, 30, 20, 10, 9, 5, 4, 3, 2, 1}
var romans = [21]string{"MMM", "MM", "M", "CM", "D", "CD", "CCC", "CC", "C", "XC", "L", "XL", "XXX", "XX", "X", "IX", "V", "IV", "III", "II", "I"}

func intToRoman(num int) string {
	roman := ""
	for num > 0 {
		for i, n := range nums {
			if num >= n {
				roman += romans[i]
				num -= n
				break
			}
		}
	}
	return roman
}

func Test12()  {
	fmt.Println(intToRoman(1994))
}
