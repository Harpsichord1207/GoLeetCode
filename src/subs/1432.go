package subs

import (
	"fmt"
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	numStr := strconv.Itoa(num)

	var large, small string
	var firstC, secondC, thirdC uint8

	for i := 0; i < len(numStr); i++ {
		if firstC == 0 {
			firstC = numStr[i]
		} else if secondC == 0 && numStr[i] != firstC {
			secondC = numStr[i]
		} else if secondC != 0 && numStr[i] != secondC && numStr[i] != firstC {
			thirdC = numStr[i]
			break
		}
	}

	if secondC == 0 {
		res, _ := strconv.Atoi(strings.Repeat("8", len(numStr)))
		return res
	}

	if firstC > '1' && firstC < '9' {
		large = strings.ReplaceAll(numStr, string([]uint8{firstC}), "9")
		small = strings.ReplaceAll(numStr, string([]uint8{firstC}), "1")
	} else if firstC == '1' {
		large = strings.ReplaceAll(numStr, "1", "9")
		if secondC == '0' {
			if thirdC == 0 {
				small = numStr
			} else {
				small = strings.ReplaceAll(numStr, string([]uint8{thirdC}), "0")
			}
		} else {
			small = strings.ReplaceAll(numStr, string([]uint8{secondC}), "0")
		}
	} else {
		large = strings.ReplaceAll(numStr, "9", "1")
		if secondC == '9' {
			if thirdC == 0 {
				small = numStr
			} else {
				small = strings.ReplaceAll(numStr, string([]uint8{thirdC}), "9")
			}
		} else {
			small = strings.ReplaceAll(numStr, string([]uint8{secondC}), "9")
		}
	}

	fmt.Println(string([]uint8{firstC, secondC, thirdC}))
	fmt.Println(large, small)
	largeInt, _ := strconv.Atoi(large)
	smallInt, _ := strconv.Atoi(small)
	if largeInt > smallInt {
		return largeInt - smallInt
	}
	return smallInt - largeInt
}

func Test1432() {
	//fmt.Println(maxDiff(1101057))
	//fmt.Println(maxDiff(9))
	//fmt.Println(maxDiff(555))
	fmt.Println(maxDiff(123456))
}
