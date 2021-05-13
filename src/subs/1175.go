package subs

import (
	"fmt"
	"math"
)


func numPrimeArrangements(n int) int {
	primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	var primeCount int
	for _, p := range primes {
		if n >= p {primeCount++} else {break}
	}
	compositeCount := n - primeCount
	m := int(math.Pow10(9) + 7)

	r := 1

	for primeCount > 1 {
		r = r * primeCount % m
		primeCount --

	}
	for compositeCount > 1 {
		r = r * compositeCount % m
		compositeCount --
	}

	return r
}

func Test1175()  {
	fmt.Println(numPrimeArrangements(100))
}
