package subs

import "fmt"

func helper390(n int, l bool) int {
	if n <= 2 {
		if l {
			return n
		} else {
			return 1
		}
	} else if (n % 2) == 0 {
		i := 1
		if l {
			i = 0
		}
		return 2 * helper390(n/2, !l) - i
	} else {
		return 2 * helper390(n/2, !l)
	}
}

func lastRemaining(n int) int {
	return helper390(n, true)
}

func Test390() {
	//fmt.Println(lastRemaining(1))
	for i:=1;i<100;i++{
		fmt.Println(lastRemaining(i))
	}
}
