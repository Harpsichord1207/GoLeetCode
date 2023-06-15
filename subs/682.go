package subs

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	var stack []int

	for _, op := range operations {
		if op == "C" {
			stack = stack[:len(stack)-1]
		} else if op == "D" {
			stack = append(stack, 2*stack[len(stack)-1])
		} else if op == "+" {
			stack = append(stack, stack[len(stack)-2]+stack[len(stack)-1])
		} else {
			v, _ := strconv.Atoi(op)
			stack = append(stack, v)
		}
	}
	r := 0
	for _, s := range stack {
		r += s
	}
	return r
}

func Test682() {
	o := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(o))
}
