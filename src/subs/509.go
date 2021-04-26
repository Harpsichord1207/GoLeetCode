package subs

import "fmt"

func fib(n int) int {
	//if n <=1 {return 1}
	//return fib(n-1) + fib(n-2)

	a := 0
	b := 1
	for i:=0;i<n;i++{
		c := a + b
		a, b = b, c
	}
	return a
}

func Test509()  {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(3))
}