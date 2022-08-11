package subs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func fractionAddition(expression string) string {
	reg, err := regexp.Compile("[+-]?\\d+/\\d+")
	if err != nil {
		panic(err)
	}
	var numerators []int
	var denominators []int
	var D int = 1
	var length int = 0
	for _, f := range reg.FindAllString(expression, -1) {
		fmt.Println(f)
		nums := strings.Split(f, "/")
		numerator, _ := strconv.Atoi(nums[0])
		numerators = append(numerators, numerator)

		denominator, _ := strconv.Atoi(nums[1])
		denominators = append(denominators, denominator)
		D *= denominator
		length += 1
	}

	var N int = 0
	for i := 0; i < length; i++ {
		N += D / denominators[i] * numerators[i]
	}
	if N == 0 {
		return "0/1"
	}

	var neg string = ""
	if N < 0 {
		neg = "-"
		N = -N
	}
	g := gcd(N, D)
	return fmt.Sprintf("%s%d/%d", neg, N/g, D/g)
}

func Test592() {
	fmt.Println(fractionAddition("1/3-1/2"))
}
