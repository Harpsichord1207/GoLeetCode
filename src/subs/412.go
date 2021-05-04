package subs

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var ans []string
	for i:=1;i<=n;i++{
		b5 := i % 5 == 0
		b3 := i % 3 == 0
		if b3 && b5 {
			ans = append(ans, "FizzBuzz")
		} else if b5 {
			ans = append(ans, "Buzz")
		} else if b3 {
			ans = append(ans, "Fizz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}


func Test412()  {
	fmt.Println(fizzBuzz(15))
}
