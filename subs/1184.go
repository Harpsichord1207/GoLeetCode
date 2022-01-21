package subs

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	 var a, b = 0, 0
	 f := false
	 if start > destination {
		 start, destination = destination, start
	 }
	 for i, d := range distance {
		 a += d
		 if i == start {
			 f = true
		 }
		 if i == destination {
			 f = false
		 }
		 if f {
			 b += d
		 }
	 }
	 c := a - b
	 if b > c {
		 return c
	 }
	 return b
}

func Test1184() {
	distance := []int{3,6,7,2,9,10,7,16,11}
	fmt.Println(distanceBetweenBusStops(distance, 6, 2))
}