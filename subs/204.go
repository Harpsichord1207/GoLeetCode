package subs

import "fmt"

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	// 默认所有的都是prime
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	// 所有2的倍数是prime，所有3的倍数是prime，4continue，所有5的倍数是prime...
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

func Test204() {
	fmt.Println(countPrimes(499979))
}
