package subs

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	var pow uint32 = 31
	for num > 0 {
		res += (num & 1) << pow
		pow--
		num >>= 1
	}
	return res
}

func Test190() {
	fmt.Println(reverseBits(43261596))
}
