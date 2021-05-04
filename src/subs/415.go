package subs

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {

	l1 := len(num1)
	l2 := len(num2)
	if l1 > l2 {
		num2, num1 = num1, num2
		l1, l2 = l2, l1
	}
	var resArray = [10000]uint8{}
	for i,j := l1-1, 9999; i>=0;i,j=i-1,j-1{
		n := num1[i]
		resArray[j] = n - 48
	}
	for i,j := l2-1, 9999; i>=0;i,j=i-1,j-1{
		n := num2[i]
		resArray[j] += n
		if resArray[j] >= 58 {
			resArray[j] -= 10
			resArray[j-1] += 1
		}
	}
	if resArray[9999-l2] > 0 {
		resArray[9999-l2] += 48
		return string(resArray[9999-l2:])
	}
	return string(resArray[10000-l2:])
}

func Test415()  {
	// 48 -> 0 ... 58 ->9
	fmt.Println(addStrings("1", "9"))
}
